package models

import (
	"database/sql"
)

// AccountStatus 账户状态
type AccountStatus int8

const (
	AccountStatusDeleted  AccountStatus = -1 // 删除
	AccountStatusDisabled AccountStatus = 0  // 禁用
	AccountStatusNormal   AccountStatus = 1  // 正常
)

// Account 账户模型
type Account struct {
	ID              string // UUID
	Email           string
	Password        string
	Status          AccountStatus // 账户状态
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	DeletedAt       sql.NullTime
	EmailVerifiedAt sql.NullTime
	LastLoginAt     sql.NullTime
}

// IsValidStatus 检查账户状态是否有效
func (a *Account) IsValidStatus() bool {
	return a.Status >= AccountStatusDeleted && a.Status <= AccountStatusNormal
}
