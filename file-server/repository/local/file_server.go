package local

import (
	"io"
	"os"
	"strings"

	"github.com/huydq189/bbb-storage/domain"
	"github.com/huydq189/bbb-storage/domain/repository"
	"github.com/labstack/echo"
)

type localFileServerRepository struct {
	recordDir     string
	maxUploadSize int64
}

// NewLocalFileServerRepository will create an object that represent the fileServer.Repository interface
func NewLocalFileServerRepository(recordDir string, maxUploadSize int64) repository.IFileServerRepository {
	return &localFileServerRepository{recordDir, maxUploadSize}
}

func (m *localFileServerRepository) DownloadFile(ctx echo.Context, id string) error {
	filePath := m.recordDir + id + "/deskshare/deskshare.webm"
	arraySlash := strings.SplitAfter(filePath, "/")
	fileName := arraySlash[len(arraySlash)-1]
	f, err := os.Open(filePath)
	if err != nil {
		return domain.ErrInternalServerError
	}
	defer f.Close()

	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	return ctx.Attachment(filePath, fileName)
}

func (m *localFileServerRepository) UploadFile(ctx echo.Context) error {
	//------------
	// Read files
	//------------

	// Multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["data"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.OpenFile(m.recordDir+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return err
}
