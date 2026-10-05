package main

import (
	"fmt"

	"github.com/eriicafes/go-dev"
	"github.com/eriicafes/go-dev/vite"
)

func main() {
	session := dev.New()
	task, err := session.RunTask(dev.Cmd{
		Dir:  dev.Dir(".."),
		Run:  dev.Package("."),
		Args: dev.Values("-dev"),
		Watch: dev.Values(
			"config.go",
			"main.go",
			"templates",
			"go.mod",
			"go.sum",
		),
		ServerAddr: ":8100",
		Commands: dev.Commands(dev.Cmd{
			Dir:   "frontend",
			Run:   dev.Binary("pnpm"),
			Args:  dev.Values("dev", dev.Pair("--port", "5273")),
			Phase: dev.Before,
		}),
		Plugins: dev.Plugins(
			vite.Refresh{Origin: "http://localhost:5273"},
		),
	})
	session.Catch(err)

	fmt.Printf("> development proxy listening on %s\n", task.URL())
	session.Catch(session.Wait())
}
