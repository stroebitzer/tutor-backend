package model

import log "github.com/sirupsen/logrus"

// TODO mandatory fields

type Training struct {
	Name   string  `yaml:"name" json:"name"`
	Topics []Topic `yaml:"topics" json:"topics"`
}

type Topic struct {
	ID        string `yaml:"id" json:"id"`
	Name      string `yaml:"name" json:"name"`
	SlidesUrl string `yaml:"slidesUrl" json:"slidesUrl"`
	VideoUrl  string `yaml:"videoUrl" json:"videoUrl"`
	Tasks     []Task `yaml:"tasks" json:"tasks"`
}

func FindTopic(training *Training, id string) Topic {
	var foundTopic Topic
	found := false
	for _, topic := range training.Topics {
		if topic.ID == id {
			foundTopic = topic
			found = true
			break
		}
	}
	// TODO this can be done better
	if !found {
		log.Warnf("Cannot find topic %v in training %v", id, training)
	}
	return foundTopic
}
