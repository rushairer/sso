package handlers

import (
	"encoding/json"
	"net/http"

	accountsRepositories "github.com/rushairer/sso/modules/accounts/repositories"
	applicationsRepositories "github.com/rushairer/sso/modules/applications/repositories"
	"github.com/rushairer/sso/modules/auth/services"
)

type AuthHandler struct {
	authService     *services.AuthService
	accountRepo     *accountsRepositories.AccountRepository
	applicationRepo *applicationsRepositories.ApplicationRepository
}

func NewAuthHandler(
	authService *services.AuthService,
	accountRepo *accountsRepositories.AccountRepository,
	applicationRepo *applicationsRepositories.ApplicationRepository,
) *AuthHandler {
	return &AuthHandler{
		authService:     authService,
		accountRepo:     accountRepo,
		applicationRepo: applicationRepo,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (h *AuthHandler) HandleAuthorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")

	if responseType != "code" {
		http.Error(w, "unsupported_response_type", http.StatusBadRequest)
		return
	}

	app, err := h.applicationRepo.GetByClientID(r.Context(), clientID)
	if err != nil || app == nil {
		http.Error(w, "invalid_client", http.StatusBadRequest)
		return
	}

	if !app.ValidateRedirectURI(redirectURI) {
		http.Error(w, "invalid_redirect_uri", http.StatusBadRequest)
		return
	}

	// TODO: 实现登录页面渲染
	// 这里应该渲染登录页面，让用户输入凭证
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid_request", http.StatusBadRequest)
		return
	}

	userID, err := h.authService.ValidateCredentials(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, "invalid_grant", http.StatusUnauthorized)
		return
	}

	sessionID, err := h.authService.CreateSession(r.Context(), userID)
	if err != nil {
		http.Error(w, "server_error", http.StatusInternalServerError)
		return
	}

	// 设置会话Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	// 重定向回授权页面
	http.Redirect(w, r, "/authorize"+r.URL.RawQuery, http.StatusFound)
}

func (h *AuthHandler) HandleToken(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid_request", http.StatusBadRequest)
		return
	}

	grantType := r.Form.Get("grant_type")
	if grantType != "authorization_code" {
		http.Error(w, "unsupported_grant_type", http.StatusBadRequest)
		return
	}

	code := r.Form.Get("code")
	clientID := r.Form.Get("client_id")
	clientSecret := r.Form.Get("client_secret")

	if !h.applicationRepo.ValidateClientCredentials(r.Context(), clientID, clientSecret) {
		http.Error(w, "invalid_client", http.StatusUnauthorized)
		return
	}

	accessToken, idToken, err := h.authService.ExchangeCodeForTokens(r.Context(), code, clientID)
	if err != nil {
		http.Error(w, "invalid_grant", http.StatusBadRequest)
		return
	}

	resp := TokenResponse{
		AccessToken: accessToken,
		IDToken:     idToken,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
