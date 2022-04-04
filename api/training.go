package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/app"
	"github.com/stroebitzer/tutor-backend/io"
)

func HandleGetTraining(w http.ResponseWriter, r *http.Request) {
	training, err := io.ReadTraining(app.GetTrainingDir(), app.GetTrainingFile())
	if err != nil {
		log.Errorf("Cannot read training, error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(training)
	if err != nil {
		log.Errorf("Cannot marshal training, error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
