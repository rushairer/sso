# SSO 系统功能流程图

---

## 目录

-   [关键组件说明](#关键组件说明)
-   [安全注意事项](#安全注意事项)
-   [SSO 单点登录流程图](#sso-单点登录流程图)
-   [注销流程](#注销流程)
-   [SSO 认证中心详细流程](#sso-认证中心详细流程)

---

### 关键组件说明

#### SSO 认证中心：

-   负责用户身份认证、会话管理、Token 颁发。

-   维护全局会话状态（如通过 Cookie 或分布式缓存）。

#### 客户端应用：

-   直接面向用户提供界面。
-   依赖 SSO 认证中心完成登录，不直接处理用户密码。

#### 资源服务器：

-   验证 Access Token 的合法性（如 JWT 签名、有效期）。
-   保存、处理和提供资源。

---

### 安全注意事项

-   Token 有效期：Access Token 应设置短有效期，使用 Refresh Token 续期。

-   HTTPS：所有通信需加密。

-   跨站防护：防范 CSRF 和 XSS 攻击（如 State 参数校验）。

---

### SSO 单点登录流程图

```mermaid
sequenceDiagram
    participant 用户
    participant 客户端应用
    participant SSO认证中心
    participant 资源服务器

    用户->>客户端应用: 1. 请求访问受保护资源
    客户端应用->>用户: 2. 重定向到SSO登录页面（携带redirect_uri）
    用户->>SSO认证中心: 3. 访问SSO登录页面
    SSO认证中心->>用户: 4. 展示登录页面
    用户->>SSO认证中心: 5. 输入用户名/密码
    SSO认证中心->>SSO认证中心: 6. 验证凭证，生成全局会话和授权码
    SSO认证中心->>用户: 7. 重定向回客户端应用（携带授权码）
    用户->>客户端应用: 8. 提交授权码
    客户端应用->>SSO认证中心: 9. 用授权码请求Access Token和ID Token
    SSO认证中心->>客户端应用: 10. 返回JWT格式的Access Token和ID Token
    客户端应用->>资源服务器: 11. 携带Access Token请求资源
    资源服务器->>资源服务器: 12. 本地验证JWT签名和有效期
    资源服务器->>客户端应用: 13. 返回受保护资源
    客户端应用->>用户: 14. 显示资源内容
```

### 分步详细说明

#### 步骤 1-2：用户触发登录

1. 用户请求资源：用户尝试访问客户端应用中的受保护资源。

2. 重定向到 SSO：客户端检测用户未登录，直接重定向到 SSO 认证中心，并附带 redirect_uri（回调地址）。

#### 步骤 3-7：用户登录并获取授权码

3. 访问登录页面：用户被重定向到 SSO 认证中心的登录页面。

4. 展示登录页面：SSO 认证中心返回登录页面。

5. 用户输入凭证：用户输入用户名和密码。

6. 验证凭证：SSO 认证中心验证凭证，生成全局会话和授权码。

7. 重定向回客户端：SSO 认证中心将用户重定向回客户端应用的 redirect_uri，并附带授权码。

#### 步骤 8-10：客户端获取 Token

8. 提交授权码：客户端应用通过后端通道接收授权码。

9. 请求 Token：客户端应用向 SSO 认证中心发送授权码，请求 Access Token 和 ID Token。

10. 返回 JWT Token：SSO 认证中心验证授权码后，返回 JWT 格式的 Access Token（签名）和 ID Token。

#### 步骤 11-14：访问资源

11. 携带 Token 请求资源：客户端应用在请求头中携带 Access Token（如 Authorization: Bearer <token>）。

12. 本地验证 JWT：资源服务器自行验证 JWT 的签名（使用预置公钥或动态获取 JWKS）和有效期（检查 exp 字段）。

13. 返回资源：验证通过后，资源服务器返回受保护数据。

14. 展示内容：客户端应用将资源展示给用户。

---

### 注销流程

```mermaid
sequenceDiagram
    participant 用户
    participant 客户端应用
    participant SSO认证中心

    用户->>客户端应用: 1. 点击注销
    客户端应用->>SSO认证中心: 2. 请求全局注销
    SSO认证中心->>SSO认证中心: 3. 销毁全局会话
    SSO认证中心->>所有关联应用: 4. 通知各客户端销毁本地会话
    SSO认证中心->>用户: 5. 重定向到登录页
```

---

### SSO 认证中心详细流程

```mermaid
sequenceDiagram
    participant 用户
    participant 客户端应用
    participant SSO认证中心
    participant 数据库
    participant Redis
    participant JWKS服务

    Note over SSO认证中心: 核心模块划分
    SSO认证中心->>SSO认证中心: 路由层（HTTP Server）
    SSO认证中心->>SSO认证中心: 业务逻辑层（Handler）
    SSO认证中心->>SSO认证中心: 数据层（DB/Redis）
    SSO认证中心->>SSO认证中心: 加密模块（JWT生成/验证）

    用户->>SSO认证中心: 1. GET /authorize?client_id=...&redirect_uri=...
    SSO认证中心->>数据库: 2. 查询客户端注册信息(client_id)
    数据库-->>SSO认证中心: 返回客户端配置
    SSO认证中心->>用户: 3. 返回登录页面（检查会话Cookie）
    alt 用户无有效会话
        用户->>SSO认证中心: 4. POST /login 提交用户名/密码
        SSO认证中心->>数据库: 5. 验证用户凭证
        数据库-->>SSO认证中心: 返回用户信息
        SSO认证中心->>Redis: 6. 创建全局会话(SessionID)
        SSO认证中心->>用户: 7. 设置会话Cookie，重定向到/authorize
    end
    SSO认证中心->>SSO认证中心: 8. 生成授权码(Authorization Code)
    SSO认证中心->>Redis: 9. 存储授权码(code→SessionID, client_id)
    SSO认证中心->>用户: 10. 重定向到redirect_uri?code=...
    用户->>客户端应用: 11. 提交授权码
    客户端应用->>SSO认证中心: 12. POST /token (code + client_secret)
    SSO认证中心->>Redis: 13. 验证授权码有效性
    Redis-->>SSO认证中心: 返回SessionID和client_id
    SSO认证中心->>加密模块: 14. 生成JWT (Access Token + ID Token)
    加密模块-->>SSO认证中心: 签名后的JWT
    SSO认证中心->>Redis: 15. 记录Token元数据（可选）
    SSO认证中心->>客户端应用: 16. 返回JSON {access_token, id_token}
    SSO认证中心->>JWKS服务: 17. 定期发布公钥（JWKS）

```
