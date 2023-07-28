package validators

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

	if err != nil {
		t.Errorf("ReadRequest() = %v; want nil", err)
	}

	if request.Name != "John" {
		t.Errorf("ReadRequest() = %v; want John", request.Name)
	}

	if request.Email != "john@example.com" {
		t.Errorf("ReadRequest() = %v; want john@example.com", request.Email)
	}
}
