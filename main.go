package main

import (
	"fmt"
	"net/http"

	"github.com/Captain-Leftovers/go_htmx_project_motor/view/layout"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



func main() {

	mainRouter := chi.NewRouter()

	mainRouter.Use(middleware.Logger)

	mainRouter.Handle("/", *templ.Handler(layout.LayoutMain()))
	

	fmt.Println("Server is running at http://localhost:3000/")

	http.ListenAndServe(":3000", mainRouter)
}
