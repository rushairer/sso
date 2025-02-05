'use client'

import { Card, CardContent } from '@/components/ui/card'

export default function TermsPage() {
    return (
        <div>
            <div className="container mx-auto py-10">
                <div className="mx-auto max-w-3xl">
                    <Card className="shadow-sm">
                        <CardContent className="p-6 pb-8 pt-8">
                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    1. 服务协议的范围
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    本服务条款是您与我们之间就SSO统一身份认证服务平台服务等相关事项所订立的契约。请您仔细阅读本服务条款，如果您对本服务条款的任何条款表示异议，您可以选择不使用我们的服务。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    2. 账户注册与安全
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    您在使用本服务前需要注册账户。您应当提供真实、准确、完整的个人资料，并在资料发生变更时及时更新。您应当妥善保管账户信息，对账户下的所有活动负责。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    3. 服务内容
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们提供统一身份认证服务，包括但不限于账户管理、身份验证、授权管理等功能。我们保留随时修改或中断服务的权利，且无需对任何第三方负责。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    4. 用户行为规范
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    您同意不从事以下行为：
                                    <br />- 违反法律法规或侵犯他人合法权益
                                    <br />- 干扰系统正常运行或未经授权访问系统
                                    <br />- 从事任何可能损害服务安全性的行为
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    5. 知识产权
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    本服务涉及的所有知识产权均归我们所有。未经我们明确书面许可，您不得使用、修改、复制、公开传播、改变、散布、发行或公开发表本服务的任何内容。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    6. 免责声明
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    在法律允许的最大范围内，我们对服务的及时性、安全性、准确性不作任何承诺，且不承担任何由于不可抗力、系统中断、数据丢失等情况给用户造成的损失。
                                </p>
                            </section>

                            <section className="mb-6">
                                <h2 className="mb-4 text-xl font-semibold">
                                    7. 协议修改
                                </h2>
                                <p className="mb-4 text-muted-foreground">
                                    我们保留随时修改本服务条款的权利。条款修改后，我们会在网站公示。如您继续使用本服务，则视为您接受修改后的条款。
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
