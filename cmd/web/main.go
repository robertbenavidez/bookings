package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robertbenavidez/bookings/pkg/config"
	"github.com/robertbenavidez/bookings/pkg/handlers"
	"github.com/robertbenavidez/bookings/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv :=
		&http.Server{
			Addr:    portNumber,
			Handler: routes(&app),
		}

	fmt.Println("Starting Application on port", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
