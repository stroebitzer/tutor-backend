package app

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var appMode string
var trainingDir string
var trainingFile string
var campusUrl string
var token string

func init() {
	appMode = readAppMode()
	trainingDir = readTrainingDir()
	trainingFile = readTrainingFile()
	campusUrl = readEnvironmentVariable("CAMPUS_URL")
	token = readEnvironmentVariable("TOKEN")
}

func GetAppMode() string {
	return appMode
}

func GetTrainingDir() string {
	return trainingDir
}

func GetTrainingFile() string {
	return trainingFile
}

func GetCampusUrl() string {
	return campusUrl
}

func GetToken() string {
	return token
}

func readEnvironmentVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment Variable %v is not set", key)
		return ""
	} else {
		return value
	}
}

func readTrainingDir() string {
	trainingDir, exists := os.LookupEnv("TRAINING_DIR")
	if !exists {
		return "/training"
	}
	return trainingDir
}

func readTrainingFile() string {
	trainingFile, exists := os.LookupEnv("TRAINING_FILE")
	if !exists {
		return ".training.yaml"
	}
	return trainingFile
}

func readAppMode() string {
	appMode, exists := os.LookupEnv("APP_MODE")
	if !exists {
		return "PROD"
	}
	return appMode
}
