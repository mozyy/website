(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{43:function(n,e,t){n.exports=t(66)},44:function(n,e,t){0},49:function(n,e,t){},50:function(n,e,t){},66:function(n,e,t){"use strict";t.r(e);t(44);var o=t(0),c=t.n(o),a=t(8),r=t.n(a),i=(t(49),t(90)),u=(t(50),t(22)),s=t(67),l=t(37),f=t.n(l),h=t(93),d=t(94),p=t(92),g=t(38),v=t.n(g),m=Object(s.a)(function(n){return{root:{width:"100%",height:"100vh",display:"block",touchAction:"none"},speedDial:{position:"absolute",bottom:n.spacing(2),right:n.spacing(3)}}}),w=function(n){var e,t={x:0,y:0};switch(n.type){case"touchstart":case"touchmove":case"touchend":e=n.changedTouches[0];break;case"mousedown":case"mousemove":case"mouseup":default:e=n}return t.x=e.clientX,t.y=e.clientY,t},b=!1,k=function(n){var e=m();console.log("run function components");var t,a,r,i=function(){var n=Object(o.useState)(),e=Object(u.a)(n,2),t=e[0],c=e[1],a=Object(o.useRef)(null);return[Object(o.useCallback)(function(n){c(n.getBoundingClientRect()),a.current=n.getContext("2d")},[]),a,t]}(),s=Object(u.a)(i,3),l=s[0],g=s[1],k=s[2],y=Object(o.useState)(""),E=Object(u.a)(y,2),j=E[0],O=E[1],C=Object(o.useState)(!1),S=Object(u.a)(C,2),x=S[0],M=S[1],T=Object(o.useRef)([]),W=Object(o.useRef)(null),A=Object(o.useCallback)(function(n){!function(n,e){n.strokeStyle=e.style,n.lineWidth=15,n.beginPath(),n.moveTo(e.points[0].x,e.points[0].y),e.points.forEach(function(e){n.lineTo(e.x,e.y)}),n.stroke()}(g.current,n)},[]),R=f()(function(n){return function(n){var e=T.current[T.current.length-1],t={x:Math.round(3240*(n.x-k.left)/k.width),y:Math.round(3240*(n.y-k.top)/k.height)};e.points.push(t),A(e)}(n)},16),U=(t=function(n){T.current.push({style:j,points:[]}),R(n)},a=function(n){R(n)},{start:function(n){b=!0,t(w(n))},move:function(n){b&&(n.stopPropagation(),a(w(n)))},stop:function(n){b=!1,r&&r(w(n))}}),N=[{icon:c.a.createElement(v.a,null),name:"color",hanler:function(){W.current.click(),M(!1)}}];return c.a.createElement("div",null,c.a.createElement("canvas",{className:e.root,ref:l,height:3240,width:3240,onMouseDown:U.start,onMouseMove:U.move,onMouseUp:U.stop,onTouchStart:U.start,onTouchMove:U.move,onTouchEnd:U.stop}),c.a.createElement(h.a,{ariaLabel:"SpeedDial",className:e.speedDial,icon:c.a.createElement(d.a,null),onClick:function(){return M(function(n){return!n})},onBlur:function(){return M(!1)},onClose:function(){return M(!1)},onFocus:function(){return M(!0)},onMouseEnter:function(){return M(!0)},onMouseLeave:function(){return M(!1)},open:x},N.map(function(n){return c.a.createElement(p.a,{key:n.name,icon:n.icon,tooltipTitle:n.name,onClick:n.hanler||console.log})})),c.a.createElement("input",{ref:W,type:"color",hidden:!0,onChange:function(n){return O(n.target.value)}}))},y=function(){return c.a.createElement(c.a.Fragment,null,c.a.createElement(i.a,null),c.a.createElement("div",{className:"App"},c.a.createElement(k,null)))},E=Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));function j(n,e){navigator.serviceWorker.register(n).then(function(n){n.onupdatefound=function(){var t=n.installing;null!=t&&(t.onstatechange=function(){"installed"===t.state&&(navigator.serviceWorker.controller?(console.log("New content is available and will be used when all tabs for this page are closed. See https://bit.ly/CRA-PWA."),e&&e.onUpdate&&e.onUpdate(n)):(console.log("Content is cached for offline use."),e&&e.onSuccess&&e.onSuccess(n)))})}}).catch(function(n){console.error("Error during service worker registration:",n)})}r.a.render(c.a.createElement(y,null),document.getElementById("root")),function(n){if("serviceWorker"in navigator){if(new URL("",window.location.href).origin!==window.location.origin)return;window.addEventListener("load",function(){var e="".concat("","/service-worker.js");E?(function(n,e){fetch(n).then(function(t){var o=t.headers.get("content-type");404===t.status||null!=o&&-1===o.indexOf("javascript")?navigator.serviceWorker.ready.then(function(n){n.unregister().then(function(){window.location.reload()})}):j(n,e)}).catch(function(){console.log("No internet connection found. App is running in offline mode.")})}(e,n),navigator.serviceWorker.ready.then(function(){console.log("This web app is being served cache-first by a service worker. To learn more, visit https://bit.ly/CRA-PWA")})):j(e,n)})}}({onSuccess:function(n){console.log("onSuccess",n)},onUpdate:function(n){console.log("onUpdate",n)}})}},[[43,1,2]]]);
//# sourceMappingURL=main.346ce105.chunk.js.map