package usecase

import "github.com/labstack/echo"

// IFileServerUsecase represent the file-sẻver's usecases
type IFileServerUsecase interface {
	DownloadFile(ctx echo.Context, id string) error
	UploadFile(ctx echo.Context) error
}
