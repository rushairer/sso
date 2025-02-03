package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID           string
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURIs []string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

func NewApplication(name string, redirectURIs []string) *Application {
	now := time.Now()
	return &Application{
		ID:           uuid.New().String(),
		Name:         name,
		ClientID:     uuid.New().String(),
		ClientSecret: uuid.New().String(),
		RedirectURIs: redirectURIs,
		CreatedAt:    sql.NullTime{Time: now, Valid: true},
		UpdatedAt:    sql.NullTime{Time: now, Valid: true},
	}
}

func (a *Application) ValidateRedirectURI(uri string) bool {
	for _, allowedURI := range a.RedirectURIs {
		if allowedURI == uri {
			return true
		}
	}
	return false
}
