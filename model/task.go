package model

// TODO mandatory fields

type Task struct {
	Directory string   `yaml:"directory" json:"directory"`
	Name      string   `yaml:"name" json:"name"`
	Checks    []*Check `yaml:"checks" json:"checks"`
}

// TODO only deliver needed fields to frontend - no command and args
type Check struct {
	ID                string `yaml:"id" json:"id"`
	Name              string `yaml:"name" json:"name"`
	Command           string `yaml:"command" json:"command"`
	Args              string `yaml:"args" json:"args"`
	Operator          string `yaml:"operator" json:"operator"`
	Expectation       string `yaml:"expectation" json:"expectation"`
	Hint              string `yaml:"hint" json:"hint"`
	LastResult        string `yaml:"lastResult" json:"lastResult"`
	LastResultContext string `yaml:"lastResultContext" json:"lastResultContext"`
	// TODO is this enough or do we need some history and timestamping
}

func NewTask(name string, command string, args string, operator string, expectation string) *Task {
	task := new(Task)
	task.Name = name
	check := NewCheck("test", "Test", command, args, operator, expectation)
	task.Checks = append(task.Checks, check)
	return task
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
