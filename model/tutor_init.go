package model

type CloudInit struct {
	Phases []string `yaml:"phases" json:"phases"`
}
