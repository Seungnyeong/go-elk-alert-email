package main

import (
	"test/swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           wkms-alert
// @version         1.0
// @description     wkms alert 서버입니다.
// @termsOfService  http://swagger.io/terms/

// @contact.name   CERT팀 김승녕 매니저
// @contact.url    http://stash.wemakeprice.com
// @contact.email  seungnyeong@wemakeprice.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/job/instance", swagger.GetAllInstance)
	e.GET("/api/v1/job/start", swagger.StartJob)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}