import * as React from 'react';

const StoryView: React.SFC<{}> = (props) => {
  const url = 'http://localhost:8900/article-proxy?q=' +
    'https%3A%2F%2Fstackshare.io%2Fstream%2Fstream-and-go-news-feeds-for-over-300-million-end-users';
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
