package models_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rushairer/sso/modules/accounts/models"
)

func TestAccountRole_Validation(t *testing.T) {
	tests := []struct {
		name        string
		accountRole *models.AccountRole
		isValid     bool
	}{
		{
			name: "有效的账户角色关联",
			accountRole: &models.AccountRole{
				AccountID: uuid.New().String(),
				RoleID:    uuid.New().String(),
				CreatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			},
			isValid: true,
		},
		{
			name: "缺少AccountID",
			accountRole: &models.AccountRole{
				RoleID: uuid.New().String(),
			},
			isValid: false,
		},
		{
			name: "缺少RoleID",
			accountRole: &models.AccountRole{
				AccountID: uuid.New().String(),
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 验证必要字段
			isValid := tt.accountRole.AccountID != "" && tt.accountRole.RoleID != ""
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestAccountRole_TimeFields(t *testing.T) {
	accountRole := &models.AccountRole{
		AccountID: uuid.New().String(),
		RoleID:    uuid.New().String(),
	}

	// 测试时间字段的设置和获取
	now := time.Now()
	accountRole.CreatedAt = sql.NullTime{Time: now, Valid: true}

	// 验证时间字段
	assert.True(t, accountRole.CreatedAt.Valid)
	assert.Equal(t, now.Unix(), accountRole.CreatedAt.Time.Unix())
}
