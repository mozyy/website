import { URLWs } from '../env'
import { Message, sendMessage } from './message'
import { encode, decode, decodeMessage } from './textEncoder'

const log = console.log.bind(console, '%c[[WebRTC]]: ', 'color: red')

interface Options {
  channel: string,
  uid?: number
}

export class RTC extends EventTarget {
  options :Options
  uid = 0
  webSocket: WebSocket | undefined
  rtcConnMaps = new Map<number,RTCPeerConnection>()
  dataChannelMaps = new Map<number,RTCDataChannel>()

  userlist: any

  constructor (options:Options = {channel:'testChannel'}) {
    super();
    this.options = options
  }

  async init() {
    this.webSocket = await this.createWebSocket()
    const joinChannel: Message = {
      kind: 'join',
      uid: this.uid,
      value: {
        channel: this.options.channel
      }
    };
  
    this.sendMessage(joinChannel)
  }

  private createWebSocket(){
    const webSocket = new WebSocket(URLWs)
    webSocket.addEventListener('close', e => {
      log('websocket closed')
      this.dispatchEvent(new CloseEvent('websocketclosed', {
        code: e.code, 
        reason: e.reason
      }))
    })
    webSocket.addEventListener('message', async e => {
      const msg = await decodeMessage(e)
      log('received websocket message: ', msg)
      this.dispatchEvent(e)
      const {kind, value, uid, target} = msg
      switch (kind) {
        case 'join-success':
          this.uid = target || 0
          break
        case 'userlist':      // Received an updated user list
          this.userlist = value
          break;
        // local message
        case 'user-joined':
          const remote = uid
          const connPeer = this.createRTCPeerConn(remote)
          const dataChannel = connPeer.createDataChannel(String(this.uid))
          this.createOffer(remote) // local
          this.createRTCDataChannel(remote, dataChannel)
          break
        case 'create-answer':
          this.receiveCreateAnswer(uid, value) // remote
          break;

        // remote message
        case 'create-offer':
          this.receiveCreateOffer(uid, value) // local remote
          break;

        case 'new-ice-candidate': // A new ICE candidate has been received
          this.receiveNewICECandidate(uid, value);
          break;
        case 'hang-up':
          this.closeRTCCall(this.rtcConnMaps.get(uid) as RTCPeerConnection)
          break;
        default:
          log('receive: Unknown message')
      }
    })
    return new Promise<WebSocket>((resolve, reject) => {
      webSocket.addEventListener('open', (e)=>{
        resolve(webSocket)
      })
      webSocket.addEventListener('error', (e)=>{
        reject(e)
      })
    })
  }
  private createRTCPeerConn(target: number) {
    const conn = new RTCPeerConnection({
      // Information about ICE servers - Use your own!
      iceServers: [{
          urls: [
            'stun:stun1.l.google.com:19302',
            'stun:stun2.l.google.com:19302',
            'stun:stun3.l.google.com:19302',
            'stun:stun4.l.google.com:19302',
            'stun:23.21.150.121',
            'stun:stun01.sipphone.com',
            'stun:stun.ekiga.net',
            'stun:stun.fwdnet.net',
            'stun:stun.ideasip.com',
            'stun:stun.iptel.org',
            'stun:stun.rixtelecom.se',
            'stun:stun.schlund.de',
            'stun:stunserver.org',
            'stun:stun.softjoys.com',
            'stun:stun.voiparound.com',
            'stun:stun.voipbuster.com',
            'stun:stun.voipstunt.com',
            'stun:stun.voxgratia.org',
            'stun:stun.xten.com',
          ]
        }]
    });
    conn.addEventListener('icecandidate', e => {
      if (e.candidate) {
        log("*** Outgoing ICE candidate: " + e.candidate.candidate);
        const msg: Message = {
          kind: 'new-ice-candidate',
          uid: this.uid,
          target,
          value: {
            candidate: e.candidate
          }
        }
        this.sendMessage(msg)
      }
    })
    conn.addEventListener('iceconnectionstatechange', () => {
      log("*** ICE connection state changed to " + conn.iceConnectionState);
      switch(conn.iceConnectionState) {
        case "closed":
        case "failed":
        case "disconnected":
          this.closeRTCCall(conn);
          break;
      }
    })
    conn.addEventListener('icegatheringstatechange', () => {
      log("*** ICE gathering state changed to: " + conn.iceGatheringState);
    })
    conn.addEventListener('signalingstatechange', () => {
      log("*** WebRTC signaling state changed to: " + conn.signalingState);
      switch(conn.signalingState) {
        case "closed":
          this.closeRTCCall(conn);
          break;
      }
    })
    conn.addEventListener('negotiationneeded', () => {
      this.createOffer(target)
    })
    conn.addEventListener('datachannel', event => {
      const dataChannel = event.channel
      this.createRTCDataChannel(Number(dataChannel.label), event.channel)
    })

    this.rtcConnMaps.set(target, conn)

    return conn
  }

