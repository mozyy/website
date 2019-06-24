

const template = document.createElement('template')
template.innerHTML = `
<style>
  div {
    min-height: 100vh;
    // background: linear-gradient(to right, rgb(160,26,26), rgb(216,51,37), rgb(160,26,26))
  }

</style>

<div>
  <clock-round value="25:00" class="clock"></clock-round> 
  <button>start</button>
</div>
`


class HomeIndex extends HTMLElement {

  constructor() {
    super()
    this._startClock = this._startClock.bind(this)
    this.attachShadow({mode:'open'})
    this.shadowRoot.appendChild(template.content.cloneNode(true))
    this.clocks = Array.from(this.shadowRoot.querySelectorAll('.clock'))
    this.shadowRoot.querySelector('button').addEventListener('click', this._startClock)
  }

  connectedCallback() {
  }

  _startClock() {
    console.log(this)
    this.clocks.forEach(clock => {
      clock.start()
    })
  }


}

customElements.define('home-index', HomeIndex)