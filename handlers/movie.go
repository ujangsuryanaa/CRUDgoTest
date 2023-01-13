package handlers

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	// "os"
	moviesdto "movie/dto/movie"
	dto "movie/dto/result"
	"movie/models"
	"movie/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var path_file_movie = "http://localhost:5000/uploads/"

type handlersMovie struct {
	MovieRepository repositories.MovieRepository
}

func HandlerMovie(MovieRepository repositories.MovieRepository) *handlersMovie {
	return &handlersMovie{MovieRepository}
}

func (h *handlersMovie) FindMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies, err := h.MovieRepository.FindMovies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range movies {
		movies[i].Image = path_file_movie + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: movies}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersMovie) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	movie, err := h.MovieRepository.GetMovie(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	movie.Image = path_file_movie + movie.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: movie}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersMovie) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	rating, _ := strconv.Atoi(r.FormValue("price"))
	request := moviesdto.MovieRequest{
		Title:  r.FormValue("title"),
		Rating: float32(rating),
		Image:  filename,
		Desc:   r.FormValue("desc"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	//cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	//resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waybeans"});

	if err != nil {
		fmt.Println(err.Error())
	}

	movie := models.Movie{
		Title:  request.Title,
		Rating: request.Rating,
		Image:  filename,
		Desc:   request.Desc,
	}

	movie, err = h.MovieRepository.CreateMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	movie, _ = h.MovieRepository.GetMovie(movie.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: movie}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersMovie) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)

	rating, _ := strconv.Atoi(r.FormValue("rating"))

	request := moviesdto.UpdateMovie{
		Title:  r.FormValue("title"),
		Rating: float32(rating),
		Image:  filename,
		Desc:   r.FormValue("desc"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	movie, err := h.MovieRepository.GetMovie(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if (request.Title) != "" {
		movie.Title = request.Title
	}

	if request.Rating != 0 {
		movie.Rating = request.Rating
	}

	if filename != "false" {
		movie.Image = request.Image
	}

	if (request.Desc) != "" {
		movie.Desc = request.Desc
	}

	movie, err = h.MovieRepository.UpdateMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseMovie(movie)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersMovie) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	movie, err := h.MovieRepository.GetMovie(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data, err := h.MovieRepository.DeleteMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertResponseMovie(u models.Movie) moviesdto.MovieResponse {
	return moviesdto.MovieResponse{
		ID:     u.ID,
		Title:  u.Title,
		Desc:   u.Desc,
		Rating: u.Rating,
		Image:  u.Image,
	}
}
