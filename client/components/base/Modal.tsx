import * as React from 'react'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import Dialog from '@mui/material/Dialog'
import { theme } from '@/utils'

export interface IModalProps {
    openBtn: React.ReactNode
    head?: React.ReactNode
    body?: React.ReactNode
    closeBtn?: React.ReactNode
    okBtn?: React.ReactNode
    isClickAwayClose?: boolean
}

export const Modal = (props: IModalProps) => {
    const [open, setOpen] = React.useState(props.isClickAwayClose ? true : false)

    const handleClickListItem = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }

    // React.useEffect(() => {
    //   handleClose()
    // }, [props.isClickAwayClose])

    return (
        <Box>
            <div onClick={handleClickListItem}>{props.openBtn}</div>

            <Dialog
                sx={{ '& .MuiDialog-paper': { width: '80%', maxHeight: 435 } }}
                maxWidth="xs"
                open={open}
            >
                <DialogTitle>
                  {props.head && props.head}
                </DialogTitle>
                <DialogContent dividers={props.head ? true : false} sx={{overflow: 'hidden'}}>
                    {props.body}
                </DialogContent>
                  <DialogActions>
                    {props.closeBtn && <Button autoFocus onClick={handleClose} sx={{color: '#757575'}}>
                        {props.closeBtn}
                      </Button>
                    }
                    { props.okBtn &&
                      <Button sx={{display: 'flex', alignItems: 'center', color: '#2196f3'}}>{props.okBtn}</Button>
                    }
                  </DialogActions>
            </Dialog>
        </Box>
    )
}
