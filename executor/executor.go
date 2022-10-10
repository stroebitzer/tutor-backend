package executor

import (
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/stroebitzer/tutor-backend/model"
)

func ExecuteTask(task *model.Task) {
	log.Infof("Execute task %v", task)
	for _, check := range task.Checks {
		ExecuteCheck(check)
	}
}

func ExecuteCheck(check *model.Check) {
	log.Infof("Execute check %v", check)

	// TODO error handling => giving error stack to frontend???

	execResult, err := Execute(check.Command, check.Args)
	if err != nil {
		log.Infof("Error on executing check: %+v, error: %s ", check, err)
	}

	execResult = strings.Trim(execResult, "'")

	result, err := compare(execResult, check.Operator, check.Expectation)
	if err != nil {
		log.Infof("Error on comparing result of execution, check: %+v, error: %s ", check, err)
	}
	check.Result = result
}

func Execute(command string, args string) (string, error) {
	splittedArgs := strings.Split(args, " ")
	result, err := exec.Command(command, splittedArgs...).Output()

	if err != nil {
		log.Infof("Error on executing command %v with args %v, %v", command, args, err)
		return "", err
	}

	log.Infof("Successful execution of command %v with args %v", command, args)
	return string(result), nil
}

func compare(result string, operator string, expectation string) (string, error) {

	log.Infof("Comparing result %v, operator %v, expectation %v", result, operator, expectation)

	if operator == "EQUALS" {
		if result == expectation {
			log.Infof("Executor SUCCESS")
			return "SUCCESS", nil
		}
		log.Infof("Executor FAILURE")
		return "FAILURE", fmt.Errorf("execution result '%v' does not match '%v'", result, expectation)
	} else if operator == "CONTAINS" {
		if strings.Contains(result, expectation) {
			log.Infof("Executor SUCCESS")
			return "SUCCESS", nil
		}
		log.Infof("Executor FAILURE")
		return "FAILURE", fmt.Errorf("execution result '%v' does not contain '%v'", result, expectation)
	}

	log.Errorf("Operator %v not implemented yet", operator)
	return "FAILURE", fmt.Errorf("operator %v not implemented yet", operator)

}
