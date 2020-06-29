package app

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// RequestHandlerFunc is a custom handler which use a database.
type RequestHandlerFunc func(*gorm.DB, *log.Logger, http.ResponseWriter, *http.Request)

// H wraps the http.HandlerFunc with a database.
func (a *App) H(handle RequestHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.Logger.Printf("[%s] %s (%d)\n", r.Method, r.URL.Path, r.ContentLength)
		handle(a.DB, a.Logger, w, r)
		a.Logger.Printf("--------------------\n")
	}
}

// GET handles GET methods.
func (a *App) GET(path string, f func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST handles POST methods.
func (a *App) POST(path string, f func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// PUT handles PUT methods.
func (a *App) PUT(path string, f func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// DELETE handles DELETE methods.
func (a *App) DELETE(path string, f func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
