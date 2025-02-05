'use client'

import { Card, CardContent } from '@/components/ui/card'

export default function PrivacyPage() {
    return (
        <div>
            <div className="mx-auto container py-10">
                <div className="mx-auto max-w-3xl">
                    <Card className="shadow-sm">
                        <CardContent className="p-6 pb-8 pt-8">
                            <h1 className="mb-8 text-3xl font-bold">
                                隐私政策
                            </h1>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    1. 信息收集
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们收集的信息包括但不限于：您的账户信息（如用户名、电子邮件地址）、身份验证信息、使用服务时产生的日志信息等。我们承诺只收集为提供服务所必需的信息。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    2. 信息使用
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们使用收集的信息用于：
                                    <br />- 提供、维护和改进我们的服务
                                    <br />- 验证用户身份和处理身份认证请求
                                    <br />- 发送服务相关通知
                                    <br />- 防范安全风险和滥用行为
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    3. 信息共享
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    除非经过您的明确同意，我们不会与第三方共享您的个人信息。但在以下情况下，我们可能会共享您的信息：
                                    <br />- 遵守法律法规的要求
                                    <br />- 保护我们的合法权益
                                    <br />- 经您授权的其他情况
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    4. 信息安全
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们采用业界标准的安全措施保护您的个人信息，包括但不限于加密存储、访问控制等技术手段。但请注意，互联网环境并非百分之百安全，我们建议您在使用服务时也要注意保护自己的个人信息。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    5. Cookie 使用
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们使用 Cookie
                                    和类似技术来提供和改进我们的服务。这些技术帮助我们了解用户如何使用我们的服务，从而进行优化和改进。您可以通过浏览器设置控制
                                    Cookie 的使用。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    6. 未成年人保护
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们非常重视对未成年人个人信息的保护。如果您是未成年人，请在监护人指导下使用我们的服务。如果我们发现自己在未获得可证实的父母同意的情况下收集了未成年人的个人信息，则会设法尽快删除相关数据。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    7. 政策更新
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们可能会不时更新本隐私政策。当我们进行重大更改时，我们会通过适当方式通知您，例如在我们的网站上发布通知。您继续使用我们的服务即表示您同意更新后的隐私政策。
                                </p>
                            </section>
                        </CardContent>
                    </Card>
                </div>
            </div>
            <footer className="mx-auto container mt-8 text-center text-sm text-muted-foreground my-10">
                <p>© 2024 SSO 统一身份认证平台. 保留所有权利。</p>
            </footer>
        </div>
    )
}
