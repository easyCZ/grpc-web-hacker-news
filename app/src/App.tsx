import * as React from 'react';
import './App.css';
import { grpc } from 'grpc-web-client';
import { HackerNewsService } from './proto/hackernews_pb_service';
import { ListStoriesRequest } from './proto/hackernews_pb';

const logo = require('./logo.svg');

class App extends React.Component {

  componentDidMount() {
    const request = new ListStoriesRequest();
    grpc.unary(HackerNewsService.ListStories, {
      request: request,
      debug: true,
      host: 'http://localhost:8900',
      onEnd: (res) => {
        console.log(res);
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
