import { str2MillionSeconds, formatTimeStr } from "../../modules/utils";


export class Clock extends HTMLElement {

  constructor() {
    super()
    this.state = 'stop'
    this._changeEvent = new Event('change')
    console.dir(new.target)
  }

  connectedCallback() {
    this._valueTotal = str2MillionSeconds(this.getAttribute('value'))
    this._value = this._valueTotal
    this._valueDom.innerText = formatTimeStr(this._value)
  }

  start() {
    if (this.state === 'start') {
      return
    }
    this._timer = setInterval(()=>{
      if (this._value >= 0) {
        this.fitClock(formatTimeStr(this._value), this._value / this._valueTotal)
        this._value -= 1000
      } else {
        this.stop()
        this.dispatchEvent(this._changeEvent)
      }
    }, 1000)
    this.state = 'start'
  }
  stop() {
    if (this._timer) {
      clearInterval(this._timer)
    }
    this.state = 'stop'
    return this._value
  }
  fitClock() {
    console.error('请实现fitClock方法')
  }

}