import * as React from 'react'
import IconButton from '@mui/material/IconButton'
import WidgetsIcon from '@mui/icons-material/Widgets'
import EmojiEmotionsIcon from '@mui/icons-material/EmojiEmotions'
import ThumbUpIcon from '@mui/icons-material/ThumbUp'
import { InputTemplate } from '@/components'

interface IStyledElement extends React.HTMLAttributes<HTMLElement> {
    style?: React.CSSProperties
    onData: (data: string) => void
}

export const ChatInput: React.FC<IStyledElement> = (
    props: React.PropsWithChildren<IStyledElement>
) => {
    const [input, setInput] = React.useState('')

    const handleSubmit = (e: any) => {
        setInput(e.target.value)
        if (e.key === 'Enter') {
            e.preventDefault()
            props.onData(input)
            setInput('')
        }
    }

    return (
        <>
            <IconButton sx={{ p: '0.5rem', mb: '0.25rem' }} aria-label="menu">
                <WidgetsIcon />
            </IconButton>

            <InputTemplate
                aria-label="Demo input"
                multiline
                placeholder="Aa"
                style={{ width: '100%', maxHeight: '8rem', marginLeft: '1rem' }}
                value={input}
                onChange={handleSubmit}
                onKeyDown={handleSubmit}
            />

            <IconButton type="button" sx={{ p: '0.5rem', mb: '0.25rem' }} aria-label="icon">
                <EmojiEmotionsIcon />
            </IconButton>

            <IconButton type="button" sx={{ p: '0.5rem', mb: '0.25rem' }} aria-label="like">
                <ThumbUpIcon />
            </IconButton>
        </>
    )
}
