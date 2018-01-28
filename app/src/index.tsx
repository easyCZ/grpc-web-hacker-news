import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './index.css';
import store from './store';
import Stories from './Stories';

ReactDOM.render(
  <Provider store={store}>
    <div>
      <Stories />
    </div>
  </Provider>,
  document.getElementById('root') as HTMLElement
);
