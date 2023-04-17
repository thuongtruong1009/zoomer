package interceptor

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
)

var (
	e = echo.New()
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	i = &interceptor{}
)

func TestInterceptor_Data(t *testing.T) {
	err := i.Data(c, http.StatusOK, "test")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"data\":\"test\",\"message\":null}\n", rec.Body.String())
}

func TestInterceptor_Error(t *testing.T) {
	code := http.StatusBadRequest
	msg := errors.New("invalid request")
	err := errors.New("invalid parameter")
	actual := i.Error(c, code, msg, err)

	httpError, ok := actual.(*echo.HTTPError)

	// Check HTTP error and status code.
	if !ok {
		t.Errorf("Expected an echo.HTTPError, but got %T", actual)
	}
	if httpError.Code != code {
		t.Errorf("Expected status code %d, but got %d", code, httpError.Code)
	}

	// Check error message
	props, ok := httpError.Message.(*InterceptorProps)
	if !ok {
		t.Errorf("Expected an *InterceptorProps, but got %T", httpError.Message)
	}
	expectedMessage := msg.Error() + " - " + err.Error()
	if props.Message != expectedMessage {
		t.Errorf("Expected message '%s', but got '%s'", expectedMessage, props.Message)
	}
}