  private createRTCDataChannel(target: number, channel: RTCDataChannel) {
    const logCurrentState = () => {
      console.log('current data channel state: ', channel.readyState)
    }
    channel.addEventListener('open', logCurrentState)
    channel.addEventListener('close', logCurrentState)
    channel.addEventListener('message', eventMessage => {
      // this.dispatchEvent(eventMessage)
      this.dispatchEvent(new CustomEvent('datachannelmessage', {detail: eventMessage}))
    })
    this.dataChannelMaps.set(target,channel)
    return channel
  }

  private async createOffer(target: number) {
    let conn = this.rtcConnMaps.get(target)
    if (!conn) {
      conn = this.createRTCPeerConn(target)
    }
    log('Creating offer')
    const offer = await conn.createOffer();
    if (conn.signalingState != "stable") {
      log("     -- The connection isn't stable yet; postponing...")
      return;
    }
    log("---> Setting local description to the offer");
    await conn.setLocalDescription(offer);
    log("---> Sending the offer to the remote peer");

    const msg: Message = {
      kind: 'create-offer',
      uid: this.uid,
      target,
      value: {
        sdp: conn.localDescription
      }
    }
    this.sendMessage(msg)
  }

  sendMessage(msg: Message) {
    if (this.webSocket) {
      this.webSocket.send(encode(msg))
    }
  }
  private closeRTCCall (peerConnection: RTCPeerConnection){
    if (!peerConnection) return
    // Close the RTCPeerConnection
    log("--> Closing the peer connection");
    // Disconnect all our event listeners; we don't want stray events
    // to interfere with the hangup while it's ongoing.
    peerConnection.ontrack = null;
    peerConnection.onicecandidate = null;
    peerConnection.oniceconnectionstatechange = null;
    peerConnection.onsignalingstatechange = null;
    peerConnection.onicegatheringstatechange = null;
    peerConnection.onnegotiationneeded = null;

    // Stop all transceivers on the connection
    peerConnection.getTransceivers().forEach(transceiver => {
      transceiver.stop();
    });

    // Stop the webcam preview as well by pausing the <video>
    // element, then stopping each of the getUserMedia() tracks
    // on it.

    // if (localVideo.srcObject) {
    //   localVideo.pause();
    //   (localVideo.srcObject as MediaStream ).getTracks().forEach(track => {
    //     track.stop();
    //   });
    // }

    // Close the peer connection

    peerConnection.close();
  }

  private async receiveCreateOffer(target: number, msg: any){
    let connPeer = this.rtcConnMaps.get(target)
    // If we're not already connected, create an RTCPeerConnection
    // to be linked to the caller.
    if (!connPeer) {
      connPeer = this.createRTCPeerConn(target)
    }
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
      ]).catch(error => log('setLocalDescription error: ', error));
      return
    } else {
      log ("  - Setting remote description");
      await connPeer.setRemoteDescription(desc);
    }
  
    log("---> Creating and sending answer to caller");
  
    await connPeer.setLocalDescription(await connPeer.createAnswer());
    const response = {
      kind: 'create-answer',
      uid: this.uid,
      target,
      value: {
        sdp: connPeer.localDescription
      }
    }
    this.sendMessage(response)
  }
  async receiveCreateAnswer(target: number, msg: any) {
    log("*** Call recipient has accepted our call");
  
    // Configure the remote description, which is the SDP payload
    // in our "video-answer" message.
  
    var desc = new RTCSessionDescription(msg.sdp);
    await (this.rtcConnMaps.get(target) as RTCPeerConnection).setRemoteDescription(desc)
  }

  async receiveNewICECandidate(target: number, msg: any) {
    var candidate = new RTCIceCandidate(msg.candidate);
  
    log("*** Adding received ICE candidate: " + JSON.stringify(candidate));
    await (this.rtcConnMaps.get(target) as RTCPeerConnection).addIceCandidate(candidate)
  }
  handleClose() {
    if (this.webSocket) {
      this.webSocket.close()
    }
    this.rtcConnMaps.forEach(connPeer => connPeer.close())
    this.dataChannelMaps.forEach(dataChannel => dataChannel.close())
  }
  dataChannelsSend(msg: Message) {
    const data = encode(msg)
    this.dataChannelMaps.forEach( dataChannel => {
      try{
        dataChannel.send(data)
      }catch(err) {
        console.log('send data channel message failed: ', err)
      }
    })
  }
}