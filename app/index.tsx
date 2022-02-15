import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './src/index.css';
import store from './src/store';
import Stories from './src/Stories';

ReactDOM.render(
  <Provider store={store}>
    <div>
      <Stories />
    </div>
  </Provider>,
  document.getElementById('root') as HTMLElement
);
