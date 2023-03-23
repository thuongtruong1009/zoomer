package delivery

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"zoomer/internal/auth/repository"
	"zoomer/internal/models"
	"zoomer/internal/resources/presenter"
	"zoomer/internal/resources/usecase"
	"zoomer/utils"
)

type resourceHandler struct {
	resourceUC usecase.ResourceUseCase
}

func NewResourceHandler(resourceUC usecase.ResourceUseCase) *resourceHandler {
	return &resourceHandler{resourceUC: resourceUC}
}

func mapResource(r *models.Resource) *presenter.ResourceResponse {
	return &presenter.ResourceResponse{
		Id:        r.Id,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		CreatedBy: r.CreatedBy,
	}
}

func mapResources(ro []*models.Resource) []*presenter.ResourceResponse {
	out := make([]*presenter.ResourceResponse, len(ro))

	for i, b := range ro {
		out[i] = mapResource(b)
	}
	return out
}

func (rh *resourceHandler) UploadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get(repository.CtxUserKey).(string)
		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer src.Close()
		_, err = rh.resourceUC.UploadResource(c.Request().Context(), src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		resource := &models.Resource{
			Name:      file.Filename,
			CreatedBy: userId,
		}
		err = rh.resourceUC.CreateResource(c.Request().Context(), resource)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, mapResource(resource))
	}
}

func (rh *resourceHandler) GetImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileName := c.Param("fileName")
		file, err := rh.resourceUC.GetResource(c.Request().Context(), fileName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.Stream(http.StatusOK, utils.GetContentType(fileName), file)
	}
}

func (rh *resourceHandler) DeleteImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileName := c.Param("fileName")
		err := rh.resourceUC.DeleteResource(c.Request().Context(), fileName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, fmt.Sprintf("File %s deleted", fileName))
	}
}
