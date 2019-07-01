

const template = document.createElement('template')
template.innerHTML = `
<style>
  div {
    min-height: 100vh;
    background: url(https://cn.bing.com/th?id=OHR.RedAnthiasCoralMayotte_EN-CN7102584078_480x800.jpg&rf=LaDigue_1920x1080.jpg&pid=hp) no-repeat center/cover;
    // background: linear-gradient(to right, rgb(160,26,26), rgb(216,51,37), rgb(160,26,26))
  }

</style>

<div>
  <clock-round value="25:00" class="clock"></clock-round> 
  <button id="start">start</button>
  <button id="stop">stop</button>
</div>
`


class HomeIndex extends HTMLElement {

  constructor() {
    super()
    this._startClock = this._startClock.bind(this)
    this._stopClock = this._stopClock.bind(this)
    this.attachShadow({mode:'open'})
    this.shadowRoot.appendChild(template.content.cloneNode(true))
    this.clocks = Array.from(this.shadowRoot.querySelectorAll('.clock'))
    this.shadowRoot.querySelector('#start').addEventListener('click', this._startClock)
    this.shadowRoot.querySelector('#stop').addEventListener('click', this._stopClock)
  }

  connectedCallback() {
  }

  _startClock() {
    this.clocks.forEach(clock => {
      clock.start()
    })
  }

  _stopClock() {
    this.clocks.forEach(clock => {
      clock.stop()
    })
  }


}

customElements.define('home-index', HomeIndex)