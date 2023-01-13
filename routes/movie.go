package routes

import (
	"movie/handlers"
	"movie/pkg/middleware"
	"movie/pkg/mysql"
	"movie/repositories"

	"github.com/gorilla/mux"
)

func MovieRoutes(r *mux.Router) {
	movieRepository := repositories.RepositoryMovie(mysql.DB)
	h := handlers.HandlerMovie(movieRepository)

	r.HandleFunc("/movies", h.FindMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", h.GetMovie).Methods("GET")
	r.HandleFunc("/movie", middleware.UploadFile(h.CreateMovie)).Methods("POST")
	r.HandleFunc("/movie/{id}", middleware.UploadFile(h.UpdateMovie)).Methods("PATCH")
	r.HandleFunc("/movie/{id}", h.DeleteMovie).Methods("DELETE")
}
