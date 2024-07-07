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

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/unrolled/secure"
)

func Routes() http.Handler {
	secureMiddleware := secure.New(secure.Options{
		IsDevelopment:      env.DefaultConfig.ENVIRONMENT == "development",
		AllowedHosts:       []string{"localhost:3000"},
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})

	router := chi.NewRouter()
	router.Use(secureMiddleware.Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

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
			http.Error(w, "Body was not provided!", http.StatusBadRequest)
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
