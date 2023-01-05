package main

import (
	_ "cep-retriever/docs"
	"cep-retriever/internal/infra/webserver"
	"cep-retriever/pkg/middlewares"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
// @contact.email eu@allancordeiro.com

// @host     localhost:8080
// @basePath /
func main() {
	log.Println("Server running")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.Metrics)

	r.Handle("/metrics", promhttp.Handler())
	r.Route("/retrieve", func(r chi.Router) {
		r.Get("/{cep}", webserver.CepHandleGet)
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	log.Fatal(http.ListenAndServe(":8080", r))
}
