import React from 'react';
import { Route } from 'react-router-dom';
import ReactDOM from 'react-dom';
import './index.css';
import { Home} from "./page";
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
      {/*<Route path="/" component={Home}/>*/}
    <Home/>
  </React.StrictMode>,
  document.getElementById('root')
);


reportWebVitals();
