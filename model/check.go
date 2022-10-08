package model

type Check struct {
	ID          string `yaml:"id" json:"id"`
	Name        string `yaml:"name" json:"name"`
	Command     string `yaml:"command"`
	Args        string `yaml:"args"`
	Operator    string `yaml:"operator"`
	Expectation string `yaml:"expectation" json:"expectation"`
	Result      string `yaml:"result" json:"result"`
	SuccessText string `yaml:"successText" json:"successText"`
	FailureText string `yaml:"failureText" json:"failureText"`
}

// TODO create a seperate DTO fro result - too much info for frontend

func NewCheck(id, name, command, args, operator, expectation, successText, failureText string) *Check {
	check := new(Check)
	check.ID = id
	check.Name = name
	check.Command = command
	check.Args = args
	check.Operator = operator
	check.Expectation = expectation
	check.SuccessText = successText
	check.FailureText = failureText
	return check
}
