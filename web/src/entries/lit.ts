import { LitElement, html, property, customElement } from 'lit-element';

@customElement('simple-greeting')
export class SimpleGreeting extends LitElement {
  @property() name = false;

  clickHandler() {
    console.log(this.name)
    // this.name = !this.name
  }

  render() {
    return html`<p>Hello, ${this.name}!</p><button @click="${this.clickHandler}">+1</button>`;
  }
}

window.sims= SimpleGreeting 

declare global {
  interface Window {
    sims: any
  }
}

window.addEventListener('load', () => {
  document.querySelector('#app').innerHTML = `
    <simple-greeting></simple-greeting>
  `
});