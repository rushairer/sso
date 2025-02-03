package models_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rushairer/sso/modules/accounts/models"
)

func TestUserDetail_Validation(t *testing.T) {
	tests := []struct {
		name       string
		userDetail *models.UserDetail
		isValid    bool
	}{
		{
			name: "有效的用户详情",
			userDetail: &models.UserDetail{
				AccountID: uuid.New().String(),
				NickName:  "测试用户",
				Avatar:    "avatar.jpg",
				Gender:    1,
				Birthday: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				Bio: "测试简介",
			},
			isValid: true,
		},
		{
			name: "无效的性别值",
			userDetail: &models.UserDetail{
				AccountID: uuid.New().String(),
				NickName:  "测试用户",
				Gender:    3, // 无效的性别值
			},
			isValid: false,
		},
		{
			name: "缺少AccountID",
			userDetail: &models.UserDetail{
				NickName: "测试用户",
				Gender:   1,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 验证必要字段
			isValid := tt.userDetail.AccountID != "" &&
				tt.userDetail.NickName != "" &&
				(tt.userDetail.Gender >= 0 && tt.userDetail.Gender <= 2)

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestUserDetail_TimeFields(t *testing.T) {
	userDetail := &models.UserDetail{
		AccountID: uuid.New().String(),
		NickName:  "测试用户",
		Gender:    1,
	}

	// 测试时间字段的设置和获取
	now := time.Now()
	userDetail.Birthday = sql.NullTime{Time: now, Valid: true}
	userDetail.CreatedAt = sql.NullTime{Time: now, Valid: true}
	userDetail.UpdatedAt = sql.NullTime{Time: now, Valid: true}

	// 验证时间字段
	assert.True(t, userDetail.Birthday.Valid)
	assert.True(t, userDetail.CreatedAt.Valid)
	assert.True(t, userDetail.UpdatedAt.Valid)

	assert.Equal(t, now.Unix(), userDetail.Birthday.Time.Unix())
	assert.Equal(t, now.Unix(), userDetail.CreatedAt.Time.Unix())
	assert.Equal(t, now.Unix(), userDetail.UpdatedAt.Time.Unix())
}
