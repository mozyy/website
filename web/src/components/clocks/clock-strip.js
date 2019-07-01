import { str2MillionSeconds, formatTimeStr } from "../../modules/utils";
import { Clock } from "./clock";

const width = 400;
const height = 100;
const lineWidth = 2;

const template = document.createElement('template')
template.innerHTML = `
<style>
  .box {
    position: relative;
    height: 200px;
  }
  .canvas-box {
    width: 100%;
    border-bottom: 4px solid rgb(127,21,21);
  }
  span {
    font-size: 20px;
  }
  canvas {
    display: block;
    width: 100%;
    transition: transform .2s;
  }
  i {
    position: absolute;
    left: 50%;
    transform: translate(-50%, calc(-50% + 2px));
    border: 10px solid white;
    border-top-color: transparent;
    border-left-color: transparent;
    border-right-color: transparent;
  }
</style>

<div class="box">
  <div class="canvas-box">
    <canvas width="${width}" height="${height}"></canvas>
  </div>
  <i></i>
</div>
`


class ClockStrip extends Clock {

  constructor() {
    super()
    this._changeEvent = new Event('change')
    this.attachShadow({mode:'open'})
    this.shadowRoot.appendChild(template.content.cloneNode(true))
    this._valueDom = this.shadowRoot.querySelector('span')
    this._circleDom = this.shadowRoot.querySelector('linear-circle')
  }

  connectedCallback() {
    super.connectedCallback()
    this.genCanvas()
  }

  genCanvas() {
    this._canvas = this.shadowRoot.querySelector('canvas')
    this._ctx = this._canvas.getContext('2d')
    const ctx = this._ctx

    ctx.beginPath()
    ctx.strokeStyle = 'white'
    ctx.fillStyle = 'white'
    ctx.font = '20px sans-serif'
    ctx.textAlign = 'center'
    ctx.lineWidth = lineWidth
    let total = Math.ceil(this._valueTotal / 60000)
    this._canvas.style.transform = `translateX(${total * 20 + 20}px)`
    console.log(this._canvas.style.transform)
    for(let i = 0; i <= total; i++) {
      const x = i * 20 + 20
      i % 5 || ctx.fillText(i, x, 55)
      ctx.moveTo(x, 70 - !(i%5) * 10)
      ctx.lineTo(x, height)
      ctx.stroke()
    }
  }

  fitClock(value, percent) {
    console.log(value, percent)

  }
}

customElements.define('clock-strip', ClockStrip)