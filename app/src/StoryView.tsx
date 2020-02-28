import * as React from 'react';
import { Story } from './proto/hackernews_pb';

type StoryViewProps = {
  story: Story.AsObject,
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
