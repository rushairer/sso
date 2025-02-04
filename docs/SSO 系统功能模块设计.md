# SSO 系统功能模块设计

## 功能描述

### 1. 帐户体系

#### 1.1 普通用户

##### 1.1.1 在平台注册帐户，可以通过以下方式：

-   邮箱注册
-   手机号注册
-   第三方平台帐户注册 \*
-   需要验证邮箱/手机号 \*
-   需要实名认证 \*

    `* 的内容为可配置。`

##### 1.1.2 在平台登录，可以通过以下方式：

-   邮箱登录
-   手机验证码登录
-   第三方平台帐户登录

##### 1.1.3 在平台修改个人信息：

-   修改邮箱，需要验证
-   修改手机号，需要验证
-   昵称等个性化信息

##### 1.1.4 用户状态:

-   基本状态：正常、禁用、已删除
-   是否已经验证邮箱/手机号 \*
-   是否已经进行实名认证 \*

    `* 的内容为可配置。`

##### 1.1.5 管理已经获取个人信息的应用。

#### 1.2 开发者用户

-   具有普通用户的所有功能。
-   具有管理应用的权限。

### 2. 应用体系
