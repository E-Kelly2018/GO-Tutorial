package main

import (
	"fmt"
	"github.com/E-Kelly2018/bookings/pkg/config"
	"github.com/E-Kelly2018/bookings/pkg/handlers"
	"github.com/E-Kelly2018/bookings/pkg/render"
	"log"
	"net/http"
)

const port = ":8080"

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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on %s", port))
	_ = http.ListenAndServe(port, nil)
}
