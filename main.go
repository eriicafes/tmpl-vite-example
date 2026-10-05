package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eriicafes/tmpl"
	"github.com/eriicafes/tmpl/vite"
)

type Entry struct {
	Title, File string
}

func main() {
	config := getConfig()
	templates := setupTemplates(!config.Prod)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spas := map[string]Entry{
			"react":  {"Tmpl React", "src/react/index.tsx"},
			"svelte": {"Tmpl Svelte", "src/svelte/index.ts"},
			"vue":    {"Tmpl Vue", "src/vue/index.ts"},
		}
		entry, ok := spas[r.URL.Query().Get("spa")]
		if !ok {
			entry = Entry{"Tmpl Vanilla", ""}
		}

		if err := templates.Render(w, tmpl.Tmpl("spa", entry)); err != nil {
			log.Println(err)
		}
	})
	http.Handle("/", templates.Vite.ServePublic(handler))
	http.ListenAndServe(config.ListenAddr(), nil)
}

type Templates struct {
	*tmpl.Templates
	*vite.Vite
}

func setupTemplates(dev bool) Templates {
	v, err := vite.New(vite.Config{
		Dev:       dev,
		Output:    os.DirFS("frontend/dist"),
		DevOrigin: "http://localhost:5273",
	})
	if err != nil {
		panic(err)
	}
	tp := tmpl.New(os.DirFS("templates")).
		Funcs(v.Funcs()).
		Load("spa").
		MustParse()

	return Templates{tp, v}
}
