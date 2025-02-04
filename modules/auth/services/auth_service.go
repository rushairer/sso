package services

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	accountsRepositories "github.com/rushairer/sso/modules/accounts/repositories"
	applicationsModels "github.com/rushairer/sso/modules/applications/models"
	applicationsRepositories "github.com/rushairer/sso/modules/applications/repositories"
	"github.com/rushairer/sso/utils/errors"
)

type AuthService struct {
	accountRepo     accountsRepositories.AccountRepository
	applicationRepo applicationsRepositories.ApplicationRepository
	redis           *redis.Client
	privateKey      *rsa.PrivateKey
	tokenExpiry     time.Duration
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	ClientID  string `json:"client_id"`
	SessionID string `json:"session_id"`
}

func NewAuthService(
	accountRepo accountsRepositories.AccountRepository,
	applicationRepo applicationsRepositories.ApplicationRepository,
	redis *redis.Client,
	privateKey *rsa.PrivateKey,
) *AuthService {
	return &AuthService{
		accountRepo:     accountRepo,
		applicationRepo: applicationRepo,
		redis:           redis,
		privateKey:      privateKey,
		tokenExpiry:     24 * time.Hour,
	}
}

func (s *AuthService) ValidateCredentials(ctx context.Context, username, password string) (string, error) {
	account, err := s.accountRepo.GetAccountByUsername(ctx, username)
	if err != nil {
		return "", errors.NewNotFoundError("User not found", err)
	}

	// TODO: 实现密码验证逻辑
	return account.ID, nil
}

func (s *AuthService) CreateSession(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()
	sessionData := map[string]interface{}{
		"user_id":    userID,
		"created_at": time.Now(),
	}

	sessionJSON, err := json.Marshal(sessionData)
	if err != nil {
		return "", errors.NewInternalError("Failed to create session", err)
	}

	err = s.redis.Set(ctx, "session:"+sessionID, sessionJSON, 24*time.Hour).Err()
	if err != nil {
		return "", errors.NewInternalError("Failed to store session", err)
	}

	return sessionID, nil
}

func (s *AuthService) ValidateSession(ctx context.Context, sessionID string) (string, error) {
	sessionJSON, err := s.redis.Get(ctx, "session:"+sessionID).Result()
	if err != nil {
		return "", errors.NewNotFoundError("Session not found", err)
	}

	var sessionData map[string]interface{}
	if err := json.Unmarshal([]byte(sessionJSON), &sessionData); err != nil {
		return "", errors.NewInternalError("Failed to parse session data", err)
	}

	userID, ok := sessionData["user_id"].(string)
	if !ok {
		return "", errors.NewValidationError("Invalid session data", nil)
	}

	return userID, nil
}

func (s *AuthService) GenerateAuthorizationCode(ctx context.Context, sessionID, clientID string) (string, error) {
	code := uuid.New().String()
	codeData := map[string]string{
		"session_id": sessionID,
		"client_id":  clientID,
	}

	codeJSON, err := json.Marshal(codeData)
	if err != nil {
		return "", errors.NewInternalError("Failed to generate authorization code", err)
	}

	err = s.redis.Set(ctx, "code:"+code, codeJSON, 10*time.Minute).Err()
	if err != nil {
		return "", errors.NewInternalError("Failed to store authorization code", err)
	}

	return code, nil
}

func (s *AuthService) ValidateClient(ctx context.Context, clientID string) (*applicationsModels.Application, error) {
	app, err := s.applicationRepo.GetByClientID(ctx, clientID)
	if err != nil {
		return nil, errors.NewInternalError("Failed to validate client", err)
	}
	if app == nil {
		return nil, errors.NewAuthorizationError("Invalid client ID", nil)
	}
	return app, nil
}

func (s *AuthService) ValidateRedirectURI(app *applicationsModels.Application, redirectURI string) bool {
	for _, uri := range app.RedirectURIs {
		if uri == redirectURI {
			return true
		}
	}
	return false
}

func (s *AuthService) ExchangeCodeForTokens(ctx context.Context, code, clientID string) (string, string, error) {
	codeJSON, err := s.redis.Get(ctx, "code:"+code).Result()
	if err != nil {
		return "", "", errors.NewNotFoundError("Authorization code not found", err)
	}

	var codeData map[string]string
	if err := json.Unmarshal([]byte(codeJSON), &codeData); err != nil {
		return "", "", errors.NewInternalError("Failed to parse authorization code", err)
	}

	if codeData["client_id"] != clientID {
		return "", "", errors.NewValidationError("Invalid client ID", nil)
	}

	userID, err := s.ValidateSession(ctx, codeData["session_id"])
	if err != nil {
		return "", "", err
	}

	// 生成访问令牌
	accessToken, err := s.generateJWT(userID, clientID, codeData["session_id"])
	if err != nil {
		return "", "", errors.NewInternalError("Failed to generate access token", err)
	}

	// 生成ID令牌（简化版本）
	idToken := accessToken

	// 删除已使用的授权码
	s.redis.Del(ctx, "code:"+code)

	return accessToken, idToken, nil
}

func (s *AuthService) generateJWT(userID, clientID, sessionID string) (string, error) {
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sso-service",
			Subject:   userID,
			ID:        uuid.New().String(),
		},
		UserID:    userID,
		ClientID:  clientID,
		SessionID: sessionID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(s.privateKey)
}
