import { LayoutProps } from '@/models/common'
import * as React from 'react'

export function AuthLayout({ children }: LayoutProps) {
    return <section style={{background: '#e1f5fe', height: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center'}}>{children}</section>
}
