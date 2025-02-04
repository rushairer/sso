package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/rushairer/sso/modules/accounts/models"
)

// AccountRepository 账户仓储接口
type AccountRepository interface {
	CreateAccount(ctx context.Context, account *models.Account) error
	GetAccountByID(ctx context.Context, id string) (*models.Account, error)
	GetAccountWithDetails(ctx context.Context, id string) (*models.Account, *models.UserDetail, error)
	GetAccountByUsername(ctx context.Context, username string) (*models.Account, error)
}

// accountRepository 账户仓储实现
type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(ctx context.Context, account *models.Account) error {
	query := `
		INSERT INTO accounts (
			id, email, password, status, created_at, updated_at
		) VALUES (
			UUID_TO_BIN(?), ?, ?, ?, ?, ?
		)
	`
	_, err := r.db.ExecContext(ctx, query,
		account.ID,
		account.Email,
		account.Password,
		account.Status,
		time.Now(),
		time.Now(),
	)
	return err
}

func (r *accountRepository) GetAccountByID(ctx context.Context, id string) (*models.Account, error) {
	account := &models.Account{}
	query := `
		SELECT
			UUID_FROM_BIN(id) as id,
			email,
			password,
			status,
			created_at,
			updated_at,
			deleted_at,
			email_verified_at,
			last_login_at
		FROM accounts
		WHERE id = UUID_TO_BIN(?) AND deleted_at IS NULL
	`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.Email,
		&account.Password,
		&account.Status,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.EmailVerifiedAt,
		&account.LastLoginAt,
	)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *accountRepository) GetAccountByUsername(ctx context.Context, username string) (*models.Account, error) {
	account := &models.Account{}
	query := `
		SELECT
			UUID_FROM_BIN(id) as id,
			email,
			password,
			status,
			created_at,
			updated_at,
			deleted_at,
			email_verified_at,
			last_login_at
		FROM accounts
		WHERE email = ? AND deleted_at IS NULL
	`
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&account.ID,
		&account.Email,
		&account.Password,
		&account.Status,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.EmailVerifiedAt,
		&account.LastLoginAt,
	)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *accountRepository) GetAccountWithDetails(ctx context.Context, id string) (*models.Account, *models.UserDetail, error) {
	query := `
		SELECT
			UUID_FROM_BIN(a.id) as account_id,
			a.email,
			a.password,
			a.status,
			a.created_at,
			a.updated_at,
			a.deleted_at,
			a.email_verified_at,
			a.last_login_at,
			UUID_FROM_BIN(ud.account_id) as detail_account_id,
			ud.nick_name,
			ud.avatar,
			ud.gender,
			ud.birthday,
			ud.bio,
			ud.created_at,
			ud.updated_at
		FROM accounts a
		LEFT JOIN user_details ud ON a.id = ud.account_id
		WHERE a.id = UUID_TO_BIN(?) AND a.deleted_at IS NULL
	`

	account := &models.Account{}
	detail := &models.UserDetail{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.Email,
		&account.Password,
		&account.Status,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.EmailVerifiedAt,
		&account.LastLoginAt,
		&detail.AccountID,
		&detail.NickName,
		&detail.Avatar,
		&detail.Gender,
		&detail.Birthday,
		&detail.Bio,
		&detail.CreatedAt,
		&detail.UpdatedAt,
	)
	if err != nil {
		return nil, nil, err
	}

	return account, detail, nil
}
