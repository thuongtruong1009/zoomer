import React, { useEffect, useRef } from "react";
import { useRouter } from "next/router";

const Room: React.FC = () => {
  const router = useRouter();
  const { roomID } = router.query;

  let userVideo = useRef<HTMLVideoElement>(null);
  let userStream = useRef<MediaStream | any>(null);
  let partnerVideo = useRef<HTMLVideoElement>(null);
  let peerRef = useRef<RTCPeerConnection | null>(null);
  let webSocketRef = useRef<WebSocket | null>(null);

  const openCamera = async () => {
    const allDevices = await navigator.mediaDevices.enumerateDevices();
    const cameras = allDevices.filter((device) => device.kind === "videoinput");

    const constraints = {
      audio: true,
      video: {
        deviceId: cameras[0].deviceId,
      },
    };

    try {
      return await navigator.mediaDevices.getUserMedia(constraints);
    } catch (err) {
      console.log(err);
      return null;
    }
  };

  // close camera
  // const closeCamera = () => {
  //   if (userStream.current) {
  //     userStream.current.getTracks().forEach((track: MediaStreamTrack) => {
  //       track.stop();
  //       userVideo.current && userVideo.current.srcObject && (userVideo.current.srcObject = null);
  //       userStream.current = null;
  //     });
  //   }
  // };

  // close socket
  // const closeSocket = () => {
  //   if (webSocketRef.current) {
  //     webSocketRef.current.close();
  //   }
  // };

  //muted audio
  const changeMuted = () => {
    if (userStream.current) {
      userStream.current.getAudioTracks()[0].enabled = false;
    }
  };

  //unmuted audio
  const changeUnMuted = () => {
    if (userStream.current) {
      userStream.current.getAudioTracks()[0].enabled = true;
    }
  }

  const callUser = () => {
    console.log("Calling Other User");
    if (userStream.current) {
      peerRef.current = createPeer();

      userStream.current.getTracks().forEach((track: MediaStreamTrack) => {
        if (peerRef.current) {
          peerRef.current.addTrack(track, userStream.current);
        }
      });
    }
  };

  const createPeer = () => {
    console.log("Creating Peer Connection");
    const peer = new RTCPeerConnection({
      iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
    });
    peer.onnegotiationneeded = handleNegotiationNeeded;
    peer.onicecandidate = handleIceCandidateEvent;
    peer.ontrack = handleTrackEvent;
    return peer;
  };

  const handleNegotiationNeeded = async () => {
    console.log("Create offer");
    if (peerRef.current) {
      try {
        const offer = await peerRef.current.createOffer();
        await peerRef.current.setLocalDescription(offer);
        if (webSocketRef.current) {
          webSocketRef.current.send(
            JSON.stringify({ offer: peerRef.current.localDescription })
          );
        }
      } catch (err) {
        console.log(err);
      }
    }
  };

  const handleIceCandidateEvent = (e: RTCPeerConnectionIceEvent) => {
    console.log("Found Ice Candidate");
    if (e.candidate && webSocketRef.current) {
      webSocketRef.current.send(JSON.stringify({ iceCandidate: e.candidate }));
    }
  };

  const handleTrackEvent = (e: RTCTrackEvent) => {
    console.log("Received Tracks");
    if (partnerVideo.current) {
      partnerVideo.current.srcObject = e.streams[0];
    }
  };

  const handleOffer = async (offer: RTCSessionDescriptionInit) => {
    console.log("Receiving and setting offer");
    peerRef.current = createPeer();

    if (peerRef.current && userStream.current) {
      await peerRef.current.setRemoteDescription(
        new RTCSessionDescription(offer)
      );

      userStream.current.getTracks().forEach((track: MediaStreamTrack) => {
        if (peerRef.current) {
          peerRef.current.addTrack(track, userStream.current);
        }
      });

      const answer = await peerRef.current.createAnswer();
      await peerRef.current.setLocalDescription(answer);

      if (webSocketRef.current) {
        webSocketRef.current.send(
          JSON.stringify({ answer: peerRef.current.localDescription })
        );
      }
    }
  };

  useEffect(() => {
    openCamera().then((stream) => {
      if (userVideo.current && stream) {
        userVideo.current.srcObject = stream;
        userStream.current = stream;
      }

      webSocketRef.current = new WebSocket(
        `ws://localhost:8081/join?streamID=${roomID}`
      );

      webSocketRef.current.addEventListener("open", () => {
        if (webSocketRef.current) {
          webSocketRef.current.send(JSON.stringify({ join: "true" }));
        }
      });

      webSocketRef.current.addEventListener("message", async (e) => {
        const message = JSON.parse(e.data);
        if (message.join) {
          callUser();
        }
        if (message.offer) {
          handleOffer(message.offer);
        }

        if (message.answer) {
          console.log("Receiving and setting answer");
          if (peerRef.current) {
            peerRef.current.setRemoteDescription(
              new RTCSessionDescription(message.answer)
            );
          }
        }
        if (message.iceCandidate) {
          console.log("Receiving and adding ICE Candidate");
          try {
            if (peerRef.current) {
              await peerRef.current.addIceCandidate(message.iceCandidate);
            }
          } catch (err) {
            console.log("Error receiving ICE Candidate", err);
          }
        }
      });
    });
  }, [roomID]);

  return (
    <div>
      {/* <video autoPlay controls ref={userVideo}></video>
      <video autoPlay controls ref={partnerVideo}></video> */}
      <video autoPlay ref={userVideo}></video>
      <video autoPlay ref={partnerVideo}></video>
      <button onClick={changeMuted}>turn of audio</button>
      <button onClick={changeUnMuted}>turn on audio</button>
      {/* <button onClick={closeCamera}>turnof camera</button>
      <button onClick={changeVideoOn}>turnon camera</button> */}
    </div>
  );
};

export default Room;
