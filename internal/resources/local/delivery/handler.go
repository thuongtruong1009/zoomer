package delivery

import (
	"log"
	"net/http"
	"os"
	"errors"
	"github.com/labstack/echo/v4"
	"zoomer/pkg/interceptor"
	"zoomer/pkg/constants"
	"zoomer/internal/resources/local/usecase"
)

func init() {
	if _, err := os.Stat("public/upload"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/upload", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	if _, err := os.Stat("public/upload"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/upload", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

type localHandler struct {
	inter  interceptor.IInterceptor
	usecase usecase.ILocalResourceUseCase
}

func NewLocalResourceHandler(inter interceptor.IInterceptor, usecase usecase.ILocalResourceUseCase) LocalHandler {
	return &localHandler{
		inter: inter,
		usecase: usecase,
	}
}

func (lh *localHandler) UploadSingleFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return lh.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		filename, err := lh.usecase.UploadSingleFile(c.Request().Context(), file); if err != nil {
			return lh.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return lh.inter.Data(c, http.StatusOK, map[string]interface{}{"file": filename})
	}
}

func (lh *localHandler) UploadMultipleFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return lh.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		files := form.File["images"]
		filepaths, err := lh.usecase.UploadMultipleFile(c.Request().Context(), files); if err != nil {
			return lh.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return lh.inter.Data(c, http.StatusOK, map[string]interface{}{"files": filepaths})
	}
}

func (lh *localHandler) DeleteSingleFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		filename := c.Param("fileName")

		err := lh.usecase.DeleteSingleFile(c.Request().Context(), filename); if err != nil {
			return lh.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return lh.inter.Data(c, http.StatusOK, map[string]interface{}{"message": "success"})
	}
}

func (lh *localHandler) DeleteMultipleFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		filenames := c.QueryParams()["fileNames"]

		err := lh.usecase.DeleteMultipleFile(c.Request().Context(), filenames); if err != nil {
			return lh.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return lh.inter.Data(c, http.StatusOK, map[string]interface{}{"message": "success"})
	}
}
