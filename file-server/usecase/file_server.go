package usecase

import (
	"context"
	"time"

	"github.com/huydq189/bbb-storage/domain/repository"
	"github.com/huydq189/bbb-storage/domain/usecase"
	"github.com/labstack/echo"
)

type fileServerUsecase struct {
	fileServerRepo repository.IFileServerRepository
	contextTimeout time.Duration
}

// NewFileServerUsecase will create new an fileServerUsecase object representation of domain.IFileServerUsecase interface
func NewFileServerUsecase(f repository.IFileServerRepository, timeout time.Duration) usecase.IFileServerUsecase {
	return &fileServerUsecase{
		fileServerRepo: f,
		contextTimeout: timeout,
	}
}

func (f *fileServerUsecase) DownloadFile(c echo.Context, id string) error {
	_, cancel := context.WithTimeout(c.Request().Context(), f.contextTimeout)
	defer cancel()
	err := f.fileServerRepo.DownloadFile(c, id)
	return err
}

func (f *fileServerUsecase) UploadFile(c echo.Context) error {
	_, cancel := context.WithTimeout(c.Request().Context(), f.contextTimeout)
	defer cancel()
	err := f.fileServerRepo.UploadFile(c)
	return err
}
