package app

import (
	"context"
	"log"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: LoadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
