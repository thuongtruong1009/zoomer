package utils

import (
	"testing"
	"strings"
	"net/http"
    "context"
	"net/http/httptest"
	"github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
    // "github.com/go-playground/validator/v10"
)

func TestRandomString(t *testing.T) {
	got := RandomString(10)
	if len(got) != 10 {
		t.Errorf("RandomString(10) = %s; want 10 characters", got)
	}
	for _, c := range got {
		if c < 'A' || c > 'z' {
			t.Errorf("RandomString(10) = %s; want only letters", got)
		}
	}
}

func TestReadRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John"}`))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	request := &struct {
		Name string `json:"name" validate:"required"`
	}{}

	if err := ReadRequest(c, request); err != nil {
		t.Errorf("ReadRequest() = %v; want nil", err)
	}

	if request.Name != "John" {
		t.Errorf("ReadRequest() = %v; want John", request.Name)
	}
}

func TestValidateStruct(t *testing.T) {
    type TestStruct struct {
        Name string `validate:"required"`
        Age  int    `validate:"required,gt=0"`
    }
    // Create a new instance of the validator
    // validate := validator.New()
    // Create a context
    ctx := context.Background()

    // Test case 1: valid input
    input1 := TestStruct{Name: "Alice", Age: 25}
    err := ValidateStruct(ctx, input1)
    assert.NoError(t, err)

    // Test case 2: missing required field
    input2 := TestStruct{Age: 30}
    err = ValidateStruct(ctx, input2)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "Name")

    // Test case 3: invalid age
    input3 := TestStruct{Name: "Bob", Age: -5}
    err = ValidateStruct(ctx, input3)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "Age")

    // Test case 4: nil input
    var input4 *TestStruct
    err = ValidateStruct(ctx, input4)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "nil")

    // Test case 5: unsupported input type
    input5 := "not a struct"
    err = ValidateStruct(ctx, input5)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "type")
}
