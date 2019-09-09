import { URLWs } from '../env'
import { encode, decode } from './textEncoder'
import { Message } from './message'


export const createSocket = () => new Promise<WebSocket>((resolve, reject) => {

  const connection = new WebSocket(URLWs)

  connection.addEventListener('open', (e)=>{
    resolve(connection)
  })
  connection.addEventListener('error', (e)=>{
    reject(e)
  })
})

interface Options {
  channel: string,
  uid?: number
}

export const connect = async (options: Options = {channel:'testChannel'}) => {
  const connection = await createSocket()
  connection.addEventListener('close', (e)=>{
    console.log(e)
  })
  connection.addEventListener('message', async(e)=>{
    const resp = await e.data.text()
    console.log(decode(resp))
  })
  const joinChannel: Message = {
    kind: 'join',
    value: {
      channel: options.channel
    }
  }

  connection.send(encode(joinChannel))
}