import dynamic from 'next/dynamic'

interface CSRProps {
    path: string
}

export const CSR = ({ path }: CSRProps) => {
    const component: any = dynamic(() => import(path), {
        ssr: false,
    })
    return component
}
