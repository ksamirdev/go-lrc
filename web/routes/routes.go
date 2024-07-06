package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/samocodes/go-lrc/env"
	"github.com/samocodes/go-lrc/helpers"
	"github.com/samocodes/go-lrc/types"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

func Routes() http.Handler {
	secureMiddleware := secure.New(secure.Options{
		IsDevelopment:      env.DefaultConfig.ENVIRONMENT == "development",
		AllowedHosts:       []string{"localhost:3000"},
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,

		// 	Allows htmx's script to be loaded
		// ContentSecurityPolicy: "default-src 'self'; script-src 'self' https://unpkg.com 'nonce-a23gbfz9e'; style-src 'self';",
	})

	router := chi.NewRouter()
	router.Use(secureMiddleware.Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// if request header doesn't accepts text/html, then return a html template
		if !helpers.SupportsHTML(r) {
			fmt.Fprint(w, "API is up and running!")
			return
		}

		w.Header().Set("Content-Type", "text/html")

		t, _ := template.ParseFiles("web/templates/index.html")
		if err := t.Execute(w, nil); err != nil {
			log.Fatalf("Error executing template: %s", err.Error())
			fmt.Fprint(w, "Error executing template")
		}
	})

	// serving css
	router.Get("/dist/main.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "web/templates/dist/main.css")
	})

	// serving js
	router.Get("/index.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")
		http.ServeFile(w, r, "web/templates/index.js")
	})

	router.Post("/lrc", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Body was not provided", http.StatusBadRequest)
			return
		}

		var music types.Music
		if err := json.Unmarshal(b, &music); err != nil {
			http.Error(w, "Unable to unmarshal the body!", http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, helpers.GenerateLRC(music))
	})

	return router
}
