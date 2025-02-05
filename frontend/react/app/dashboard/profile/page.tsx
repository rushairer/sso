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
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'

export const metadata: Metadata = {
    title: '个人信息 - SSO',
    description: '查看和编辑您的个人信息',
}

export default function ProfilePage() {
    return (
        <div className="space-y-6">
            <Card>
                <CardHeader>
                    <CardTitle>个人信息</CardTitle>
                    <CardDescription>查看和更新您的个人信息</CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="space-y-6">
                        <div className="flex items-center space-x-4">
                            <Avatar className="h-24 w-24">
                                <AvatarImage
                                    src="/placeholder-avatar.jpg"
                                    alt="用户头像"
                                />
                                <AvatarFallback>用户</AvatarFallback>
                            </Avatar>
                            <Button variant="outline">更换头像</Button>
                        </div>

                        <div className="grid gap-4 md:grid-cols-2">
                            <div className="space-y-2">
                                <Label htmlFor="username">用户名</Label>
                                <Input
                                    id="username"
                                    placeholder="请输入用户名"
                                />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="email">邮箱</Label>
                                <Input
                                    id="email"
                                    type="email"
                                    placeholder="请输入邮箱"
                                />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="phone">手机号码</Label>
                                <Input
                                    id="phone"
                                    type="tel"
                                    placeholder="请输入手机号码"
                                />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="nickname">昵称</Label>
                                <Input id="nickname" placeholder="请输入昵称" />
                            </div>
                        </div>

                        <div className="flex justify-end">
                            <Button>保存更改</Button>
                        </div>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
