import { MainLayout } from '@/layouts';
import { NextPageWithLayout } from '@/models';
import { Box, Button, Card, ClickAwayListener, Grow, MenuItem, MenuList, Paper, Popper, Stack, SxProps, Theme } from '@mui/material';
import { useEffect, useRef, useState } from 'react';
import DnsIcon from '@mui/icons-material/Dns';
import CallEndIcon from '@mui/icons-material/CallEnd';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
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

const BasicCard: NextPageWithLayout = () => {
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
  }

  const [isVideoOn, setIsVideoOn] = useState(true);
  const changeVideoOn = () => {
    setIsVideoOn(!isVideoOn);
  }

  const findAndClick = (selector: string) => {
    const element = document.querySelector(selector) as HTMLElement;
    if (element) {
      element.click();
    }
  }

  useEffect(() => {
    findAndClick('[aria-label="Turn off microphone (⌘ + D)"]');
    findAndClick('[aria-label="Turn off camera (⌘ + E)"]');
  }, [isMuted, isVideoOn])

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
    <div style={{display: 'flex', flexDirection: 'column', alignItems: 'center',  background: '#E7F3FF', padding: '1rem'}}>
      <Box
        component="ul"
        sx={containerStyle}
      >
        <Card component="li" sx={videoBigStyle}>
            <video
              autoPlay
              // controls
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
              // controls
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
        <Item><CallEndIcon /></Item>
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
}

BasicCard.Layout = MainLayout

export default BasicCard
