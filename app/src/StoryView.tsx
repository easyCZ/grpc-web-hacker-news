import * as React from 'react';
import { Item } from './proto/hackernews_pb';

type StoryViewProps = {
  story: Item.AsObject,
};

const StoryView: React.SFC<StoryViewProps> = (props) => {
  return (
    <iframe
      frameBorder="0"
      style={{
        height: '100vh',
        width: '100%',
      }}
      src={props.story.url}
    />
  );
};

export default StoryView;
