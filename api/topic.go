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

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	training, err := io.ReadTraining(app.GetTrainingDir(), app.GetTrainingFile())
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
	if err != nil {
		log.Errorf("cannot marshal topic with id %v, error: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
