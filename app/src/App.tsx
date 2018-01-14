import * as React from 'react';
import './App.css';
import { grpc } from 'grpc-web-client';
import { HackerNewsService } from './proto/hackernews_pb_service';
import { ListStoriesRequest, ListStoriesResponse } from './proto/hackernews_pb';

const logo = require('./logo.svg');

class App extends React.Component {

  componentDidMount() {
    const request = new ListStoriesRequest();
    grpc.invoke(HackerNewsService.ListStories, {
      request: request,
      debug: true,
      host: 'http://localhost:8900',
      onMessage: (res: ListStoriesResponse) => {
        // const obj: ListStoriesResponse = res.toObject();
        console.log(res.getStory()!.getTitle());
      },
      onEnd: (res) => {
        console.log('end', res);
      }
    });
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo"/>
          <h2>Welcome to React</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.tsx</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;
