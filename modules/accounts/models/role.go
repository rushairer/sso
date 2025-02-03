package models

import (
	"database/sql"
)

// Role 角色模型
type Role struct {
	ID          string
	Name        string
	Description string
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
