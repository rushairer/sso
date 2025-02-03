package repositories_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rushairer/sso/modules/applications/models"
	"github.com/rushairer/sso/modules/applications/repositories"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock db 失败: %v", err)
	}
	defer db.Close()

	repo := repositories.NewApplicationRepository(db)
	ctx := context.Background()

	tests := []struct {
		name    string
		app     *models.Application
		mockFn  func()
		wantErr bool
	}{
		{
			name: "成功创建应用",
			app: &models.Application{
				ID:           uuid.New().String(),
				Name:         "测试应用",
				ClientID:     uuid.New().String(),
				ClientSecret: uuid.New().String(),
				RedirectURIs: []string{"http://localhost:3000"},
				CreatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
			},
			mockFn: func() {
				mock.ExpectExec("INSERT INTO applications").WithArgs(
					sqlmock.AnyArg(),
					"测试应用",
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "创建应用失败 - 数据库错误",
			app: &models.Application{
				ID:           uuid.New().String(),
				Name:         "测试应用",
				ClientID:     uuid.New().String(),
				ClientSecret: uuid.New().String(),
				RedirectURIs: []string{"http://localhost:3000"},
				CreatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
			},
			mockFn: func() {
				mock.ExpectExec("INSERT INTO applications").WithArgs(
					sqlmock.AnyArg(),
					"测试应用",
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
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

			err := repo.Create(ctx, tt.app)
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

func TestGetByClientID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock db 失败: %v", err)
	}
	defer db.Close()

	repo := repositories.NewApplicationRepository(db)
	ctx := context.Background()
	testID := uuid.New().String()
	testTime := time.Now()

	tests := []struct {
		name     string
		clientID string
		mockFn   func()
		wantErr  bool
	}{
		{
			name:     "成功获取应用",
			clientID: "test-client-id",
			mockFn: func() {
				redirectURIs := []string{"http://localhost:3000"}
				redirectURIsJSON, _ := json.Marshal(redirectURIs)
				rows := sqlmock.NewRows([]string{
					"id", "name", "client_id", "client_secret",
					"redirect_uris", "created_at", "updated_at",
				}).AddRow(
					testID, "测试应用", "test-client-id", "test-secret",
					string(redirectURIsJSON), testTime, testTime,
				)
				mock.ExpectQuery("SELECT BIN_TO_UUID\\(id\\), name, client_id").WithArgs("test-client-id").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:     "应用不存在",
			clientID: "non-exist-client-id",
			mockFn: func() {
				mock.ExpectQuery("SELECT BIN_TO_UUID\\(id\\), name, client_id").WithArgs("non-exist-client-id").WillReturnError(sql.ErrNoRows)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			app, err := repo.GetByClientID(ctx, tt.clientID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, app)
			} else {
				if err == nil && app != nil {
					assert.Equal(t, tt.clientID, app.ClientID)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("存在未满足的期望: %s", err)
			}
		})
	}
}

func TestValidateClientCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock db 失败: %v", err)
	}
	defer db.Close()

	repo := repositories.NewApplicationRepository(db)
	ctx := context.Background()

	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		mockFn       func()
		want         bool
	}{
		{
			name:         "验证成功",
			clientID:     "test-client-id",
			clientSecret: "test-secret",
			mockFn: func() {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
				mock.ExpectQuery("SELECT COUNT\\(\\*\\)").WithArgs(
					"test-client-id", "test-secret",
				).WillReturnRows(rows)
			},
			want: true,
		},
		{
			name:         "验证失败",
			clientID:     "test-client-id",
			clientSecret: "wrong-secret",
			mockFn: func() {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
				mock.ExpectQuery("SELECT COUNT\\(\\*\\)").WithArgs(
					"test-client-id", "wrong-secret",
				).WillReturnRows(rows)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			got := repo.ValidateClientCredentials(ctx, tt.clientID, tt.clientSecret)
			assert.Equal(t, tt.want, got)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("存在未满足的期望: %s", err)
			}
		})
	}
}
