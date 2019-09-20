import './bootstrap';
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import { isDev } from './env'

ReactDOM.render(<App />, document.getElementById('root'));


const serviceConfig = {
  onSuccess: (registration: ServiceWorkerRegistration)=> {
    console.log('onSuccess', registration)
  },
  onUpdate: (registration: ServiceWorkerRegistration) => {
    console.log('onUpdate', registration)
  }
}

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
if (isDev) {
  serviceWorker.unregister();
} else {
  serviceWorker.register(serviceConfig);
  navigator.serviceWorker.ready.then(registration=> {
    // registration.addEventListener()
    console.log(registration)
  })
}