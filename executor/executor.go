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
	execResult := execute(check.Command, check.Args)
	execResult = strings.Trim(execResult, "'")
	result, err := compare(execResult, check.Operator, check.Expectation)
	check.LastResult = result
	if err != nil {
		check.LastResultContext = err.Error()
	}
}

func execute(command string, args string) string {
	splittedArgs := strings.Split(args, " ")
	result, err := exec.Command(command, splittedArgs...).Output()
	if err != nil {
		log.Infof("Error on executing command %v with args %v, %v", command, args, err)
	} else {
		log.Infof("Successful execution of command %v with args %v", command, args)
	}
	return string(result)
}

func compare(result string, operator string, expectation string) (string, error) {

	log.Infof("Comparing result %v, operator %v, expectation %v", result, operator, expectation)

	if operator == "EQUAL" {
		if result == expectation {
			log.Infof("Executor SUCCESS")
			return "SUCCESS", nil
		}
		log.Infof("Executor FAILURE")
		return "FAILURE", fmt.Errorf("result '%v' does not match expectation '%v'", result, expectation)
	}

	log.Errorf("Operator %v not implemented yet", operator)
	return "FAILURE", fmt.Errorf("operator %v not implemented yet", operator)

}
