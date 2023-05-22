import * as React from 'react'
import Button from '@mui/material/Button'
import { Modal } from '@/components'
import { SearchField } from '@/components/room'
import GroupAddIcon from '@mui/icons-material/GroupAdd';
import PersonAddIcon from '@mui/icons-material/PersonAdd'
import { useSelector, useDispatch } from 'react-redux';
import { addItem } from '@/store';
import { RootState } from '@/store/configureStore';

export const SearchPopup = () => {
  // const items = useSelector((state: RootState) => state.items);
  const dispatch = useDispatch();

  const [isOk, setIsOk] = React.useState(false)

  const handleOk = (username: string) => {
    const data = {username: username, last_activity: Date.now() / 1000}
    dispatch(addItem(data))
    console.log(data)
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
