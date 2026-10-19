package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"game-service/internal/di"
	"game-service/internal/web"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		fx.Invoke(runHTTPServer),
	).Run()
}

func runHTTPServer(lc fx.Lifecycle, h *web.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/game/", h)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("HTTP server starting on :8080")
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Printf("server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("HTTP server shutting down")
			return srv.Shutdown(ctx)
		},
	})
}
