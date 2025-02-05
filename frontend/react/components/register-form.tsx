import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

export function RegisterForm({
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
                                <h1 className="text-2xl font-bold">创建账户</h1>
                                <p className="text-balance text-muted-foreground">
                                    加入 SSO，开启您的个性化体验
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
                            <div className="grid gap-2">
                                <Label htmlFor="password">密码</Label>
                                <Input
                                    id="password"
                                    type="password"
                                    placeholder="至少8个字符"
                                    required
                                />
                            </div>
                            <div className="grid gap-2">
                                <Label htmlFor="confirmPassword">
                                    确认密码
                                </Label>
                                <Input
                                    id="confirmPassword"
                                    type="password"
                                    placeholder="再次输入密码"
                                    required
                                />
                            </div>
                            <Button type="submit" className="w-full">
                                注册
                            </Button>
                            <div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
                                <span className="relative z-10 bg-background px-2 text-muted-foreground">
                                    或使用以下方式注册
                                </span>
                            </div>
                            <div className="grid grid-cols-4 gap-4">
                                <Button variant="outline" className="w-full">
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            d="M8.691 2.188C3.891 2.188 0 5.476 0 9.53c0 2.212 1.17 4.203 3.002 5.55a.59.59 0 0 1 .213.665l-.39 1.48c-.019.07-.048.141-.048.213 0 .163.13.295.29.295a.326.326 0 0 0 .167-.054l1.903-1.114a.864.864 0 0 1 .717-.098 10.16 10.16 0 0 0 2.837.403c.276 0 .543-.027.811-.05-.857-2.578.157-4.972 1.932-6.446 1.703-1.415 3.882-1.98 5.853-1.838-.576-3.583-4.196-6.348-8.596-6.348zM5.785 5.991c.642 0 1.162.529 1.162 1.182a1.17 1.17 0 0 1-1.162 1.182A1.17 1.17 0 0 1 4.623 7.173c0-.653.52-1.182 1.162-1.182zm5.813 0c.642 0 1.162.529 1.162 1.182a1.17 1.17 0 0 1-1.162 1.182 1.17 1.17 0 0 1-1.162-1.182c0-.653.52-1.182 1.162-1.182zm5.34 2.99c-1.682-.054-3.38.47-4.725 1.603-1.345 1.132-2.084 2.72-2.084 4.283 0 3.272 3.058 5.948 6.809 5.948a7.84 7.84 0 0 0 2.188-.313.686.686 0 0 1 .569.077l1.509.884a.26.26 0 0 0 .132.043c.127 0 .231-.105.231-.234 0-.057-.023-.114-.038-.17l-.308-1.173a.476.476 0 0 1 .169-.527C22.674 17.614 24 15.585 24 13.877c0-3.223-2.973-5.896-7.062-5.896zm-2.625 3.066c.513 0 .929.422.929.943 0 .521-.416.943-.929.943a.936.936 0 0 1-.929-.943c0-.521.416-.943.929-.943zm5.25 0c.513 0 .929.422.929.943 0 .521-.416.943-.929.943a.936.936 0 0 1-.929-.943c0-.521.416-.943.929-.943z"
                                            fill="currentColor"
                                        />
                                    </svg>
                                    <span className="sr-only">
                                        使用微信账号注册
                                    </span>
                                </Button>
                                <Button variant="outline" className="w-full">
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.676-1.48 3.676-2.948 1.156-1.688 1.636-3.325 1.662-3.415-.039-.013-3.182-1.221-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.61 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.559-1.701"
                                            fill="currentColor"
                                        />
                                    </svg>
                                    <span className="sr-only">
                                        使用 Apple 账号注册
                                    </span>
                                </Button>
                                <Button variant="outline" className="w-full">
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
                                            fill="currentColor"
                                        />
                                    </svg>
                                    <span className="sr-only">
                                        使用 Google 账号注册
                                    </span>
                                </Button>
                                <Button variant="outline" className="w-full">
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z"
                                            fill="currentColor"
                                        />
                                    </svg>
                                    <span className="sr-only">
                                        使用 Github 账号注册
                                    </span>
                                </Button>
                            </div>
                            <div className="text-center text-sm">
                                已有账户？{' '}
                                <a
                                    href="/login"
                                    className="underline underline-offset-4"
                                >
                                    登录
                                </a>
                            </div>
                        </div>
                    </form>
                </CardContent>
            </Card>
            <div className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">
                点击继续，即表示您同意我们的 <a href="#">服务条款</a> 和{' '}
                <a href="#">隐私政策</a>。
            </div>
        </div>
    )
}
