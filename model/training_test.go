package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTopic(t *testing.T) {

	// given
	training := getTraining()

	// when
	topic, err := training.FindTopic("topic_02")

	// then
	assert.NoError(t, err)
	assert.Equal(t, topic.Name, "Second Topic")

}

func getTraining() Training {
	topics := []*Topic{
		NewTopic("topic_01", "First Topic", "http://example.com", "", nil),
		NewTopic("topic_02", "Second Topic", "http://example.com", "", nil),
	}
	training := NewTraining("Test Training", topics)
	return *training
}
