package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTraining(t *testing.T) {

	// given
	dir := "reader_test"
	file := ".training.yaml"

	// when
	training := ReadTraining(dir, file)

	// then
	assert.Equal(t, "Test", training.Name)
	assert.Equal(t, 1, len(training.Topics))
	assert.Equal(t, "Test", training.Topics[0].Name)
	assert.Equal(t, 1, len(training.Topics[0].Tasks))
	assert.Equal(t, "Test", training.Topics[0].Tasks[0].Name)
	assert.Equal(t, 1, len(training.Topics[0].Tasks[0].Checks))
}

func TestReadTrainingDoesNotFail(t *testing.T) {

	// given
	dir := "not_existing"
	file := "not_existing"

	// when
	training := ReadTraining(dir, file)

	// then
	assert.Empty(t, training)
}

func TestReadTask(t *testing.T) {

	// given
	trainingDir := "reader_test"
	taskDir := "task"

	// when
	task := ReadTask(trainingDir, taskDir)

	// then
	assert.Equal(t, "Test", task.Name)
	assert.Equal(t, 1, len(task.Checks))
}

func TestReadTaskDoesNotFail(t *testing.T) {

	// given
	trainingDir := "not_existing"
	taskDir := "not_existing"

	// when
	task := ReadTask(trainingDir, taskDir)

	// then
	assert.Empty(t, task)
}

func TestReadTaskMarkdown(t *testing.T) {

	// given
	trainingDir := "reader_test"
	taskDir := "task"

	// when
	md := ReadTaskMarkdown(trainingDir, taskDir)

	// then
	assert.NotEmpty(t, md)

}

func TestReadTaskMarkdownDoesNotFail(t *testing.T) {

	// given
	trainingDir := "not_existing"
	taskDir := "not_existing"

	// when
	md := ReadTaskMarkdown(trainingDir, taskDir)

	// then
	assert.Empty(t, md)

}
