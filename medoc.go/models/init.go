package models

import (
	"fmt"
	"os"
	"time"
)

//Topic struct
type Topic struct {
	TopicID  string
	Title    string
	Time     time.Time
	Tag      []*TopicTag
	Content  string
	IsPublic bool //true for public，false for protected
}

//TopicTag struct
type TopicTag struct {
	TagID   string
	TagName string
	Topics  []*Topic
}

//TopicMonth show the topic group by month
type TopicMonth struct {
	Month  string
	Topics []*Topic
}

var (
	topicTagJSONFile    = "posts/tag.json"
	topicMarkdownFolder = "posts"

	//Topics store all the topic
	//md 文档
	Topics []*Topic
	//TopicsGroupByMonth store the topic by month
	TopicsGroupByMonth []*TopicMonth
	//TopicsGroupByTag store all the tag
	TopicsGroupByTag []*TopicTag
)

func init() {
	if err := InitTopicList(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
