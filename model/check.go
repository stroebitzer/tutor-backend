package model

type Check struct {
	ID                string `yaml:"id" json:"id"`
	Name              string `yaml:"name" json:"name"`
	Command           string `yaml:"command"`
	Args              string `yaml:"args"`
	Operator          string `yaml:"operator"`
	Expectation       string `yaml:"expectation" json:"expectation"`
	Hint              string `yaml:"hint" json:"hint"`
	LastResult        string `yaml:"lastResult" json:"lastResult"`
	LastResultContext string `yaml:"lastResultContext"`
}

func NewCheck(id string, name string, command string, args string, operator string, expectation string) *Check {
	check := new(Check)
	check.ID = id
	check.Name = name
	check.Command = command
	check.Args = args
	check.Operator = operator
	check.Expectation = expectation
	return check
}
