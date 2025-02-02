package bootstrap_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/sso/frontend/web/bootstrap"
	"gotest.tools/assert"
)

func TestSetupRouterAlive(t *testing.T) {
	server := gin.Default()

	bootstrap.SetupServer(server)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test/alive", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
