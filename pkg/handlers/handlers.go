package handlers

import (
	"github.com/E-Kelly2018/bookings/pkg/config"
	"github.com/E-Kelly2018/bookings/pkg/models"
	"github.com/E-Kelly2018/bookings/pkg/render"
	"net/http"
)

//Repo the repository used by the handles
var Repo *Repository

//Repository is the repository typer
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handles
func NewHandlers(r *Repository) {
	Repo = r
}

//Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//preform so logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Template"
	//Send data to template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
