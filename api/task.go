package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/app"
	"github.com/stroebitzer/tutor-backend/executor"
	"github.com/stroebitzer/tutor-backend/io"
)

func HandleGetTask(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	directory := params["id"]

	task, err := io.ReadTask(app.GetTrainingDir(), directory)
	if err != nil {
		log.Errorf("cannot read task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(task)
	if err != nil {
		log.Errorf("cannot marshal task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func HandleGetTaskMarkdown(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	directory := params["id"]

	md, err := io.ReadTaskMarkdown(app.GetTrainingDir(), directory)
	if err != nil {
		log.Errorf("cannot read task markdown on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(md)
}

func HandleExecuteTask(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	directory := params["id"]

	task, err := io.ReadTask(app.GetTrainingDir(), directory)
	if err != nil {
		log.Errorf("cannot read task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	executor.ExecuteTask(task)

	json, err := json.Marshal(task)
	if err != nil {
		log.Errorf("cannot marshal task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func HandleExecuteTaskCheck(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	directory := params["taskId"]
	checkId := params["checkId"]

	task, err := io.ReadTask(app.GetTrainingDir(), directory)
	if err != nil {
		log.Errorf("cannot read task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	check := task.FindCheck(checkId)
	if check == nil {
		log.Errorf("cannot find check with id %v on directory %v", checkId, directory)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	executor.ExecuteCheck(check)

	json, err := json.Marshal(check)
	if check == nil {
		log.Errorf("cannot marshal check %v on directory %v, error: %v", check, directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
