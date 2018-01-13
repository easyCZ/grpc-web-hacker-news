package hackernews

import (
	"golang.org/x/net/context"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
)

type hackerNewsService struct{}

func New() *hackerNewsService {
	return &hackerNewsService{}
}

func (s *hackerNewsService) ListStories(ctx context.Context, request *hackernews_pb.ListStoriesRequest) (*hackernews_pb.ListStoriesResponse, error) {
	var stories []*hackernews_pb.Item
	for i := 0; i < 10; i++ {
		stories = append(stories, &hackernews_pb.Item{Id: string(i)})
	}
	return &hackernews_pb.ListStoriesResponse{
		Stories: stories,
	}, nil
}
