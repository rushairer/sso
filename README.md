# SSO

一个基于 golang 的单点登录系统。

---

## 使用前的准备

### 安装 air

```bash
go install github.com/air-verse/air@latest
# export PATH=$PATH:/GO_PATH/bin

```

### 生成开发模式的自签名

```bash
openssl req -x509 -newkey rsa:4096 -keyout ./frontend/api/resources/dev.key.pem -out ./frontend/api/resources/dev.cert.pem -days 365 -nodes

openssl req -x509 -newkey rsa:4096 -keyout ./frontend/web/resources/dev.key.pem -out ./frontend/web/resources/dev.cert.pem -days 365 -nodes
```

### 热更新的方式启动 api

```bash
air -d -c .air/api
```

# 文档

-   [功能设计](./docs/功能设计.md)
