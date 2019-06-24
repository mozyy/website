// import * as serviceWorker from './serviceWorker';
import './modules/bootstrap'


// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
// serviceWorker.unregister();



 if ('serviceWorker' in navigator) {
   window.addEventListener('load', () => {
   //   navigator.serviceWorker.register('/service-worker.js').then(registration => {
   //     console.log('SW registered: ', registration);
   //   }).catch(registrationError => {
   //     console.log('SW registration failed: ', registrationError);
   //   });
   });
 }

 window.addEventListener('load', () => {
   document.querySelector('#app').innerHTML = `
      <home-index></home-index>
   `

   });