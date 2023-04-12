import * as React from 'react'
import Button from '@mui/material/Button'
import SearchIcon from '@mui/icons-material/Search'
import { Modal } from '@/components'

export const SearchPopup = () => {
    return (
        <Modal
            openBtn={
                <Button>
                    <SearchIcon />
                </Button>
            }
            dialog={'Search user'}
            closeBtn={'Close'}
        />
    )
}
