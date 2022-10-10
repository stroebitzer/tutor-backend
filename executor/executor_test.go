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

// TODO CONTAINS operator tests

func TestExecute(t *testing.T) {

	// given

	// when
	result, err := Execute("curl", "-I https://www.google.at")

	// then
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "HTTP/2 200")

}

func TestExecuteFail(t *testing.T) {

	// given

	// when
	result, err := Execute("cat", "/not/existing/file")

	// then
	assert.Error(t, err)
	assert.Empty(t, result)

}
