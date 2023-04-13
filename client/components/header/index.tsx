import * as React from 'react'
import { HeaderDesktop } from './HeaderDesktop'
import { HeaderMobile } from './HeaderMobile'

export function Header() {
    return (
        <>
            <HeaderMobile />
            <HeaderDesktop />
        </>
    )
}
