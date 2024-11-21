package main

import (
	"errors"
	"fmt"
	"github.com/anmho/inference/gen/protos/v1/inferencev1connect"
	"github.com/anmho/inference/inference"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log/slog"
	"net/http"
	"os"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("loading env", slog.Any("error", err))
		os.Exit(1)
	}

	config, err := env.ParseAs[Config]()
	if err != nil {
		slog.Error("parsing env", slog.Any("error", err))
		os.Exit(1)
	}

	openaiClient := openai.NewClient()
	inferenceService := inference.NewService(openaiClient)

	mux := http.NewServeMux()
	mux.Handle(inferencev1connect.NewInferenceServiceHandler(inferenceService))

	slog.Info("server listening", slog.Int("port", config.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), h2c.NewHandler(mux, &http2.Server{}))
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("unknown server error", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
