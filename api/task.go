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
	params := mux.Vars(r)
	directory := params["id"]

	task, err := io.ReadTask(app.TrainingDir, directory)
	if err != nil {
		log.Errorf("cannot read task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(task)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		log.Errorf("cannot marshal task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func HandleGetTaskMarkdown(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	directory := params["id"]

	md, err := io.ReadTaskMarkdown(app.TrainingDir, directory)
	if err != nil {
		log.Errorf("cannot read task markdown on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(md)
}

func HandleExecuteTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	directory := params["id"]

	task, err := io.ReadTask(app.TrainingDir, directory)
	if err != nil {
		log.Errorf("cannot read task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO error handling?
	executor.ExecuteTask(task)

	json, err := json.Marshal(task)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		log.Errorf("cannot marshal task on directory %v, error: %v", directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func HandleExecuteTaskCheck(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	directory := params["taskId"]
	checkId := params["checkId"]

	task, err := io.ReadTask(app.TrainingDir, directory)
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

	// TODO error handling
	executor.ExecuteCheck(check)

	json, err := json.Marshal(check)
	// TODO create new DTO, only stuff for frontend
	if check == nil {
		log.Errorf("cannot marshal check %v on directory %v, error: %v", check, directory, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
