package main

import (
	"flag"
	"log"

	"github.com/mirzaahmedov/golang-url-shortner/internal/config"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage/postgres"
	"github.com/mirzaahmedov/golang-url-shortner/internal/transfer/http"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-file", "configs/config.toml", "path to the config file")
}

func main() {
	flag.Parse()

	c, err := config.Load(configPath)
	if err != nil {
		log.Println("Could not load config file")
		log.Fatal(err)
	}

	s := postgres.NewStorage(c.DatabaseURL)

	err = s.Open()
	if err != nil {
		log.Println("Could connect to the database")
		log.Fatal(err)
	}
	defer s.Close()

	r := http.NewRouter(s)

	log.Fatal(r.Run(c.BindingAddr))
}
