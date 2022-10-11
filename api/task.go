package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func HandleSetupTask(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	directory := params["id"]

	// TODO move to business package and test it

	absPath := app.GetTrainingDir() + "/" + directory + "/.setup/setup.sh"
	_, err = ioutil.ReadFile(absPath)
	if err != nil {
		log.Errorf("Setup File with path %s does not exist, error: %v", absPath, err)
		w.WriteHeader(http.StatusOK)
		return
	}

	// TODO what to do with the result of the execution?
	_, err = executor.Execute("/bin/bash", absPath)
	if err != nil {
		log.Errorf("Error on executing setup file on path %s, error: %v", absPath, err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleTeardownTask(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	err := verifyToken(token)
	if err != nil {
		log.Errorf("invalid token %v, error: %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	directory := params["id"]

	// TODO move to business package and test it

	absPath := app.GetTrainingDir() + "/" + directory + "/.teardown/teardown.sh"
	_, err = ioutil.ReadFile(absPath)
	if err != nil {
		log.Errorf("Teardown File with path %s does not exist, error: %v", absPath, err)
		w.WriteHeader(http.StatusOK)
		return
	}

	// TODO what to do with the result of the execution?
	_, err = executor.Execute("/bin/bash", absPath)
	if err != nil {
		log.Errorf("Error on executing teardown file on path %s, error: %v", absPath, err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
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

func HandleExecuteTaskCheck(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")

	log.Infof("TOKEN %s", token)
	log.Infof("HEADERS %+v", r.Header)

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
