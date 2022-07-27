package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/app"
	"github.com/stroebitzer/tutor-backend/io"
)

func HandleGetTraining(w http.ResponseWriter, r *http.Request) {

	// token := r.Header.Get("Token")
	// err := verifyToken(token)
	// if err != nil {
	// 	log.Errorf("invalid token %v, error: %v", token, err)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

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
