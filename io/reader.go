package io

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/stroebitzer/tutor-backend/model"
)

func ReadTraining(trainingDir string, trainingFile string) (*model.Training, error) {
	absPath := trainingDir + "/" + trainingFile
	log.Infof("Reading training from %v", absPath)
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read training from path %v, error: %v", absPath, err)
	}
	training := new(model.Training)
	err = yaml.Unmarshal(yamlFile, training)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal training from path %v, error: %v", absPath, err)
	}

	for _, topic := range training.Topics {
		for i, task := range topic.Tasks {
			fullTask, err := ReadTask(trainingDir, task.Directory)
			if err != nil {
				return nil, fmt.Errorf("cannot read task  %v, error: %v", task, err)
			}
			topic.Tasks[i] = fullTask
		}
	}

	return training, nil
}

func ReadTask(trainingDir string, directory string) (*model.Task, error) {
	absPath := trainingDir + "/" + directory + "/.task.yaml"
	log.Infof("Reading task from %v", absPath)
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read task from path %v, error: %v", absPath, err)
	}
	task := new(model.Task)
	task.Directory = directory
	err = yaml.Unmarshal(yamlFile, task)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal task from path %v, error: %v", absPath, err)
	}
	return task, nil
}

func ReadTaskMarkdown(trainingDir string, directory string) ([]byte, error) {
	absPath := trainingDir + "/" + directory + "/.task.md"
	log.Infof("Reading markdown for task from %v", absPath)
	markdownFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read task markdown from path %v, error: %v", absPath, err)
	}
	return markdownFile, nil
}

func ReadCloudInit(path string) (*model.CloudInit, error) {
	log.Infof("Reading tutor init file from %v", path)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read tutor init file from path %v, error: %v", path, err)
	}
	cloudInit := new(model.CloudInit)
	err = yaml.Unmarshal(yamlFile, cloudInit)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal tutor init file from path %v, error: %v", path, err)
	}
	return cloudInit, nil
}
