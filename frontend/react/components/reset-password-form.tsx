import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

export function ResetPasswordForm({
    className,
    ...props
}: React.ComponentProps<'div'>) {
    return (
        <div
            className={cn('flex flex-col gap-6 max-w-md mx-auto', className)}
            {...props}
        >
            <Card className="overflow-hidden">
                <CardContent className="p-0">
                    <form className="p-6">
                        <div className="flex flex-col gap-6">
                            <div className="flex flex-col items-center text-center">
                                <h1 className="text-2xl font-bold">找回密码</h1>
                                <p className="text-balance text-muted-foreground">
                                    输入您的邮箱，我们将向您发送重置密码的链接
                                </p>
                            </div>
                            <div className="grid gap-2">
                                <Label htmlFor="email">邮箱</Label>
                                <Input
                                    id="email"
                                    type="email"
                                    placeholder="me@example.com"
                                    required
                                />
                            </div>
                            <Button type="submit" className="w-full">
                                发送重置链接
                            </Button>
                            <div className="text-center text-sm">
                                记起密码了？{' '}
                                <a
                                    href="/login"
                                    className="underline underline-offset-4"
                                >
                                    返回登录
                                </a>
                            </div>
                        </div>
                    </form>
                </CardContent>
            </Card>
        </div>
    )
}
