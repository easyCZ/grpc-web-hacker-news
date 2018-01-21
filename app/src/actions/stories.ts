export const STORIES_INIT = 'STORIES_INIT';

type ListStoriesInit = {
  type: typeof STORIES_INIT,
};

export const listStoriesInit = (): ListStoriesInit => ({type: STORIES_INIT});

export type StoryActionTypes =
  | ListStoriesInit;
