import { Metadata } from 'next'
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'

export const metadata: Metadata = {
    title: '应用授权 - SSO',
    description: '管理已授权的第三方应用',
}

export default function ApplicationsPage() {
    return (
        <div className="space-y-6">
            <Card>
                <CardHeader>
                    <CardTitle>已授权应用</CardTitle>
                    <CardDescription>
                        管理已授权的第三方应用访问权限
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="space-y-6">
                        <div className="flex items-center justify-between">
                            <div className="space-y-1">
                                <div className="font-medium">示例应用 A</div>
                                <div className="text-sm text-muted-foreground">
                                    访问您的基本信息和邮箱
                                </div>
                            </div>
                            <div className="flex items-center space-x-2">
                                <Switch defaultChecked />
                                <Button variant="ghost" size="sm">
                                    撤销授权
                                </Button>
                            </div>
                        </div>

                        <div className="flex items-center justify-between">
                            <div className="space-y-1">
                                <div className="font-medium">示例应用 B</div>
                                <div className="text-sm text-muted-foreground">
                                    访问您的基本信息
                                </div>
                            </div>
                            <div className="flex items-center space-x-2">
                                <Switch defaultChecked />
                                <Button variant="ghost" size="sm">
                                    撤销授权
                                </Button>
                            </div>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>授权记录</CardTitle>
                    <CardDescription>查看应用的授权历史记录</CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="text-sm text-muted-foreground">
                        最近的授权活动将显示在这里
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
