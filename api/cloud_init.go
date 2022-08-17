package api

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/io"
)

func HandleIsCloudInitDone(w http.ResponseWriter, r *http.Request) {

	// TODO cant this duplication of code not be avoided?

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tutorInit, err := io.ReadCloudInit("/root/tutor-init.yaml")
	if err != nil {
		log.Errorf("Cannot read tutor init file on path %v, error: %v", tutorInit, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprint(w, strconv.FormatBool(false))
		return
	}
	if tutorInit == nil {
		log.Errorf("Cannot read tutor init file on path %v, error: %v", tutorInit, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprint(w, strconv.FormatBool(false))
		return
	}

	for _, phase := range tutorInit.Phases {
		if phase == "Finished Successfully" {
			fmt.Fprint(w, strconv.FormatBool(true))
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	fmt.Fprint(w, strconv.FormatBool(false))
	w.WriteHeader(http.StatusOK)
}
