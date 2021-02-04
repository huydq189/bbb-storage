package repository

import "github.com/labstack/echo"

// IFileServerRepository represent the file-server's repository contract
type IFileServerRepository interface {
	DownloadFile(ctx echo.Context, id string) error
	UploadFile(ctx echo.Context) error
}
