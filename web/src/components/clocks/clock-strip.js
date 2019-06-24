import { str2MillionSeconds, formatTimeStr } from "../../modules/utils";

const string2dom = (str)=>{
  const container = document.createElement('div')
  container.innerHTML = str
  return container.firstChild
}

const template = document.createElement('template')
template.innerHTML = `
<style>
  div {
    display: flex;
    width: 200px;
    height: 200px;
    align-items: center;
    justify-content: center
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


class ClockStrip extends HTMLElement {

  constructor() {
    super()
    this._changeEvent = new Event('change')
    this.attachShadow({mode:'open'})
    this.shadowRoot.appendChild(template.content.cloneNode(true))
    this._valueDom = this.shadowRoot.querySelector('span')
    this._circleDom = this.shadowRoot.querySelector('linear-circle')
  }

  connectedCallback() {
    this._valueTotal = str2MillionSeconds(this.getAttribute('value'))
    this._value = this._valueTotal
    this._valueDom.innerText = formatTimeStr(this._value)
  }

  start() {
    this._timer = setInterval(()=>{
      if (this._value >= 0) {
        this._valueDom.innerText = formatTimeStr(this._value)
        this._circleDom.setAttribute('percent', this._value / this._valueTotal)
        this._value -= 1000
      } else {
        this.stop()
        this.dispatchEvent(this._changeEvent)
      }
    }, 1000)
  }
  stop() {
    if (this._timer) {
      clearInterval(this._timer)
    }
    return this._value
  }

}

customElements.define('clock-strip', ClockStrip)