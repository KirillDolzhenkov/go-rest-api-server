package main

import (
	"go-rest-api-server/internal/config"
	"go-rest-api-server/internal/user"
	"go-rest-api-server/pkg/logging"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg := config.MustLoad()

	log := logging.SetupLogger(cfg.Env)

	log.Info("starting app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	log.Info("create router")
	router := httprouter.New()

	log.Info("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	log.Info("start http server")
	start(router, log, cfg)
}

func start(router *httprouter.Router, log *slog.Logger, cfg *config.Config) {
	log.Info("start server")

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("server is listening", slog.String("address", cfg.Address))

	if err := server.Serve(listener); err != nil {
		log.Error("server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
