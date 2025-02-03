package models

import (
	"database/sql"
)

// Account 账户模型
type Account struct {
	ID              string // UUID
	Email           string
	Password        string
	Status          int8 // -1: 删除, 0: 禁用, 1: 正常
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	DeletedAt       sql.NullTime
	EmailVerifiedAt sql.NullTime
	LastLoginAt     sql.NullTime
}
