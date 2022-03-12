package model

type Task struct {
	Directory string   `yaml:"directory" json:"directory"`
	Name      string   `yaml:"name" json:"name"`
	Checks    []*Check `yaml:"checks" json:"checks"`
}

func NewTask(name string, command string, args string, operator string, expectation string) *Task {
	task := new(Task)
	task.Name = name
	check := NewCheck("test", "Test", command, args, operator, expectation)
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
