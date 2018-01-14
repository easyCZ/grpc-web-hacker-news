package hackernews

import (
	"gopkg.in/zabawaba99/firego.v1"
	"net/http"
	"log"
	"fmt"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
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
	Item  *hackernews_pb.Item
	Error error
}

func NewHackerNewsApi(client *http.Client) *hackerNewsApi {
	firebase := firego.New("https://hacker-news.firebaseio.com", client)
	return &hackerNewsApi{
		Firebase: firebase,
	}
}

func (api *hackerNewsApi) GetStory(id int, stories chan<- *hackernews_pb.Item) {
	ref, err := api.storyRef(id)
	if err != nil {
		log.Fatalf("Failed to get story reference")
	}
	var value Item
	//var v map[string]interface{}
	if err := ref.Value(&value); err != nil {
		log.Fatal("failed to get Story %d", id, err)
	}

	var itemType hackernews_pb.ItemType
	switch value.Type {
	case "story":
		itemType = hackernews_pb.ItemType_STORY
	default:
		itemType = hackernews_pb.ItemType_UNKNOWN

	}
	stories <- &hackernews_pb.Item{
		Id:    &hackernews_pb.ItemId{Id: value.Id},
		By:    value.By,
		Score: value.Score,
		Time:  value.Time,
		Title: value.Title,
		Type:  itemType,
		Url:   value.Url,
	}
}

func (api *hackerNewsApi) TopStories(stories chan<- *hackernews_pb.Item) error {
	ref, err := api.topStoriesRef()
	if err != nil {
		return err
	}

	var ids []float64
	if err := ref.Value(&ids); err != nil {
		log.Fatalf("Failed to get top stories")
	}

	for _, id := range ids {
		go api.GetStory(int(id), stories)
	}

	return nil
}

func (api *hackerNewsApi) topStoriesRef() (*firego.Firebase, error) {
	return api.Firebase.Ref("/v0/topstories")
}

func (api *hackerNewsApi) storyRef(id int) (*firego.Firebase, error) {
	return api.Firebase.Ref(fmt.Sprintf("/v0/item/%d", id))
}
