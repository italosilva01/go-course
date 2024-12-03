package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	BuildDb()
	r := chi.NewRouter();

	r.Use(middleware.Logger)
	r.Get("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World !"))
	})
	r.Get("/see-details", func(w http.ResponseWriter, r *http.Request) {
		for key, athlete :=range dataBase {
			fmt.Fprintf(w, "key: %s. Athlet: %+v\n",key,*athlete)
		}
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("AI PAPAI. Essa rota não existe. Tente outra"))
	})

	http.ListenAndServe(":3000",r)
}