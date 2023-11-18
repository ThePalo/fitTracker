package server

import (
	"encoding/json"
	"net/http"

	app "fitTracker/src/cmd/pkg"

	"github.com/gorilla/mux"
)

type api struct {
	repository app.ExerciseRepository
}

func New(repo ExerciseRepository) *mux.Router {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/gophers", a.fetchExercises).Methods(http.MethodGet)
	r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.fetchExercise).Methods(http.MethodGet)

	return r
}

func (a *api) fetchExercises(w http.ResponseWriter, r *http.Request) {
	gophers, _ := a.repository.FetchGophers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gophers)
}

func (a *api) fetchExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gopher, err := a.repository.FetchGopherByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}

	json.NewEncoder(w).Encode(gopher)
}
