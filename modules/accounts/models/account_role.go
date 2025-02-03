package models

import (
	"database/sql"
)

// AccountRole 账户角色关联
type AccountRole struct {
	AccountID string
	RoleID    string
	CreatedAt sql.NullTime
}
