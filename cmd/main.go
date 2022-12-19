package main

import (
	_ "cep-retriever/docs"
	"cep-retriever/internal/infra/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title          CEP Retriever
// @version        1.0
// @description    CEP retriever document. Fetch values like address, district and city through cep code.
// @termsOfService http://www.swagger.io/terms

// @contact.name Allan Cordeiro
// @contact.url  http://www.allancordeiro.com
// @contat.email eu@allancordeiro.com

// @host     localhost:8080
// @basePath /
func main() {
	log.Println("Server running")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/retrieve", func(r chi.Router) {
		r.Get("/{cep}", webserver.CepHandleGet)
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	log.Fatal(http.ListenAndServe(":8080", r))
}
