import { Action } from 'redux';
import { Item, ListStoriesRequest, ListStoriesResponse } from '../proto/hackernews_pb';
import { GrpcAction, grpcRequest } from '../middleware/grpc';
import { Code, Metadata } from 'grpc-web-client';
import { HackerNewsService } from '../proto/hackernews_pb_service';
import { PingRequest, PingResponse } from '../proto/ping_pb';
import { PingService } from '../proto/ping_pb_service';

export const STORIES_INIT = 'STORIES_INIT';
export const ADD_STORY = 'ADD_STORY';
export const SELECT_STORY = 'SELECT_STORY';

export const ping = () => {
  return grpcRequest<PingRequest, PingResponse>({
    request: new PingRequest(),
    onEnd: (code: Code, message: string | undefined, trailers: Metadata): Action | void => {
      return;
    },
    host: 'http://localhost:8900',
    methodDescriptor: PingService.Ping,
    onMessage: (message) => {
      console.log(message);
      return;
    }
  });
};

type AddStory = {
  type: typeof ADD_STORY,
  payload: Item,
};
export const addStory = (story: Item) => ({ type: ADD_STORY, payload: story });

type ListStoriesInit = {
  type: typeof STORIES_INIT,
};
export const listStoriesInit = (): ListStoriesInit => ({type: STORIES_INIT});

export const listStories = () => {
  return grpcRequest<ListStoriesRequest, ListStoriesResponse>({
    request: new ListStoriesRequest(),
    onStart: () => listStoriesInit(),
    onEnd: (code: Code, message: string | undefined, trailers: Metadata): Action | void => {
      console.log(code, message, trailers);
      return;
    },
    host: 'http://localhost:8900',
    methodDescriptor: HackerNewsService.ListStories,
    onMessage: message => {
      const story = message.getStory();
      if (story) {
        return addStory(story);
      }
      return;
    },
  });
};

type SelectStory = {
  type: typeof SELECT_STORY,
  payload: number,
};
export const selectStory = (storyId: number): SelectStory => ({ type: SELECT_STORY, payload: storyId });

export type StoryActionTypes =
  | ListStoriesInit
  | AddStory
  | SelectStory
  | GrpcAction<ListStoriesRequest, ListStoriesResponse>;
