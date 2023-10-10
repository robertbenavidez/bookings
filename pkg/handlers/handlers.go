package handlers

import (
	"net/http"

	"github.com/robertbenavidez/bookings/pkg/config"
	"github.com/robertbenavidez/bookings/pkg/models"
	"github.com/robertbenavidez/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateDate{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// test data to be sent to the template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, template."

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateDate{
		StringMap: stringMap,
	})
}

// General renders general-suite page
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "general-suite.page.tmpl", &models.TemplateDate{})
}

// Executive renders executive-suite page
func (m *Repository) Executive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "executive-suite.page.tmpl", &models.TemplateDate{})
}

// Availabilty renders the search availabilty page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "executive-suite.page.tmpl", &models.TemplateDate{})
}
