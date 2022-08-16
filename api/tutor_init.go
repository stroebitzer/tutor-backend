package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/io"
)

func HandleGetTutorInit(w http.ResponseWriter, r *http.Request) {

	// TODO cant this duplication of code not be avoided?

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tutorInit, err := io.ReadTutorInit("/root/tutor-init.yaml")

	if err != nil {
		log.Errorf("Cannot read tutor init file on path %v, error: %v", tutorInit, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(tutorInit)
	if err != nil {
		log.Errorf("Cannot marshal tutor init file, error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
