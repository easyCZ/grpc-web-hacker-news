import { RootAction } from '../actions';
import { STORIES_INIT } from '../actions/stories';

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
  readonly stories: Map<StoryId, Story>,
  readonly error: Error | null,
  readonly loading: boolean,
};

const initialState = {
  stories: new Map<StoryId, Story>(),
  error: null,
  loading: false
};

export default function (state: StoryState = initialState, action: RootAction): StoryState {

  switch (action.type) {
    case STORIES_INIT:
      return {...state, loading: true};

    default:
      return state;
  }

}
