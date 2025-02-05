import { Metadata } from 'next'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import Link from 'next/link'

export const metadata: Metadata = {
    title: '用户中心 - SSO',
    description: '管理您的个人信息和应用授权',
}

export default function DashboardPage() {
    return (
        <div className="space-y-6">
            <Card>
                <CardHeader>
                    <CardTitle>欢迎来到用户中心</CardTitle>
                </CardHeader>
                <CardContent className="space-y-4">
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <Link href="/dashboard/profile">
                            <Button
                                variant="outline"
                                className="w-full h-24 flex flex-col items-center justify-center space-y-2"
                            >
                                <span className="text-lg">个人信息</span>
                                <span className="text-sm text-muted-foreground">
                                    管理您的个人资料
                                </span>
                            </Button>
                        </Link>
                        <Link href="/dashboard/security">
                            <Button
                                variant="outline"
                                className="w-full h-24 flex flex-col items-center justify-center space-y-2"
                            >
                                <span className="text-lg">安全设置</span>
                                <span className="text-sm text-muted-foreground">
                                    管理账户安全选项
                                </span>
                            </Button>
                        </Link>
                        <Link href="/dashboard/applications">
                            <Button
                                variant="outline"
                                className="w-full h-24 flex flex-col items-center justify-center space-y-2"
                            >
                                <span className="text-lg">应用授权</span>
                                <span className="text-sm text-muted-foreground">
                                    管理第三方应用授权
                                </span>
                            </Button>
                        </Link>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
