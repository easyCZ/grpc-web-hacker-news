import { RootAction } from '../actions';
import { ADD_STORY, SELECT_STORY, STORIES_INIT } from '../actions/stories';
import { Story } from '../proto/hackernews_pb';

export type StoryState = {
  readonly stories: { [storyId: number]: Story.AsObject },
  readonly error: Error | null,
  readonly loading: boolean,
  readonly selected: Story.AsObject | null,
};

const initialState = {
  stories: {},
  error: null,
  loading: false,
  selected: null,
};

export default function (state: StoryState = initialState, action: RootAction): StoryState {

  switch (action.type) {

    case STORIES_INIT:
      return {...state, loading: true};

    case ADD_STORY:
      const story: Story.AsObject = action.payload.toObject();
      const selected = state.selected !== null ? state.selected : story;
      if (story.id && story.id) {
        return {
          ...state,
          loading: false,
          stories: {...state.stories, [story.id]: story},
          selected,
        };
      }
      return state;

    case SELECT_STORY:
      return {...state, selected: state.stories[action.payload]};

    default:
      return state;
  }

}
