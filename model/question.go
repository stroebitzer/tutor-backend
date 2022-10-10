package model

type Question struct {
	Type        string   `yaml:"type" json:"type"`
	Options     []string `yaml:"options" json:"options"`
	SuccessText string   `yaml:"successText" json:"successText"`
	FailureText string   `yaml:"failureText" json:"failureText"`
	RightAnswer string   `yaml:"rightAnswer" json:"rightAnswer"`
}

// TODO test
