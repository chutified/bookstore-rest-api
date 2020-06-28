package app

import (
	"fmt"
	"tommychu/workdir/026_api-example/app/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) setRouter() {

	// books
	a.POST("/books", a.H(handlers.NewBook))
	a.GET("/books", a.H(handlers.GetAllBooks))
	a.GET("/books/{id}", a.H(handlers.GetBook))
	a.PUT("/books/{id}", a.H(handlers.UpdateBook))
	a.DELETE("/books/{id}", a.H(handlers.RemoveBook))

	// docs
	a.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", a.Port)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	),
	)
}
