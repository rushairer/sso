import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
    /* config options here */
    compiler: {
        styledComponents: {
            ssr: false,
        },
    },
    output: 'export',
}

export default nextConfig
