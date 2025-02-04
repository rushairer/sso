package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorType 定义错误类型
type ErrorType string

const (
	ErrorTypeValidation    ErrorType = "VALIDATION_ERROR"
	ErrorTypeAuthorization ErrorType = "AUTHORIZATION_ERROR"
	ErrorTypeNotFound      ErrorType = "NOT_FOUND"
	ErrorTypeInternal      ErrorType = "INTERNAL_ERROR"
	ErrorTypeBadRequest    ErrorType = "BAD_REQUEST"
)

// AppError 定义应用错误结构
type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Err     error     `json:"error,omitempty"`
}

// Error 实现error接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// NewValidationError 创建验证错误
func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
		Err:     err,
	}
}

// NewAuthorizationError 创建授权错误
func NewAuthorizationError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeAuthorization,
		Message: message,
		Err:     err,
	}
}

// NewNotFoundError 创建未找到错误
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeNotFound,
		Message: message,
		Err:     err,
	}
}

// NewInternalError 创建内部错误
func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeInternal,
		Message: message,
		Err:     err,
	}
}

// NewBadRequestError 创建错误请求错误
func NewBadRequestError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeBadRequest,
		Message: message,
		Err:     err,
	}
}

// HTTPError 将AppError转换为HTTP响应
func HTTPError(w http.ResponseWriter, err error) {
	var appErr *AppError
	if e, ok := err.(*AppError); ok {
		appErr = e
	} else {
		appErr = NewInternalError("An unexpected error occurred", err)
	}

	statusCode := getHTTPStatusCode(appErr.Type)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"error": map[string]interface{}{
			"type":    appErr.Type,
			"message": appErr.Message,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// getHTTPStatusCode 根据错误类型返回对应的HTTP状态码
func getHTTPStatusCode(errorType ErrorType) int {
	switch errorType {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeAuthorization:
		return http.StatusUnauthorized
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
