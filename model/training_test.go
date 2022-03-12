package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTopic(t *testing.T) {

	// given
	training := getTraining()

	// when
	topic := FindTopic(&training, "topic_02")

	// then
	assert.Equal(t, topic.Name, "Second Topic")

}

func getTraining() Training {
	training := new(Training)
	training.Name = "Test Training"
	training.Topics = make([]Topic, 2)
	training.Topics[0].ID = "topic_01"
	training.Topics[0].Name = "First Topic"
	training.Topics[0].SlidesUrl = "http://example.com"
	training.Topics = make([]Topic, 2)
	training.Topics[1].ID = "topic_02"
	training.Topics[1].Name = "Second Topic"
	training.Topics[1].SlidesUrl = "http://example.com"
	return *training
}
