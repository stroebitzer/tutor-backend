package io

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/stroebitzer/tutor-backend/model"
)

func ReadTraining(trainingDir string, trainingFile string) *model.Training {
	absPath := trainingDir + "/" + trainingFile
	log.Infof("Reading training from %v", absPath)
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Warnf("Cannot read training file %v", err)
		return new(model.Training)
	}
	// TODO has to use defer
	// TODO static factory
	training := new(model.Training)
	err = yaml.Unmarshal(yamlFile, training)
	if err != nil {
		log.Warnf("Cannot unmarshal yaml %v", err)
		return new(model.Training)
	}

	// this kind of sucks
	for _, topic := range training.Topics {
		for i, task := range topic.Tasks {
			fullTask := ReadTask(trainingDir, task.Directory)
			topic.Tasks[i] = *fullTask
		}
	}

	return training
}

func ReadTask(trainingDir string, directory string) *model.Task {

	// TODO add error in result

	absPath := trainingDir + "/" + directory + "/.task.yaml"
	log.Infof("Reading task from %v", absPath)
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Warnf("Cannot read task file %v", err)
		return new(model.Task)
	}
	// TODO has to use defer
	task := new(model.Task)
	task.Directory = directory
	err = yaml.Unmarshal(yamlFile, task)
	if err != nil {
		log.Warnf("Cannot unmarshal yaml %v", err)
		return new(model.Task)
	}
	return task
}

func ReadTaskMarkdown(trainingDir string, directory string) []byte {
	absPath := trainingDir + "/" + directory + "/.task.md"
	log.Infof("Reading markdown for task from %v", absPath)
	markdownFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Warnf("Cannot read markdown task file %v", err)
		return []byte("")
	}
	// TODO has to use defer
	return markdownFile
}
