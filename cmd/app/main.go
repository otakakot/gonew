package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/otakakot/gonew/internal/cache"
	"github.com/otakakot/gonew/internal/config"
	"github.com/otakakot/gonew/internal/database"
	"github.com/otakakot/gonew/internal/env"
	"github.com/otakakot/gonew/internal/handler"
)

func main() {
	env.Init()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.PostgresURL)
	if err != nil {
		panic(err)
	}

	cc, err := cache.New(cfg.RedisURL)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	hdl := handler.New(db, cc)

	mux.HandleFunc("/database", hdl.Database)
	mux.HandleFunc("/cache", hdl.Cache)
	mux.HandleFunc("/env", hdl.Env)

	srv := &http.Server{ //nolint:gosec
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		slog.Info("server started")

		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("server stopped")

	timeout := 30

	ctx, cansel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	defer cansel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
