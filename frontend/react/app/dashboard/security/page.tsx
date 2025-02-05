import { Metadata } from 'next'
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Switch } from '@/components/ui/switch'

export const metadata: Metadata = {
    title: '安全设置 - SSO',
    description: '管理您的账户安全设置',
}

export default function SecurityPage() {
    return (
        <div className="space-y-6">
            <Card>
                <CardHeader>
                    <CardTitle>修改密码</CardTitle>
                    <CardDescription>
                        定期更新密码可以提高账户安全性
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="space-y-4">
                        <div className="space-y-2">
                            <Label htmlFor="current-password">当前密码</Label>
                            <Input
                                id="current-password"
                                type="password"
                                placeholder="请输入当前密码"
                            />
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="new-password">新密码</Label>
                            <Input
                                id="new-password"
                                type="password"
                                placeholder="请输入新密码"
                            />
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="confirm-password">确认新密码</Label>
                            <Input
                                id="confirm-password"
                                type="password"
                                placeholder="请再次输入新密码"
                            />
                        </div>
                        <div className="flex justify-end">
                            <Button>更新密码</Button>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>两步验证</CardTitle>
                    <CardDescription>
                        启用两步验证以增加账户安全性
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="flex items-center justify-between">
                        <div className="space-y-1">
                            <div>Google Authenticator</div>
                            <div className="text-sm text-muted-foreground">
                                使用 Google Authenticator 进行两步验证
                            </div>
                        </div>
                        <Switch />
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>登录历史</CardTitle>
                    <CardDescription>查看您的账户登录记录</CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="text-sm text-muted-foreground">
                        最近的登录活动将显示在这里
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
