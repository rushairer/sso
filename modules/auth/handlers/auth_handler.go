package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/rushairer/sso/modules/auth/services"
	"github.com/rushairer/sso/utils/errors"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
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

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		errors.HTTPError(w, errors.NewBadRequestError("Invalid request body", err))
		return
	}

	// 验证用户凭据
	userID, err := h.authService.ValidateCredentials(r.Context(), loginReq.Username, loginReq.Password)
	if err != nil {
		errors.HTTPError(w, errors.NewAuthorizationError("Invalid credentials", err))
		return
	}

	// 创建会话
	sessionID, err := h.authService.CreateSession(r.Context(), userID)
	if err != nil {
		errors.HTTPError(w, errors.NewInternalError("Failed to create session", err))
		return
	}

	// 生成授权码
	clientID := r.URL.Query().Get("client_id")
	code, err := h.authService.GenerateAuthorizationCode(r.Context(), sessionID, clientID)
	if err != nil {
		errors.HTTPError(w, errors.NewInternalError("Failed to generate authorization code", err))
		return
	}

	// 返回授权码
	redirectURI := r.URL.Query().Get("redirect_uri")
	redirectURL, err := url.Parse(redirectURI)
	if err != nil {
		errors.HTTPError(w, errors.NewBadRequestError("Invalid redirect URI", err))
		return
	}

	q := redirectURL.Query()
	q.Set("code", code)
	redirectURL.RawQuery = q.Encode()

	http.Redirect(w, r, redirectURL.String(), http.StatusFound)
}

func (h *AuthHandler) HandleToken(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errors.HTTPError(w, errors.NewBadRequestError("Invalid request", err))
		return
	}

	grantType := r.Form.Get("grant_type")
	if grantType != "authorization_code" {
		errors.HTTPError(w, errors.NewBadRequestError("Unsupported grant type", nil))
		return
	}

	code := r.Form.Get("code")
	clientID := r.Form.Get("client_id")
	if code == "" || clientID == "" {
		errors.HTTPError(w, errors.NewBadRequestError("Missing required parameters", nil))
		return
	}

	// 验证客户端应用
	_, err := h.authService.ValidateClient(r.Context(), clientID)
	if err != nil {
		errors.HTTPError(w, err)
		return
	}

	// 交换授权码获取令牌
	accessToken, idToken, err := h.authService.ExchangeCodeForTokens(r.Context(), code, clientID)
	if err != nil {
		errors.HTTPError(w, errors.NewInternalError("Failed to exchange code for tokens", err))
		return
	}

	// 返回令牌
	response := TokenResponse{
		AccessToken: accessToken,
		IDToken:     idToken,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) HandleAuthorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")

	if responseType != "code" {
		errors.HTTPError(w, errors.NewBadRequestError("Unsupported response type", nil))
		return
	}

	if clientID == "" || redirectURI == "" {
		errors.HTTPError(w, errors.NewBadRequestError("Missing required parameters", nil))
		return
	}

	// 验证客户端应用
	app, err := h.authService.ValidateClient(r.Context(), clientID)
	if err != nil {
		errors.HTTPError(w, err)
		return
	}

	// 验证重定向URI
	if !h.authService.ValidateRedirectURI(app, redirectURI) {
		errors.HTTPError(w, errors.NewBadRequestError("Invalid redirect URI", nil))
		return
	}

	// TODO: 实现用户认证和授权页面的渲染
	// 这里应该渲染一个登录页面或授权页面，让用户进行身份验证并授权应用访问
	// 用户授权后，生成授权码并重定向到应用的回调地址
}
