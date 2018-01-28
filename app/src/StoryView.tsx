import * as React from 'react';
import { Item } from './proto/hackernews_pb';

type StoryViewProps = {
  story: Item.AsObject,
};

const StoryView: React.SFC<StoryViewProps> = (props) => {
  const url = `http://localhost:8900/article-proxy?q=${encodeURIComponent(props.story.url)}`;
  return (
    <iframe
      frameBorder="0"
      style={{
        height: '100vh',
        width: '100%',
      }}
      src={url}
    />
  );
};

export default StoryView;
