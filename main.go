package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/api"
)

func main() {

	router := mux.NewRouter()

	// vital signs
	router.HandleFunc("/liveness", handleLiveness).Methods("GET")
	router.HandleFunc("/readiness", handleReadiness).Methods("GET")

	// training
	router.HandleFunc("/training", api.HandleGetTraining).Methods("GET")

	// topics
	router.HandleFunc("/topics/{id}", api.HandleGetTopic).Methods("GET")

	// tasks
	// TODO tasks should get tasks - fix REST API
	// TODO find proper REST method for executes
	router.HandleFunc("/tasks/{id}", api.HandleGetTask).Methods("GET")
	router.HandleFunc("/tasksmd/{id}", api.HandleGetTaskMarkdown).Methods("GET")
	router.HandleFunc("/tasksexecute/{id}", api.HandleExecuteTask).Methods("GET")
	router.HandleFunc("/taskscheckexecute/{taskId}/{checkId}", api.HandleExecuteTaskCheck).Methods("GET")

	// TODO make this save
	handler := cors.AllowAll().Handler(router)
	// TODO automate host and port log line = > also do in campus
	log.Infof("Starting on Port %v", 8080)
	http.ListenAndServe(":8080", handler)

}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// TODO firefox issue
