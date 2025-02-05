# SSO

一个基于 golang 的单点登录系统。

---

## 使用前的准备

### 安装 air

```bash
go install github.com/air-verse/air@latest
# export PATH=$PATH:/GO_PATH/bin

```

### 热更新的方式启动 web

```bash
air -d -c .air/frontend/web.toml
```

### 热更新的方式启动 api

```bash
air -d -c .air/frontend/api.toml
```

# 文档

-   [功能设计](./docs/功能设计.md)
