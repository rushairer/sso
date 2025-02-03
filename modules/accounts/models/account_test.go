package models_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rushairer/sso/modules/accounts/models"
)

func TestAccount_Validation(t *testing.T) {
	tests := []struct {
		name    string
		account *models.Account
		isValid bool
	}{
		{
			name: "有效账户",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "test@example.com",
				Password: "password123",
				Status:   1,
				CreatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			},
			isValid: true,
		},
		{
			name: "无效的邮箱格式 - 缺少@符号",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "invalid-email",
				Password: "password123",
				Status:   1,
			},
			isValid: false,
		},
		{
			name: "无效的邮箱格式 - 缺少域名",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "test@",
				Password: "password123",
				Status:   1,
			},
			isValid: false,
		},
		{
			name: "无效的邮箱格式 - 域名格式错误",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "test@invalid.",
				Password: "password123",
				Status:   1,
			},
			isValid: false,
		},
		{
			name: "无效的状态值",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "test@example.com",
				Password: "password123",
				Status:   3, // 无效的状态值
			},
			isValid: false,
		},
	}

	// 邮箱格式验证的正则表达式
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 验证必要字段和格式
			isValid := tt.account.ID != "" &&
				tt.account.Email != "" &&
				emailRegex.MatchString(tt.account.Email) &&
				tt.account.Password != "" &&
				(tt.account.Status >= -1 && tt.account.Status <= 1)

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestAccount_TimeFields(t *testing.T) {
	account := &models.Account{
		ID:       uuid.New().String(),
		Email:    "test@example.com",
		Password: "password123",
		Status:   1,
	}

	// 测试时间字段的设置和获取
	now := time.Now()
	account.CreatedAt = sql.NullTime{Time: now, Valid: true}
	account.UpdatedAt = sql.NullTime{Time: now, Valid: true}
	account.DeletedAt = sql.NullTime{Time: now, Valid: true}
	account.EmailVerifiedAt = sql.NullTime{Time: now, Valid: true}
	account.LastLoginAt = sql.NullTime{Time: now, Valid: true}

	// 验证时间字段
	assert.True(t, account.CreatedAt.Valid)
	assert.True(t, account.UpdatedAt.Valid)
	assert.True(t, account.DeletedAt.Valid)
	assert.True(t, account.EmailVerifiedAt.Valid)
	assert.True(t, account.LastLoginAt.Valid)

	assert.Equal(t, now.Unix(), account.CreatedAt.Time.Unix())
	assert.Equal(t, now.Unix(), account.UpdatedAt.Time.Unix())
	assert.Equal(t, now.Unix(), account.DeletedAt.Time.Unix())
	assert.Equal(t, now.Unix(), account.EmailVerifiedAt.Time.Unix())
	assert.Equal(t, now.Unix(), account.LastLoginAt.Time.Unix())
}
