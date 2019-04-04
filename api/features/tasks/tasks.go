package tasks

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Task struct {
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetTask)
	router.Post("/", CreateTask)
	router.Put("/{taskID}", UpdateTask)
	router.Delete("/{taskID}", DeleteTask)
	return router
}

func GetTask(w http.ResponseWriter, r *http.Request) {}

func CreateTask(w http.ResponseWriter, r *http.Request) {}

func UpdateTask(w http.ResponseWriter, r *http.Request) {}

func DeleteTask(w http.ResponseWriter, r *http.Request) {}
