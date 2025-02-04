package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/rushairer/sso/modules/applications/models"
)

// ApplicationRepository 应用仓储接口
type ApplicationRepository interface {
	Create(ctx context.Context, app *models.Application) error
	GetByClientID(ctx context.Context, clientID string) (*models.Application, error)
	ValidateClientCredentials(ctx context.Context, clientID, clientSecret string) bool
}

// applicationRepository 应用仓储实现
type applicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) ApplicationRepository {
	return &applicationRepository{db: db}
}

func (r *applicationRepository) Create(ctx context.Context, app *models.Application) error {
	redirectURIsJSON, err := json.Marshal(app.RedirectURIs)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO applications (id, name, client_id, client_secret, redirect_uris, created_at, updated_at)
		VALUES (UUID_TO_BIN(?), ?, ?, ?, ?, ?, ?)
	`
	_, err = r.db.ExecContext(ctx, query,
		app.ID,
		app.Name,
		app.ClientID,
		app.ClientSecret,
		string(redirectURIsJSON),
		app.CreatedAt.Time,
		app.UpdatedAt.Time,
	)
	return err
}

func (r *applicationRepository) GetByClientID(ctx context.Context, clientID string) (*models.Application, error) {
	query := `
		SELECT BIN_TO_UUID(id), name, client_id, client_secret, redirect_uris, created_at, updated_at
		FROM applications
		WHERE client_id = ?
	`

	app := &models.Application{}
	var redirectURIsJSON string

	err := r.db.QueryRowContext(ctx, query, clientID).Scan(
		&app.ID,
		&app.Name,
		&app.ClientID,
		&app.ClientSecret,
		&redirectURIsJSON,
		&app.CreatedAt,
		&app.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if err := json.Unmarshal([]byte(redirectURIsJSON), &app.RedirectURIs); err != nil {
		return nil, err
	}
	return app, nil
}

func (r *applicationRepository) ValidateClientCredentials(ctx context.Context, clientID, clientSecret string) bool {
	query := `
		SELECT COUNT(*)
		FROM applications
		WHERE client_id = ? AND client_secret = ?
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, clientID, clientSecret).Scan(&count)
	return err == nil && count > 0
}
