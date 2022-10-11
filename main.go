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
	log.Infof("Running in %v mode", app.GetAppMode())

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
	router.HandleFunc("/task/{id}/setup", api.HandleSetupTask).Methods(http.MethodGet)
	router.HandleFunc("/task/{id}/teardown", api.HandleTeardownTask).Methods(http.MethodGet)
	router.HandleFunc("/task/{id}/markdown", api.HandleGetTaskMarkdown).Methods(http.MethodGet)
	router.HandleFunc("/task/{taskId}/check/{checkId}", api.HandleExecuteTaskCheck).Methods(http.MethodGet)

	// init
	router.HandleFunc("/cloud_init_done", api.HandleIsCloudInitDone).Methods(http.MethodGet)

	allowedOrigins := getAllowedOrigins()
	log.Infof("Allowed Origins for CORS: %v", allowedOrigins)
	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{http.MethodGet},
		AllowedHeaders:   []string{"Token"},
		AllowCredentials: true,
		// TODO remove allow credentials??? -> no basic auth is used
	})

	handler := cors.Handler(router)
	log.Info("Running tutor-backend on port 8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Cannot run tutor-backend, error %v", err)
	}

}

func getAllowedOrigins() []string {
	if app.GetAppMode() == "DEV" {
		return []string{"*"}
	}
	return []string{
		app.GetCampusUrl(),
	}
}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
