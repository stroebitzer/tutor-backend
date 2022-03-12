package model

// TODO mandatory fields

type Topic struct {
	ID        string  `yaml:"id" json:"id"`
	Name      string  `yaml:"name" json:"name"`
	SlidesUrl string  `yaml:"slidesUrl" json:"slidesUrl"`
	VideoUrl  string  `yaml:"videoUrl" json:"videoUrl"`
	Tasks     []*Task `yaml:"tasks" json:"tasks"`
}

func NewTopic(id, name, slidesUrl, videoUrl string, tasks []*Task) *Topic {
	topic := new(Topic)
	topic.ID = id
	topic.Name = name
	topic.SlidesUrl = slidesUrl
	topic.VideoUrl = videoUrl
	topic.Tasks = tasks
	return topic
}
