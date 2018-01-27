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
	stories := make(chan *hackernews_pb.Item)
	defer close(stories)
	s.api.TopStories(stories)
	for story := range stories {
		resp.Send(&hackernews_pb.ListStoriesResponse{
			Story: story,
		})
	}

	return nil
}

//func (s *hackerNewsService) GetStory(ctx context.Context, req *hackernews_pb.GetStoryRequest) (*hackernews_pb.GetStoryResponse, error) {
//	if req.Id == nil {
//		return nil, errors.New("Req id is nil")
//	}
//	story, err := s.api.GetStory(int(req.Id.Id))
//	if err != nil {
//		return nil, err
//	}
//
//	body := []byte("")
//	if story.Url != "" {
//		body, err = GetPageBody(story.Url)
//		if err != nil {
//			fmt.Printf("Failed to get %v", story.Url)
//		}
//	}
//
//
//
//	return &hackernews_pb.GetStoryResponse{
//		Html:  body,
//		Story: story,
//	}, nil
//}
