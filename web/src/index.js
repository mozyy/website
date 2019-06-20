if ('serviceWorker' in navigator) {
 window.addEventListener('load', ()=>{
   navigator.serviceWorker.register('/service-worker.js').then((registration)=>{
      console.log('sw registration: ', registration)
   }).catch(e=>{
      console.log('error: ', e)
   })
 })
}
console.log(222)