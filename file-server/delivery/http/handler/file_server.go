package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/huydq189/bbb-storage/domain"
	"github.com/huydq189/bbb-storage/domain/usecase"
)

// ResponseMessage represent the reseponse error struct
type ResponseMessage struct {
	Message string `json:"message"`
}

// FileServerHandler  represent the httphandler for article
type FileServerHandler struct {
	FUsecase usecase.IFileServerUsecase
}

// NewFileServerHandler will initialize the fileServer/ resources endpoint
func NewFileServerHandler(e *echo.Echo, us usecase.IFileServerUsecase) {
	handler := &FileServerHandler{
		FUsecase: us,
	}
	e.GET("/upload", handler.UploadFile)
	e.GET("/download/:id", handler.DownloadFile)
}

// UploadFile will fetch the article based on given params
func (f *FileServerHandler) UploadFile(c echo.Context) error {
	err := f.FUsecase.UploadFile(c)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "Success"})
}

// DownloadFile will store the article by given request body
func (f *FileServerHandler) DownloadFile(c echo.Context) (err error) {
	id := c.Param("id")
	err = f.FUsecase.DownloadFile(c, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
