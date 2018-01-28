package hackernews

import (
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
)

type hackerNewsService struct {
	api *hackerNewsApi
}

func NewHackerNewsService(api *hackerNewsApi) *hackerNewsService {
	if api == nil {
		api = NewHackerNewsApi(nil)
	}
	return &hackerNewsService{api}
}

func (s *hackerNewsService) ListStories(req *hackernews_pb.ListStoriesRequest, resp hackernews_pb.HackerNewsService_ListStoriesServer) error {
	stories, err := s.api.TopStories()
	defer close(stories)
	if err != nil {
		return err
	}
	for story := range stories {
		resp.Send(&hackernews_pb.ListStoriesResponse{
			Story: story,
		})
	}

	return nil
}