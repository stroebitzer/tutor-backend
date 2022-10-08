package executor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stroebitzer/tutor-backend/model"
)

func TestExecuteTask(t *testing.T) {

	// given
	task := model.NewTask("test", "echo", "test", "EQUALS", "test\n")

	// when
	ExecuteTask(task)

	// then
	assert.Equal(t, "SUCCESS", task.Checks[0].Result)

}

func TestExecuteCheck(t *testing.T) {

	// given
	check := model.NewCheck("test", "Test", "echo", "test", "EQUALS", "test\n", "testSuccessText", "testFailureText")

	// when
	ExecuteCheck(check)

	// then
	assert.Equal(t, "SUCCESS", check.Result)

}

// TODO CONTAINS tests
