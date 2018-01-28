package hackernews

import (
	"gopkg.in/zabawaba99/firego.v1"
	"net/http"
	"log"
	"fmt"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
	"errors"
)

type Item struct {
	Id    int32  `json:id`
	Score int32  `json:score`
	Title string `json:title`
	By    string `json:by`
	Time  int32  `json:time`
	Url   string `json:url`
	Type  string `json:type`
}

type hackerNewsApi struct {
	*firego.Firebase
}

type ItemResult struct {
	Item  *hackernews_pb.Story
	Error error
}

func NewHackerNewsApi(client *http.Client) *hackerNewsApi {
	firebase := firego.New("https://hacker-news.firebaseio.com", client)
	return &hackerNewsApi{
		Firebase: firebase,
	}
}

func (api *hackerNewsApi) GetStory(id int) (*hackernews_pb.Story, error) {
	ref, err := api.storyRef(id)
	if err != nil {
		log.Fatalf("Failed to get story reference")
	}
	var value Item
	if err := ref.Value(&value); err != nil {
		log.Fatal("failed to get Story %d", id, err)
	}

	return &hackernews_pb.Story{
		Id:    value.Id,
		By:    value.By,
		Score: value.Score,
		Time:  value.Time,
		Title: value.Title,
		Url:   value.Url,
	}, nil
}

func (api *hackerNewsApi) TopStories() (chan *hackernews_pb.Story, error) {
	stories := make(chan *hackernews_pb.Story)
	ref, err := api.topStoriesRef()
	if err != nil {
		return nil, err
	}

	var ids []float64
	if err := ref.Value(&ids); err != nil {
		return nil, errors.New("failed to get top stories")
	}

	ids = ids[:10]
	for _, id := range ids {
		go func(id int) {
			story, _ := api.GetStory(id)
			stories <- story
		}(int(id))
	}

	return stories, nil
}

func (api *hackerNewsApi) topStoriesRef() (*firego.Firebase, error) {
	return api.Firebase.Ref("/v0/topstories")
}

func (api *hackerNewsApi) storyRef(id int) (*firego.Firebase, error) {
	return api.Firebase.Ref(fmt.Sprintf("/v0/item/%d", id))
}
