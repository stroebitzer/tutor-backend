package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/app"
	"github.com/stroebitzer/tutor-backend/io"
)

func HandleGetTopic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	training, err := io.ReadTraining(app.TrainingDir, app.TrainingFile)
	if err != nil {
		log.Errorf("cannot read training, error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topic, err := training.FindTopic(id)
	if err != nil {
		log.Errorf("cannot find topic with id %v, error: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(topic)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		log.Errorf("cannot marshal topic with id %v, error: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
