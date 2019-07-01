const width = 500;
const strokeWidth = 10;


const center = width / 2;
const radius = (width - strokeWidth) / 2;
const perimeter = Math.ceil(2 * Math.PI * radius)

const template = document.createElement('template')
template.innerHTML = `
<style>
  svg {
    width: 100%;
    height: 100%;
    transform: rotate(-0.05deg);
  }
  circle {
    transition: stroke-dasharray .2s;
  }
</style>

<svg viewBox="0 0 ${width} ${width}" class="center">
  <defs>
    <linearGradient x1="1" y1="0" x2="0" y2="0" id="gradient1">
      <stop offset="0%" stop-color="#e52c5c"></stop>
      <stop offset="100%" stop-color="#ab5aea"></stop>
    </linearGradient>
    <linearGradient x1="1" y1="0" x2="0" y2="0" id="gradient2">
      <stop offset="0%" stop-color="#4352f3"></stop>
      <stop offset="100%" stop-color="#ab5aea"></stop>
    </linearGradient>
  </defs>
  <g transform="matrix(0,-1,1,0,0,${width})">
      <circle cx="${center}" cy="${center}" r="${radius}" stroke-width="${strokeWidth}" stroke="#f0f1f5" fill="none" stroke-dasharray="${perimeter} ${perimeter}"></circle>
      <circle cx="${center}" cy="${center}" r="${radius}" stroke-width="${strokeWidth}" stroke="url('#gradient1')" fill="none" stroke-dasharray="${perimeter} ${perimeter}"></circle>
      <circle cx="${center}" cy="${center}" r="${radius}" stroke-width="${strokeWidth}" stroke="url('#gradient2')" fill="none" stroke-dasharray="${perimeter/2} ${perimeter}"></circle>
  </g>
</svg>
`

class LinearCircle extends HTMLElement {

  static get observedAttributes() {
    return ['percent']
  }

  constructor () {
    super()
    this.attachShadow({mode: 'open'})
    this.shadowRoot.appendChild(template.content.cloneNode(true))
    this.circles = this.shadowRoot.querySelectorAll('circle')
  }

  updatePercent(percent) {
    if (isNaN(Number(percent)) || percent > 1 ) {
      console.log('请输入正确比例')
      return
    }
    const value = percent*perimeter
    this.circles[1].setAttribute('stroke-dasharray', `${value} ${perimeter}`)
    if (percent < 0.5) {
      this.circles[2].setAttribute('stroke-dasharray', `${value} ${perimeter}`)
    }
  }


  attributeChangedCallback(name, oldValue, newValue) {
    switch(name) {
      case 'percent': 
        this.updatePercent(newValue)
        break;
    }
  }

  set percent (value) {
    this.setAttribute('percent', value)
  }
}

customElements.define('linear-circle', LinearCircle)
