import { combineReducers, createStore } from 'redux';
import stories, { StoryState } from './reducers/stories';

interface StoreEnhancerState {
}

export interface RootState extends StoreEnhancerState {
  stories: StoryState;
}

const reducers = combineReducers<RootState>({
  stories,
});

export default createStore(reducers);
