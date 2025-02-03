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

### MariaDB 创建 uuid 相关函数

```sql
DELIMITER //

    CREATE FUNCTION UuidToBin(_uuid BINARY(36))
        RETURNS BINARY(16)
        LANGUAGE SQL  DETERMINISTIC  CONTAINS SQL  SQL SECURITY INVOKER
    RETURN
        UNHEX(CONCAT(
            SUBSTR(_uuid, 15, 4),
            SUBSTR(_uuid, 10, 4),
            SUBSTR(_uuid,  1, 8),
            SUBSTR(_uuid, 20, 4),
            SUBSTR(_uuid, 25) ));
    //
    CREATE FUNCTION UuidFromBin(_bin BINARY(16))
        RETURNS BINARY(36)
        LANGUAGE SQL  DETERMINISTIC  CONTAINS SQL  SQL SECURITY INVOKER
    RETURN
        LCASE(CONCAT_WS('-',
            HEX(SUBSTR(_bin,  5, 4)),
            HEX(SUBSTR(_bin,  3, 2)),
            HEX(SUBSTR(_bin,  1, 2)),
            HEX(SUBSTR(_bin,  9, 2)),
            HEX(SUBSTR(_bin, 11))
                 ));

    //
    DELIMITER ;
```

### 热更新的方式启动 api

```bash
air -d -c .air/frontend/api.toml
```

# 文档

-   [功能设计](./docs/功能设计.md)
