import { RootAction } from '../actions';
import { ADD_STORY, STORIES_INIT } from '../actions/stories';
import { Item } from '../proto/hackernews_pb';

export type StoryId = number;

export type Story = {
  id: StoryId,
  score: number,
  title: string,
  by: string,
  time: number,
  url: string,
  type: string,
};

export type StoryState = {
  readonly stories: { [storyId: number]: Item.AsObject },
  readonly error: Error | null,
  readonly loading: boolean,
};

const initialState = {
  stories: {},
  error: null,
  loading: false
};

export default function (state: StoryState = initialState, action: RootAction): StoryState {

  switch (action.type) {

    case STORIES_INIT:
      return {...state, loading: true};

    case ADD_STORY:
      const story: Item.AsObject = action.payload.toObject();
      if (story.id && story.id.id) {
        return {...state, loading: false, stories: {...state.stories, [story.id.id]: story}};
      }
      return state;

    default:
      return state;
  }

}
