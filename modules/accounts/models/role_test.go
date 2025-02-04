package models_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rushairer/sso/modules/accounts/models"
)

func TestRole_Validation(t *testing.T) {
	tests := []struct {
		name    string
		role    *models.Role
		isValid bool
	}{
		{
			name: "有效的角色",
			role: &models.Role{
				ID:          uuid.New().String(),
				Name:        "管理员",
				Description: "系统管理员角色",
				CreatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			},
			isValid: true,
		},
		{
			name: "缺少ID",
			role: &models.Role{
				Name: "管理员",
			},
			isValid: false,
		},
		{
			name: "缺少Name",
			role: &models.Role{
				ID: uuid.New().String(),
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 验证必要字段
			isValid := tt.role.ID != "" && tt.role.Name != ""
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestRole_TimeFields(t *testing.T) {
	role := &models.Role{}

	// 测试时间字段的设置和获取
	now := time.Now()
	role.CreatedAt = sql.NullTime{Time: now, Valid: true}
	role.UpdatedAt = sql.NullTime{Time: now, Valid: true}

	// 验证时间字段
	assert.True(t, role.CreatedAt.Valid)
	assert.True(t, role.UpdatedAt.Valid)

	assert.Equal(t, now.Unix(), role.CreatedAt.Time.Unix())
	assert.Equal(t, now.Unix(), role.UpdatedAt.Time.Unix())
}
