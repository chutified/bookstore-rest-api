package app

import "tommychu/workdir/026_api-example/app/handlers"

func (a *App) setRouter() {

	// books
	a.POST("/books", a.H(handlers.NewBook))
	a.GET("/books", a.H(handlers.GetAllBooks))
	a.GET("/books/{id}", a.H(handlers.GetBook))
	a.PUT("/books/{id}", a.H(handlers.UpdateBook))
	a.DELETE("/books/{id}", a.H(handlers.RemoveBook))
}
