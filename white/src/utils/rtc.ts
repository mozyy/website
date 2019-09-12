import { URLWs } from '../env'
import { encode, decode } from './textEncoder'
import { Message, sendMessage } from './message'
import { createPeerConnection, closeVideoCall } from './connection'

const log = console.log
const log_error = console.warn



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

const decodeMessage = (e: MessageEvent): Promise<Message>=> e.data.arrayBuffer().then(decode)

class Connection extends EventTarget {
  options :Options
  conn: undefined | WebSocket
  userlist = []
  connPeer: undefined | RTCPeerConnection
  uid = 0
  constructor (options:Options = {channel:'testChannel'}) {
    super();
    this.options = options
  }
  async init() {
    this.conn = await createSocket()
    this.addListener()
    await this.join()
    this.connPeer = await this.createPeerConnection()
  }
  join() {
    const joinChannel: Message = {
      kind: 'join',
      value: {
        channel: this.options.channel
      }
    };
    const conn = this.conn as WebSocket
  
    conn.send(encode(joinChannel))
    return new Promise<Message>((resolve, reject) => {
      const joinMessage = async (e: MessageEvent): Promise<Message> => {
        const message = await decodeMessage(e)
        switch (message.kind) {
          case 'join-success':
            this.dispatchEvent(new CustomEvent('join-success', {detail: message.value}))
            conn.removeEventListener('message', joinMessage)
            resolve(message)
          case 'join-failure':
            this.dispatchEvent(new CustomEvent('join-failure', {detail: message.value}))
            conn.removeEventListener('message', joinMessage)
            reject('join-failure')
          default:
            return message
        }
      }
      conn.addEventListener('message', joinMessage)
    })
  }

  createPeerConnection() {
    return createPeerConnection.call(this, this.conn as WebSocket)
  }

  addListener() {
    const conn = this.conn as WebSocket
    conn.addEventListener('close', (e)=>{
      console.log(e)
      this.dispatchEvent(e)
    })
    conn.addEventListener('message', async (e)=>{
      const msg = await decodeMessage(e)
      console.log('message: ', msg)
      this.dispatchEvent(e)
      const {kind, value} = msg
      switch (kind) {
        case 'join-success':
          this.uid = value.uid
          break
        case "userlist":      // Received an updated user list
          this.userlist = value
          break;
  
      // Signaling messages: these messages are used to trade WebRTC
      // signaling information during negotiations leading up to a video
      // call.
  
      case "video-offer":  // Invitation and offer to chat
        this.handleVideoOfferMsg(value);
        break;
  
      case "video-answer":  // Callee has answered our offer
        this.handleVideoAnswerMsg(msg);
        break;
  
      case "new-ice-candidate": // A new ICE candidate has been received
        this.handleNewICECandidateMsg(msg);
        break;
  
      case "hang-up": // The other peer has hung up the call
        this.handleHangUpMsg(msg);
        break;
  
      // Unknown message; output to console for debugging.
  
      default:
        log_error("Unknown message received:");
        log_error(msg);
      }
    })
  }

  handleVideoOfferMsg = async (msg: any) => {
    const target = msg.uid;
    const connPeer = this.connPeer as RTCPeerConnection
    // If we're not already connected, create an RTCPeerConnection
    // to be linked to the caller.
  
    log("Received video chat offer from " + target);
  
  
    // We need to set the remote description to the received SDP offer
    // so that our local WebRTC layer knows how to talk to the caller.
  
    var desc = new RTCSessionDescription(msg.sdp);
  
    // If the connection isn't stable yet, wait for it...
  
    if (connPeer.signalingState != "stable") {
      log("  - But the signaling state isn't stable, so triggering rollback");
  
      // Set the local and remove descriptions for rollback; don't proceed
      // until both return.
      await Promise.all([
        connPeer.setLocalDescription({type: "rollback"}),
        connPeer.setRemoteDescription(desc)
      ]);
      return;
    } else {
      log ("  - Setting remote description");
      await connPeer.setRemoteDescription(desc);
    }
  
    // Get the webcam stream if we don't already have it
  
    // if (!webcamStream) {
    //   try {
    //     webcamStream = await navigator.mediaDevices.getUserMedia(mediaConstraints);
    //   } catch(err) {
    //     handleGetUserMediaError(err);
    //     return;
    //   }
  
    //   document.getElementById("local_video").srcObject = webcamStream;
  
    //   // Add the camera stream to the RTCPeerConnection
  
    //   try {
    //     webcamStream.getTracks().forEach(
    //       transceiver = track => connPeer.addTransceiver(track, {streams: [webcamStream]})
    //     );
    //   } catch(err) {
    //     handleGetUserMediaError(err);
    //   }
    // }
  
    log("---> Creating and sending answer to caller");
  
    await connPeer.setLocalDescription(await connPeer.createAnswer());
    const response = {
      kind: 'video-answer',
      value: {
        uid: this.uid,
      }
    }
    sendMessage(this.conn as WebSocket, msg)
  
    // sendToServer({
    //   name: myUsername,
    //   target: targetUsername,
    //   type: "video-answer",
    //   sdp: connPeer.localDescription
    // });
  }
  async handleVideoAnswerMsg(msg: any) {
    log("*** Call recipient has accepted our call");
  
    // Configure the remote description, which is the SDP payload
    // in our "video-answer" message.
  
    var desc = new RTCSessionDescription(msg.sdp);
    await (this.connPeer as RTCPeerConnection).setRemoteDescription(desc)
  }

  async handleNewICECandidateMsg(msg: any) {
    var candidate = new RTCIceCandidate(msg.candidate);
  
    log("*** Adding received ICE candidate: " + JSON.stringify(candidate));
    await (this.connPeer as RTCPeerConnection).addIceCandidate(candidate)
  }
  
  handleHangUpMsg(msg: any) {
    closeVideoCall(this.connPeer as RTCPeerConnection)
  }
}

export const connect = async (options: Options = {channel:'testChannel'}) => {
  const conn = new Connection()
  
  await conn.init()
}