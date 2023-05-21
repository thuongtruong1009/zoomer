import React, { useEffect, useRef, useState } from "react";
import { useRouter } from "next/router";
import { Box, Button, Card, ClickAwayListener, Grow, MenuItem, MenuList, Paper, Popper, Stack, SxProps, Theme } from '@mui/material';
import DnsIcon from '@mui/icons-material/Dns';
import CallEndIcon from '@mui/icons-material/CallEnd';
import AutoAwesomeIcon from '@mui/icons-material/AutoAwesome';
import ViewCarouselIcon from '@mui/icons-material/ViewCarousel';
import LaptopChromebookIcon from '@mui/icons-material/LaptopChromebook';
import VideocamIcon from '@mui/icons-material/Videocam';
import VideocamOffIcon from '@mui/icons-material/VideocamOff';
import VolumeUpIcon from '@mui/icons-material/VolumeUp';
import VolumeOffIcon from '@mui/icons-material/VolumeOff';
import AutoAwesomeMosaicIcon from '@mui/icons-material/AutoAwesomeMosaic';

enum EVIEW_MODE {
  GRID = 'grid',
  NEST = 'nest',
  FLEX = 'flex'
}

import { styled } from '@mui/material/styles';

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  height: 'fit-content',
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
  borderRadius: '50%'
}));

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
      // video: cameraOn ? { deviceId: cameras[0].deviceId } : false,
    };

    try {
      return await navigator.mediaDevices.getUserMedia(constraints);
    } catch (err) {
      console.log(err);
      return null;
    }
  };

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
        userVideo.current.srcObject = stream
        userStream.current = stream
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

  const [viewMode, setViewMode] = useState<EVIEW_MODE>(EVIEW_MODE.NEST);
  const handleViewMode = (event: React.SyntheticEvent, newValue: EVIEW_MODE) => {
    setViewMode(newValue);
    handleClose(event);
  };

  const [containerStyle, setContainerStyle] = useState<SxProps<Theme> | undefined>(undefined);
  const [videoBigStyle, setVideoBigStyle] = useState<SxProps<Theme> | undefined>(undefined);
  const [videoSmallStyle, setVideoSmallStyle] = useState<SxProps<Theme> | undefined>(undefined);

  useEffect(() => {
    if(viewMode === EVIEW_MODE.FLEX) {
      setContainerStyle({
        display: 'grid',
        gridTemplateColumns: '1fr 1fr',
        gap: '1rem',
        width: '100%',
        transition: '1s ease-in-out'
      })
      setVideoBigStyle({
        borderRadius: '0.75rem',
        boxShadow: '0 0 0.6rem rgba(0, 0, 0, 0.55)',
      })
      setVideoSmallStyle({
        borderRadius: '0.75rem',
        boxShadow: '0 0 0.6rem rgba(0, 0, 0, 0.55)',
      })
    }
    if(viewMode === EVIEW_MODE.GRID) {
      setContainerStyle({
        display: 'grid',
        gridTemplateRows: 'auto auto',
        justifyContent: 'center',
        transition: '1s ease-in-out',
        gap: '1rem',
        width: '100%',
      })
      setVideoBigStyle({
        borderRadius: '1rem',
        boxShadow: '0 0 0.6rem rgba(0, 0, 0, 0.55)',
        width: '100%',
      })
      setVideoSmallStyle({
        borderRadius: '1rem',
        boxShadow: '0 0 0.6rem rgba(0, 0, 0, 0.55)',
        width: '100%',
      })
    }
    if(viewMode === EVIEW_MODE.NEST) {
      setContainerStyle({
        position: 'relative',
        transition: '1s ease-in-out',
      })
      setVideoBigStyle({
        position: 'relative',
        zIndex: 1,
        borderRadius: '1rem',
      })
      setVideoSmallStyle({
        position: 'absolute',
        bottom: '1rem',
        right: '1rem',
        zIndex: 2,
        width: '20%',
        height: '20%',
        borderRadius: '0.5rem',
      })
    }
  }, [viewMode])

  // menu
  const [open, setOpen] = useState(false);
  const anchorRef = useRef<HTMLButtonElement>(null);

  const handleToggle = () => {
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClose = (event: Event | React.SyntheticEvent) => {
    if (
      anchorRef.current &&
      anchorRef.current.contains(event.target as HTMLElement)
    ) {
      return;
    }
    setOpen(false);
  };

  function handleListKeyDown(event: React.KeyboardEvent) {
    if (event.key === 'Tab') {
      event.preventDefault();
      setOpen(false);
    } else if (event.key === 'Escape') {
      setOpen(false);
    }
  }

  // return focus to the button when we transitioned from !open -> open
  const prevOpen = useRef(open);
  useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current!.focus();
    }
    prevOpen.current = open;
  }, [open]);

  //logic video
  const [isMuted, setIsMuted] = useState(false);

  const changeMuted = () => {
    setIsMuted(!isMuted);
    userStream.current.getAudioTracks()[0].enabled = isMuted;
  }

  const [isVideoOn, setIsVideoOn] = useState(true);
  const changeVideoOn = () => {
    setIsVideoOn(!isVideoOn);
  }

  const endCall = () => {
    if (webSocketRef.current) {
      // await webSocketRef.current.send(JSON.stringify({ leave: "true" }));
      webSocketRef.current.close();
      window.close()
    }
  }

  const checkMatchMode  =  (mode: string) => {
    if(viewMode === mode) {
      return {
        color: '#3f51b5',
        fontWeight: 'bold',
        backgroundColor: '#E7F3FF'
      }
    }
  }

  return (
    // <div>
    //   <video autoPlay ref={userVideo}></video>
    //   <video autoPlay ref={partnerVideo}></video>

    //   <button onClick={changeMuted}>turn of audio</button>
    //   <button onClick={changeUnMuted}>turn on audio</button>
    // </div>
    <div style={{display: 'flex', flexDirection: 'column', alignItems: 'center',  background: '#E7F3FF', padding: '1rem'}}>
      <Box
        component="ul"
        sx={containerStyle}
      >
        <Card component="li" sx={videoBigStyle}>
            <video
              autoPlay
              ref={userVideo}
              loop
              poster="https://assets.codepen.io/6093409/river.jpg"
            >
              <source
                src="https://assets.codepen.io/6093409/river.mp4"
                type="video/mp4"
              />
            </video>
        </Card>
        <Card component="li" sx={videoSmallStyle}>
            <video
              autoPlay
              ref={partnerVideo}
              loop
              poster="https://assets.codepen.io/6093409/river.jpg"
            >
              <source
                src="https://assets.codepen.io/6093409/river.mp4"
                type="video/mp4"
              />
            </video>
        </Card>
      </Box>

      <Stack direction={{ xs: 'column', sm: 'row' }} spacing={{ xs: 1, sm: 2, md: 4 }} marginTop={2}>
        <Item onClick={endCall}><CallEndIcon /></Item>
        {
          isVideoOn ? <Item><VideocamIcon onClick={changeVideoOn} /></Item> : <Item><VideocamOffIcon onClick={changeVideoOn} /></Item>
        }
        {
          isMuted ? <Item><VolumeOffIcon onClick={changeMuted} /></Item> : <Item><VolumeUpIcon onClick={changeMuted} /></Item>
        }
        <Item><AutoAwesomeIcon /></Item>
        <Button
          ref={anchorRef}
          id="composition-button"
          aria-controls={open ? 'composition-menu' : undefined}
          aria-expanded={open ? 'true' : undefined}
          aria-haspopup="true"
          onClick={handleToggle}
          aria-label='more'
          sx={{m: 0, p: 0}}
        >
          <Item><AutoAwesomeMosaicIcon /></Item>
        </Button>
        <Popper
          open={open}
          anchorEl={anchorRef.current}
          role={undefined}
          placement="bottom-start"
          transition
          disablePortal
          style={{zIndex: 3}}
        >
          {({ TransitionProps, placement }) => (
            <Grow
              {...TransitionProps}
              style={{
                transformOrigin:
                  placement === 'bottom-start' ? 'left top' : 'left bottom',
              }}
            >
              <Paper>
                <ClickAwayListener onClickAway={handleClose}>
                  <MenuList
                    autoFocusItem={open}
                    id="composition-menu"
                    aria-labelledby="composition-button"
                    onKeyDown={handleListKeyDown}
                  >
                    <MenuItem onClick={(e) => handleViewMode(e, EVIEW_MODE.FLEX)} sx={checkMatchMode(EVIEW_MODE.FLEX)}><ViewCarouselIcon sx={{mr: 1}} /><h5>{EVIEW_MODE.FLEX}</h5></MenuItem>
                    <MenuItem onClick={(e) => handleViewMode(e, EVIEW_MODE.NEST)} sx={checkMatchMode(EVIEW_MODE.NEST)}><LaptopChromebookIcon sx={{mr: 1}} /> {EVIEW_MODE.NEST}</MenuItem>
                    <MenuItem onClick={(e) => handleViewMode(e, EVIEW_MODE.GRID)} sx={checkMatchMode(EVIEW_MODE.GRID)}><DnsIcon sx={{mr: 1}} /> {EVIEW_MODE.GRID}</MenuItem>
                  </MenuList>
                </ClickAwayListener>
              </Paper>
            </Grow>
          )}
        </Popper>
      </Stack>
    </div>
  );
};

export default Room;
