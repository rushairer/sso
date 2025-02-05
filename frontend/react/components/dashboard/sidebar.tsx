'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { Button } from '@/components/ui/button'
import { UserCircle, Shield, AppWindow } from 'lucide-react'

const sidebarItems = [
    {
        title: '个人信息',
        href: '/dashboard/profile',
        icon: UserCircle,
    },
    {
        title: '安全设置',
        href: '/dashboard/security',
        icon: Shield,
    },
    {
        title: '应用授权',
        href: '/dashboard/applications',
        icon: AppWindow,
    },
]

export function Sidebar() {
    const pathname = usePathname()

    return (
        <aside className="w-64 border-r bg-background p-6">
            <nav className="space-y-2">
                {sidebarItems.map((item) => (
                    <Link key={item.href} href={item.href}>
                        <Button
                            variant={
                                pathname === item.href ? 'secondary' : 'ghost'
                            }
                            className="w-full justify-start"
                        >
                            <item.icon className="mr-2 h-4 w-4" />
                            {item.title}
                        </Button>
                    </Link>
                ))}
            </nav>
        </aside>
    )
}
