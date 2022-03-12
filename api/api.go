package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/executor"
	"github.com/stroebitzer/tutor-backend/io"
	"github.com/stroebitzer/tutor-backend/model"
)

// TODO do it with a class and DI
var Training *model.Training

var trainingDir string = "/training"
var trainingFile string = ".training.yaml"

func init() {
	// value, present := os.LookupEnv("TRAINING_DIR")
	// if present {
	// 	trainingDir = value
	// } else {
	// 	log.Warn("Environment Variable TRAINING_DIR is not set")
	// }

	// value, present = os.LookupEnv("TRAINING_FILE")
	// if present {
	// 	trainingFile = value
	// } else {
	// 	log.Warn("Environment Variable TRAINING_FILE is not set")
	// }
}

func HandleGetTraining(w http.ResponseWriter, r *http.Request) {

	log.Infof("HTTP call %v", r.URL.Path)

	training := io.ReadTraining(trainingDir, trainingFile)
	json, err := json.Marshal(training)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func HandleGetTopic(w http.ResponseWriter, r *http.Request) {

	log.Infof("HTTP call %v", r.URL.Path)
	params := mux.Vars(r)
	id := params["id"]

	training := io.ReadTraining(trainingDir, trainingFile)
	topic := model.FindTopic(training, id)

	json, err := json.Marshal(topic)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	log.Infof("HTTP call %v", r.URL.Path)
	params := mux.Vars(r)
	directory := params["id"]
	task := io.ReadTask(trainingDir, directory)
	json, err := json.Marshal(task)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func HandleGetTaskMarkdown(w http.ResponseWriter, r *http.Request) {
	log.Infof("HTTP call %v", r.URL.Path)
	params := mux.Vars(r)
	directory := params["id"]
	md := io.ReadTaskMarkdown(trainingDir, directory)
	w.Write(md)
}

func HandleExecuteTask(w http.ResponseWriter, r *http.Request) {
	log.Infof("HTTP call %v", r.URL.Path)
	params := mux.Vars(r)
	directory := params["id"]
	task := io.ReadTask(trainingDir, directory)

	executor.ExecuteTask(task)

	json, err := json.Marshal(task)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func HandleExecuteTaskCheck(w http.ResponseWriter, r *http.Request) {
	log.Infof("HTTP call %v", r.URL.Path)
	params := mux.Vars(r)
	directory := params["taskId"]
	checkId := params["checkId"]

	task := io.ReadTask(trainingDir, directory)

	var foundCheck model.Check
	for _, check := range task.Checks {
		if check.ID == checkId {
			foundCheck = *check
			break
		}
	}

	// TODO   => address cannot be nil
	// TODO what if check was not found
	if &foundCheck == nil {
		log.Warnf("Cannot find task %v check %v", directory, checkId)
	}

	executor.ExecuteCheck(&foundCheck)

	json, err := json.Marshal(foundCheck)
	// TODO create new DTO, only stuff for frontend
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
