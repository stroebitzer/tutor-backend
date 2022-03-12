package model

import (
	"fmt"
)

type Training struct {
	Name   string   `yaml:"name" json:"name"`
	Topics []*Topic `yaml:"topics" json:"topics"`
}

func NewTraining(name string, topics []*Topic) *Training {
	training := new(Training)
	training.Name = name
	training.Topics = topics
	return training
}

func (training *Training) FindTopic(id string) (*Topic, error) {
	var foundTopic *Topic
	for _, topic := range training.Topics {
		if topic.ID == id {
			foundTopic = topic
			break
		}
	}
	if foundTopic == nil {
		return nil, fmt.Errorf("cannot find topic %v in training %v", id, training)
	}
	return foundTopic, nil
}
