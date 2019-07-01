import { str2MillionSeconds, formatTimeStr } from "../../modules/utils";
import { Clock } from "./clock";


const template = document.createElement('template')
template.innerHTML = `
<style>
  div {
    display: flex;
    width: 150px;
    height: 150px;
    align-items: center;
    justify-content: center;
  }
  span {
    position: absolute;
    font-size: 20px;
  }
</style>

<div>
  <linear-circle></linear-circle>
  <span></span>
</div>
`


class ClockRound extends Clock {

  constructor() {
    super()
    this.attachShadow({mode:'open'})
    this.shadowRoot.appendChild(template.content.cloneNode(true))
    this._valueDom = this.shadowRoot.querySelector('span')
    this._circleDom = this.shadowRoot.querySelector('linear-circle')
  }

  connectedCallback() {
    super.connectedCallback()
    this._valueDom.innerText = formatTimeStr(this._value)
  }

  fitClock(value, percent) {
    this._valueDom.innerText = value
    this._circleDom.setAttribute('percent', percent)
  }
}

customElements.define('clock-round', ClockRound)