// Package main starts the Stashbook server: it loads the config,
// connects to the database, and serves the API.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hbabb/stashbook/internal/config"
	"github.com/hbabb/stashbook/internal/handler"
	"github.com/hbabb/stashbook/internal/repository"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Server stopped", "err", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("create database pool: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}

	queries := repository.New(pool)
	accounts := handler.NewAccounts(queries)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/api/accounts", accounts.List)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	slog.Info("listening", "addr", srv.Addr)
	return srv.ListenAndServe()
}
