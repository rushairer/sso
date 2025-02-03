package models

import (
	"database/sql"
)

// UserDetail 用户详情模型
type UserDetail struct {
	AccountID string // 关联Account.ID
	NickName  string
	Avatar    string
	Gender    int8 // 0: 未知, 1: 男, 2: 女
	Birthday  sql.NullTime
	Bio       string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
