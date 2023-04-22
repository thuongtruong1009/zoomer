package validators

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestReadRequest(t *testing.T) {
	e := echo.New()

	jsonStr := `{"name":"John","email":"john@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a new response recorder
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	var request struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	err := ReadRequest(ctx, &request)

	assert.NoError(t, err)

	assert.Equal(t, "John", request.Name)
	assert.Equal(t, "john@example.com", request.Email)
}
