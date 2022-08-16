package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO fix mess with test data files

func TestReadTraining(t *testing.T) {

	// given
	dir := "reader_test"
	file := ".training.yaml"

	// when
	training, err := ReadTraining(dir, file)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "Test", training.Name)
	assert.Equal(t, 1, len(training.Topics))
	assert.Equal(t, "Test", training.Topics[0].Name)
	assert.Equal(t, 1, len(training.Topics[0].Tasks))
	assert.Equal(t, "Test", training.Topics[0].Tasks[0].Name)
	assert.Equal(t, 1, len(training.Topics[0].Tasks[0].Checks))
}

func TestReadTrainingDoesFail(t *testing.T) {

	// given
	dir := "not_existing"
	file := "not_existing"

	// when
	training, err := ReadTraining(dir, file)

	// then
	assert.Error(t, err)
	assert.Empty(t, training)
}

func TestReadTask(t *testing.T) {

	// given
	trainingDir := "reader_test"
	taskDir := "task"

	// when
	task, err := ReadTask(trainingDir, taskDir)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "Test", task.Name)
	assert.Equal(t, 1, len(task.Checks))
}

func TestReadTaskDoesFail(t *testing.T) {

	// given
	trainingDir := "not_existing"
	taskDir := "not_existing"

	// when
	task, err := ReadTask(trainingDir, taskDir)

	// then
	assert.Error(t, err)
	assert.Empty(t, task)
}

func TestReadTaskMarkdown(t *testing.T) {

	// given
	trainingDir := "reader_test"
	taskDir := "task"

	// when
	md, err := ReadTaskMarkdown(trainingDir, taskDir)

	// then
	assert.NoError(t, err)
	assert.NotEmpty(t, md)

}

func TestReadTaskMarkdownDoesFail(t *testing.T) {

	// given
	trainingDir := "not_existing"
	taskDir := "not_existing"

	// when
	md, err := ReadTaskMarkdown(trainingDir, taskDir)

	// then
	assert.Error(t, err)
	assert.Empty(t, md)

}

func TestReadTutorInit(t *testing.T) {

	// given

	// when
	tutorInit, err := ReadTutorInit("testdata/tutor-init.yaml")

	// then
	assert.NoError(t, err)
	assert.NotNil(t, tutorInit)
	assert.Equal(t, 10, len(tutorInit.Phases))
}
