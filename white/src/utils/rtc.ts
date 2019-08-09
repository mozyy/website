import { URLWs } from '../env'

export const connect = () => {
  const connection = new WebSocket(URLWs)

}