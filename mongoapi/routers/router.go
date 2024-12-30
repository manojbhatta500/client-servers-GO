package routers

import (
	"mongoapi/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("<h1> server is working <h1>"))

	})

	router.HandleFunc("/api/movies", controllers.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.AddOneMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOneMovies).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movies/read/{id}", controllers.MarkAsRead).Methods("PATCH")

	return router

}
