import { sendMessage, Message } from './message'

// Called by the WebRTC layer to let us know when it's time to
// begin, resume, or restart ICE negotiation.
const log = console.log
const reportError = console.error

export const closeVideoCall = (peerConnection: RTCPeerConnection) => {
  var localVideo = document.getElementById("local_video") as HTMLVideoElement;

  log("Closing the call");

  // Close the RTCPeerConnection

  if (peerConnection) {
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

    if (localVideo.srcObject) {
      localVideo.pause();
      (localVideo.srcObject as MediaStream ).getTracks().forEach(track => {
        track.stop();
      });
    }

    // Close the peer connection

    peerConnection.close();
  }

}

const handleNegotiationNeededEvent = (peerConnection: RTCPeerConnection, conn: WebSocket) => async () => {
  log("*** Negotiation needed");

  try {
    log("---> Creating offer");
    const offer = await peerConnection.createOffer();

    // If the connection hasn't yet achieved the "stable" state,
    // return to the caller. Another negotiationneeded event
    // will be fired when the state stabilizes.

    if (peerConnection.signalingState != "stable") {
      log("     -- The connection isn't stable yet; postponing...")
      return;
    }

    // Establish the offer as the local peer's current
    // description.

    log("---> Setting local description to the offer");
    await peerConnection.setLocalDescription(offer);

    // Send the offer to the remote peer.

    log("---> Sending the offer to the remote peer");

    const msg: Message = {
      kind: 'video-offer',
      value: {
        sdp: peerConnection.localDescription
      }
    }
    sendMessage(conn, msg)
    // sendToServer({
    //   name: myUsername,
    //   target: targetUsername,
    //   type: "video-offer",
    //   sdp: peerConnection.localDescription
    // });
  } catch(err) {
    log("*** The following error occurred while handling the negotiationneeded event:");
    reportError(err);
  };
}

// Called by the WebRTC layer when events occur on the media tracks
// on our WebRTC call. This includes when streams are added to and
// removed from the call.
//
// track events include the following fields:
//
// RTCRtpReceiver       receiver
// MediaStreamTrack     track
// MediaStream[]        streams
// RTCRtpTransceiver    transceiver
//
// In our case, we're just taking the first stream found and attaching
// it to the <video> element for incoming media.

const handleTrackEvent = (event: RTCTrackEvent) => {
  log("*** Track event", event.streams[0]);
  // document.getElementById("received_video").srcObject = event.streams[0];
  // document.getElementById("hangup-button").disabled = false;
}

// Handles |icecandidate| events by forwarding the specified
// ICE candidate (created by our local ICE agent) to the other
// peer through the signaling server.

const handleICECandidateEvent = (conn: WebSocket) => (event: RTCPeerConnectionIceEvent) => {
  if (event.candidate) {
    log("*** Outgoing ICE candidate: " + event.candidate.candidate);


    const msg: Message = {
      kind: 'new-ice-candidate',
      value: {
        candidate: event.candidate
      }
    }
    sendMessage(conn, msg)
    // sendToServer({
    //   type: "new-ice-candidate",
    //   target: targetUsername,
    //   candidate: event.candidate
    // });
  }
}

// Handle |iceconnectionstatechange| events. This will detect
// when the ICE connection is closed, failed, or disconnected.
//
// This is called when the state of the ICE agent changes.

const handleICEConnectionStateChangeEvent = (peerConnection: RTCPeerConnection) => (event: Event) => {
  log("*** ICE connection state changed to " + peerConnection.iceConnectionState);

  switch(peerConnection.iceConnectionState) {
    case "closed":
    case "failed":
    case "disconnected":
      closeVideoCall(peerConnection);
      break;
  }
}

// Set up a |signalingstatechange| event handler. This will detect when
// the signaling connection is closed.
//
// NOTE: This will actually move to the new RTCPeerConnectionState enum
// returned in the property RTCPeerConnection.connectionState when
// browsers catch up with the latest version of the specification!

const handleSignalingStateChangeEvent = (peerConnection: RTCPeerConnection) => (event: Event) => {
  log("*** WebRTC signaling state changed to: " + peerConnection.signalingState);
  switch(peerConnection.signalingState) {
    case "closed":
      closeVideoCall(peerConnection);
      break;
  }
}

// Handle the |icegatheringstatechange| event. This lets us know what the
// ICE engine is currently working on: "new" means no networking has happened
// yet, "gathering" means the ICE engine is currently gathering candidates,
// and "complete" means gathering is complete. Note that the engine can
// alternate between "gathering" and "complete" repeatedly as needs and
// circumstances change.
//
// We don't need to do anything when this happens, but we log it to the
// console so you can see what's going on when playing with the sample.

const handleICEGatheringStateChangeEvent = (peerConnection: RTCPeerConnection) => (event: Event) => {
  log("*** ICE gathering state changed to: " + peerConnection.iceGatheringState);
}


const myHostname = window.location.hostname;



export const createPeerConnection = async (conn: WebSocket) => {

  // Create an RTCPeerConnection which knows to use our chosen
  // STUN server.

  const peerConnection = new RTCPeerConnection({
    iceServers: [     // Information about ICE servers - Use your own!
      {
        urls: "turn:" + myHostname,  // A TURN server
        username: "webrtc",
        credential: "turnserver"
      }
    ]
  });

  // Set up event handlers for the ICE negotiation process.

  peerConnection.onicecandidate = handleICECandidateEvent(conn);
  peerConnection.oniceconnectionstatechange = handleICEConnectionStateChangeEvent(peerConnection);
  peerConnection.onicegatheringstatechange = handleICEGatheringStateChangeEvent(peerConnection);
  peerConnection.onsignalingstatechange = handleSignalingStateChangeEvent(peerConnection);
  peerConnection.onnegotiationneeded = handleNegotiationNeededEvent(peerConnection, conn);
  peerConnection.ontrack = handleTrackEvent;

  return peerConnection
}