import { str2MillionSeconds, formatTimeStr } from "../../modules/utils";
import { Clock } from "./clock";

const width = 400;
const height = 100;
const lineWidth = 4;

const template = document.createElement('template')
template.innerHTML = `
<style>
  div {
    width: 200px;
    height: 200px;
  }
  span {
    position: absolute;
    font-size: 20px;
  }
  canvas {
  }
</style>

<div>
  <canvas width="${width}" height="${height}"></canvas>
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
    console.log(this._ctx)
    const ctx = this._ctx
    
    ctx.lineWidth = lineWidth
    ctx.strokeStyle = 'rgb(127,21,21)'
    ctx.moveTo(0, height - lineWidth / 2)
    ctx.lineTo(width, height - lineWidth / 2)
    ctx.stroke()

    ctx.beginPath()
    ctx.strokeStyle = 'white'
    ctx.fillStyle = 'white'
    ctx.font = '20px sans-serif'
    ctx.textAlign = 'center'
    ctx.lineWidth = lineWidth / 2
    for(let i = 0; i <= 15; i++) {
      const x = i * 20 + 20
      i % 5 || ctx.fillText(i, x, 55)
      ctx.moveTo(x, 70 - !(i%5) * 10)
      ctx.lineTo(x, height - lineWidth)
      ctx.stroke()
    }
    
    
  }

}

customElements.define('clock-strip', ClockStrip)