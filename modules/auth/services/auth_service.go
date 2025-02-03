package services

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	db          *sql.DB
	redis       *redis.Client
	privateKey  *rsa.PrivateKey
	tokenExpiry time.Duration
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	ClientID  string `json:"client_id"`
	SessionID string `json:"session_id"`
}

func NewAuthService(db *sql.DB, redis *redis.Client, privateKey *rsa.PrivateKey) *AuthService {
	return &AuthService{
		db:          db,
		redis:       redis,
		privateKey:  privateKey,
		tokenExpiry: 24 * time.Hour,
	}
}

func (s *AuthService) ValidateCredentials(ctx context.Context, username, password string) (string, error) {
	var (
		userID         string
		hashedPassword string
	)
	err := s.db.QueryRowContext(ctx,
		"SELECT id, password FROM accounts WHERE username = $1",
		username).Scan(&userID, &hashedPassword)
	if err != nil {
		return "", err
	}

	// TODO: 实现密码验证逻辑
	return userID, nil
}

func (s *AuthService) CreateSession(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()
	sessionData := map[string]interface{}{
		"user_id":    userID,
		"created_at": time.Now(),
	}

	sessionJSON, err := json.Marshal(sessionData)
	if err != nil {
		return "", err
	}

	err = s.redis.Set(ctx, "session:"+sessionID, sessionJSON, 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func (s *AuthService) ValidateSession(ctx context.Context, sessionID string) (string, error) {
	sessionJSON, err := s.redis.Get(ctx, "session:"+sessionID).Result()
	if err != nil {
		return "", err
	}

	var sessionData map[string]interface{}
	if err := json.Unmarshal([]byte(sessionJSON), &sessionData); err != nil {
		return "", err
	}

	userID, ok := sessionData["user_id"].(string)
	if !ok {
		return "", errors.New("invalid session data")
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
		return "", err
	}

	err = s.redis.Set(ctx, "code:"+code, codeJSON, 10*time.Minute).Err()
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *AuthService) ExchangeCodeForTokens(ctx context.Context, code, clientID string) (string, string, error) {
	codeJSON, err := s.redis.Get(ctx, "code:"+code).Result()
	if err != nil {
		return "", "", err
	}

	var codeData map[string]string
	if err := json.Unmarshal([]byte(codeJSON), &codeData); err != nil {
		return "", "", err
	}

	if codeData["client_id"] != clientID {
		return "", "", errors.New("invalid client_id")
	}

	userID, err := s.ValidateSession(ctx, codeData["session_id"])
	if err != nil {
		return "", "", err
	}

	// 生成访问令牌
	accessToken, err := s.generateJWT(userID, clientID, codeData["session_id"])
	if err != nil {
		return "", "", err
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
