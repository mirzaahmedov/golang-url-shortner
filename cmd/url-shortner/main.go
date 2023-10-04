package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

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

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Println("Could not load config file")
		log.Fatal(err)
	}

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		}),
	)

	storage := postgres.NewStorage(cfg.DatabaseURL)

	err = storage.Open()
	if err != nil {
		log.Println("Could connect to the database")
		log.Fatal(err)
	}
	defer storage.Close()

	router := http.NewRouter(storage, logger)

	err = router.Run(cfg.BindingAddr)
	if err != nil {
		log.Fatal(err)
	}
}
