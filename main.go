package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/api"
	"github.com/stroebitzer/tutor-backend/app"
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
		AllowedOrigins:   getAllowedOrigins(),
		AllowedMethods:   []string{http.MethodGet, http.MethodPatch},
		AllowCredentials: true,
	})
	handler := cors.Handler(router)
	log.Info("Running campus-backend on port 8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Cannot run campus-backend, error %v", err)
	}

}

func getAllowedOrigins() []string {
	devOrigins := []string{"*"}
	prodOrigins := []string{"https://*.academy." + app.Domain + ":443"}
	if app.AppMode == "DEV" {
		return devOrigins
	}
	return prodOrigins
}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
