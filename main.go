package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/api"
)

func main() {

	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(formatter)

	router := mux.NewRouter()

	// vital signs
	router.HandleFunc("/liveness", handleLiveness).Methods("GET")
	router.HandleFunc("/readiness", handleReadiness).Methods("GET")

	// training
	router.HandleFunc("/training", api.HandleGetTraining).Methods("GET")

	// topic
	router.HandleFunc("/topics/{id}", api.HandleGetTopic).Methods("GET")

	// task
	// TODO tasks should get tasks - fix REST API
	// TODO find proper REST method for executes
	router.HandleFunc("/tasks/{id}", api.HandleGetTask).Methods("GET")
	router.HandleFunc("/tasksmd/{id}", api.HandleGetTaskMarkdown).Methods("GET")
	router.HandleFunc("/tasksexecute/{id}", api.HandleExecuteTask).Methods("GET")
	router.HandleFunc("/taskscheckexecute/{taskId}/{checkId}", api.HandleExecuteTaskCheck).Methods("GET")

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})
	handler := cors.Handler(router)
	log.Info("Running campus-backend on port 8080")
	http.ListenAndServe(":8080", handler)

}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
