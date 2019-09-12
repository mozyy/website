
import { encode } from './textEncoder'

type requestKinds = 'join' | 'leave' | 'video-offer' | 'new-ice-candidate' | string // TODO: add every kinds
type responseKinds = 'join-success'|'join-failure'

export interface Message {
  kind: requestKinds | responseKinds;
  value: any
}

export const sendMessage = (conn: WebSocket, msg: Message) => {
  conn.send(encode(msg))
}