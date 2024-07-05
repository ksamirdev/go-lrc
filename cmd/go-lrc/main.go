package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samocodes/go-lrc/env"
	"github.com/samocodes/go-lrc/web/routes"
)

type Application struct {
	config env.Config
}

func (app *Application) Serve() error {
	port := app.config.PORT

	log.Printf("🚀 Server listening to port %s", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes.Routes(),
	}

	return srv.ListenAndServe()

}

func init() {
	env.Load()
}

func main() {
	app := Application{
		config: env.DefaultConfig,
	}

	if err := app.Serve(); err != nil {
		log.Fatalf("Error starting server: %s\b", err.Error())
	}
}
