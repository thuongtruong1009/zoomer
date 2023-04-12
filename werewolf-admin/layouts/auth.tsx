import { LayoutProps } from '@/models/common'
import * as React from 'react'

export function AuthLayout({ children }: LayoutProps) {
    return <section>{children}</section>
}
