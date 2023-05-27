import * as React from 'react'
import Button from '@mui/material/Button'
import { Modal } from '@/components'
import { SearchField } from '@/components/room'
import PersonAddIcon from '@mui/icons-material/PersonAdd'
import { addItem } from '@/store';
import { useAppDispatch } from '@/store/configureStore';

export const SearchPopup = () => {
  const dispatch = useAppDispatch();

  const [isOk, setIsOk] = React.useState(false)

  const handleOk = (username: string) => {
    const data = {username: username, last_activity: Date.now() / 1000}
    dispatch(addItem(data))
    setIsOk(true)
  }

    return (
        <Modal
            openBtn={
                <Button>
                    <PersonAddIcon sx={{color: '#2196f3'}} />
                </Button>
            }
            body={<SearchField handleOk={handleOk} />}
            closeBtn={'Close'}
            isClickAwayClose={isOk}
        />

    )
}
