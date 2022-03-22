package app

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// TODO make that all unchangeable

var AppMode string
var TrainingDir string
var TrainingFile string
var CampusDomain string

func init() {
	AppMode = getAppMode()
	TrainingDir = getTrainingDir()
	TrainingFile = getTrainingFile()
	CampusDomain = getEnvironmentVariable("CAMPUS_DOMAIN")
}

func getEnvironmentVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment Variable %v is not set", key)
		return ""
	} else {
		return value
	}
}

func getTrainingDir() string {
	trainingDir, exists := os.LookupEnv("TRAINING_DIR")
	if !exists {
		return "/training"
	}
	return trainingDir
}

func getTrainingFile() string {
	trainingFile, exists := os.LookupEnv("TRAINING_FILE")
	if !exists {
		return ".training.yaml"
	}
	return trainingFile
}

func getAppMode() string {
	appMode, exists := os.LookupEnv("APP_MODE")
	if !exists {
		return "PROD"
	}
	return appMode
}

func GetCampusUrl() string {
	return "https://" + CampusDomain + ":443"
}
