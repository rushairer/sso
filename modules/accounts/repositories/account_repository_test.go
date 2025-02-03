package repositories_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rushairer/sso/modules/accounts/models"
	"github.com/rushairer/sso/modules/accounts/repositories"
)

func TestCreateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock db 失败: %v", err)
	}
	defer db.Close()

	repo := repositories.NewAccountRepository(db)
	ctx := context.Background()

	tests := []struct {
		name    string
		account *models.Account
		mockFn  func()
		wantErr bool
	}{
		{
			name: "成功创建账户",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "test@example.com",
				Password: "password123",
				Status:   1,
			},
			mockFn: func() {
				mock.ExpectExec("INSERT INTO accounts").WithArgs(
					sqlmock.AnyArg(),
					"test@example.com",
					"password123",
					int8(1),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "数据库错误",
			account: &models.Account{
				ID:       uuid.New().String(),
				Email:    "test@example.com",
				Password: "password123",
				Status:   1,
			},
			mockFn: func() {
				mock.ExpectExec("INSERT INTO accounts").WithArgs(
					sqlmock.AnyArg(),
					"test@example.com",
					"password123",
					int8(1),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnError(errors.New("数据库错误"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			err := repo.CreateAccount(ctx, tt.account)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("存在未满足的期望: %s", err)
			}
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock db 失败: %v", err)
	}
	defer db.Close()

	repo := repositories.NewAccountRepository(db)
	ctx := context.Background()
	testID := uuid.New().String()
	testTime := time.Now()

	tests := []struct {
		name    string
		id      string
		mockFn  func()
		wantErr bool
	}{
		{
			name: "成功获取账户",
			id:   testID,
			mockFn: func() {
				rows := sqlmock.NewRows([]string{
					"id", "email", "password", "status",
					"created_at", "updated_at", "deleted_at",
					"email_verified_at", "last_login_at",
				}).AddRow(
					testID, "test@example.com", "password123", 1,
					testTime, testTime, nil,
					testTime, testTime,
				)
				mock.ExpectQuery("SELECT (.+) FROM accounts").WithArgs(testID).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "账户不存在",
			id:   testID,
			mockFn: func() {
				mock.ExpectQuery("SELECT (.+) FROM accounts").WithArgs(testID).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			account, err := repo.GetAccountByID(ctx, tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, account)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, account)
				assert.Equal(t, tt.id, account.ID)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("存在未满足的期望: %s", err)
			}
		})
	}
}

func TestGetAccountWithDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock db 失败: %v", err)
	}
	defer db.Close()

	repo := repositories.NewAccountRepository(db)
	ctx := context.Background()
	testID := uuid.New().String()
	testTime := time.Now()

	tests := []struct {
		name    string
		id      string
		mockFn  func()
		wantErr bool
	}{
		{
			name: "成功获取账户和详情",
			id:   testID,
			mockFn: func() {
				rows := sqlmock.NewRows([]string{
					"account_id", "email", "password", "status",
					"created_at", "updated_at", "deleted_at",
					"email_verified_at", "last_login_at",
					"detail_account_id", "nick_name", "avatar",
					"gender", "birthday", "bio",
					"created_at", "updated_at",
				}).AddRow(
					testID, "test@example.com", "password123", 1,
					testTime, testTime, nil,
					testTime, testTime,
					testID, "测试用户", "avatar.jpg",
					1, testTime, "测试简介",
					testTime, testTime,
				)
				mock.ExpectQuery("SELECT (.+) FROM accounts").WithArgs(testID).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "账户不存在",
			id:   testID,
			mockFn: func() {
				mock.ExpectQuery("SELECT (.+) FROM accounts").WithArgs(testID).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			account, detail, err := repo.GetAccountWithDetails(ctx, tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, account)
				assert.Nil(t, detail)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, account)
				assert.NotNil(t, detail)
				assert.Equal(t, tt.id, account.ID)
				assert.Equal(t, tt.id, detail.AccountID)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("存在未满足的期望: %s", err)
			}
		})
	}
}
