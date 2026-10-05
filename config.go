package main

import (
	"cmp"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Prod bool
	Port string
}

func (c Config) ListenAddr() string {
	return fmt.Sprintf(":%s", c.Port)
}

func getConfig() Config {
	dev := flag.Bool("dev", false, "development mode")
	// port := flag.String("port", "8000", "application port")
	flag.Parse()

	port := cmp.Or(os.Getenv("PORT"), "8000")

	return Config{
		Prod: !*dev,
		Port: port,
	}
}
