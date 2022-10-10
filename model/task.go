package model

// TODO maybe two different DTOs checks and question
type Task struct {
	Directory string    `yaml:"directory" json:"directory"`
	Name      string    `yaml:"name" json:"name"`
	Checks    []*Check  `yaml:"checks" json:"checks"`
	Question  *Question `yaml:"question" json:"question"`
}

func NewTask(name, command, args, operator, expectation string) *Task {
	task := new(Task)
	task.Name = name
	// TODO why is test stuff in here?
	check := NewCheck("test", "Test", command, args, operator, expectation, "testSuccessText", "testFailureText")
	task.Checks = append(task.Checks, check)
	return task
}

func (task *Task) FindCheck(id string) *Check {
	for _, check := range task.Checks {
		if check.ID == id {
			return check
		}
	}
	return nil
}
