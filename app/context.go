package app

import "os"

var TrainingDir string
var TrainingFile string

func init() {
	TrainingDir = getTrainingDir()
	TrainingFile = getTrainingFile()
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
