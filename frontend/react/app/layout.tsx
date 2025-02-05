import type { Metadata } from 'next'
import { Geist, Geist_Mono } from 'next/font/google'
import { ThemeProvider } from 'next-themes'
import { ThemeToggle } from '@/components/theme-toggle'
import { Logo } from '@/components/logo'
import './globals.css'

const geistSans = Geist({
    variable: '--font-geist-sans',
    subsets: ['latin'],
})

const geistMono = Geist_Mono({
    variable: '--font-geist-mono',
    subsets: ['latin'],
})

export const metadata: Metadata = {
    title: 'SSO',
    description: '一个基于 golang 的单点登录系统',
}

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode
}>) {
    return (
        <html lang="en" suppressHydrationWarning>
            <body
                className={`${geistSans.variable} ${geistMono.variable} antialiased`}
                suppressHydrationWarning
            >
                <ThemeProvider
                    attribute="class"
                    defaultTheme="system"
                    enableSystem
                    disableTransitionOnChange
                >
                    <div className="fixed top-0 left-0 right-0 z-50 h-16 flex items-center justify-between px-8 backdrop-blur-md bg-background/90 border-b border-border/40 shadow-sm">
                        <Logo />
                        <ThemeToggle />
                    </div>
                    <div className="pt-16">{children}</div>
                </ThemeProvider>
            </body>
        </html>
    )
}
