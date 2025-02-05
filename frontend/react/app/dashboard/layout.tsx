import { Metadata } from 'next'
import { Sidebar } from '@/components/dashboard/sidebar'

export const metadata: Metadata = {
    title: '用户中心 - SSO',
    description: '管理您的个人信息、安全设置和应用授权',
}

interface DashboardLayoutProps {
    children: React.ReactNode
}

export default function DashboardLayout({ children }: DashboardLayoutProps) {
    return (
        <div className="flex min-h-screen">
            <Sidebar />
            <main className="flex-1 p-6">
                <div className="mx-auto max-w-4xl">{children}</div>
            </main>
        </div>
    )
}
