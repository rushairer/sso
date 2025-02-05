'use client'
import { Geist, Geist_Mono } from 'next/font/google'
import { ThemeProvider } from 'next-themes'
import { ThemeToggle } from '@/components/theme-toggle'
import { Logo } from '@/components/logo'
import { usePathname } from 'next/navigation'
import './globals.css'

const geistSans = Geist({
    variable: '--font-geist-sans',
    subsets: ['latin'],
})

const geistMono = Geist_Mono({
    variable: '--font-geist-mono',
    subsets: ['latin'],
})

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode
}>) {
    const pathname = usePathname()
    const noBackgroundPaths = ['/', '/login', '/register', '/reset-password']
    const shouldShowBackground = !noBackgroundPaths.includes(pathname)

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
                    {shouldShowBackground ? (
                        <>
                            <div
                                className={`fixed top-0 left-0 right-0 z-50 h-16 flex items-center justify-between px-8 backdrop-blur-md bg-background/90 border-b border-border/40 shadow-sm`}
                            >
                                <Logo />
                                <ThemeToggle />
                            </div>
                            <div className="pt-16">{children}</div>
                        </>
                    ) : (
                        <>
                            <Logo />
                            <ThemeToggle />
                            {children}
                        </>
                    )}
                </ThemeProvider>
            </body>
        </html>
    )
}
