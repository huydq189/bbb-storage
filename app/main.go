package main

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_fileServerHttpDelivery "github.com/huydq189/bbb-storage/file-server/delivery/http/handler"
	_fileServerHttpDeliveryMiddleware "github.com/huydq189/bbb-storage/file-server/delivery/http/middleware"
	_fileServerRepo "github.com/huydq189/bbb-storage/file-server/repository/local"
	_fileServerUcase "github.com/huydq189/bbb-storage/file-server/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	middL := _fileServerHttpDeliveryMiddleware.InitMiddleware()
	recordDir := viper.GetString(`bbb.presentationDir`)
	maxUploadSize := int64(1 << viper.GetInt(`bbb.maxUploadSize`))
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	e.Use(middL.CORS)

	fileServerRepo := _fileServerRepo.NewLocalFileServerRepository(recordDir, maxUploadSize)
	fu := _fileServerUcase.NewFileServerUsecase(fileServerRepo, timeoutContext)
	_fileServerHttpDelivery.NewFileServerHandler(e, fu)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
