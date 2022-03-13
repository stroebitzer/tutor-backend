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
	router.HandleFunc("/liveness", handleLiveness).Methods(http.MethodGet)
	router.HandleFunc("/readiness", handleReadiness).Methods(http.MethodGet)

	// training
	router.HandleFunc("/training", api.HandleGetTraining).Methods(http.MethodGet)

	// topic
	router.HandleFunc("/topic/{id}", api.HandleGetTopic).Methods(http.MethodGet)

	// task
	router.HandleFunc("/task/{id}", api.HandleGetTask).Methods(http.MethodGet)
	router.HandleFunc("/task/{id}/markdown", api.HandleGetTaskMarkdown).Methods(http.MethodGet)
	router.HandleFunc("/task/{id}", api.HandleExecuteTask).Methods(http.MethodPatch)
	router.HandleFunc("/task/{taskId}/check/{checkId}", api.HandleExecuteTaskCheck).Methods(http.MethodPatch)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPatch},
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
