package swagger

import (
	"fmt"
	"net/http"
	"strings"
	_ "test/docs"
	"test/elastic"
	"test/keyinfo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @Summary Get All Job
// @Description 현재 실행되고 있는 잡을 알수있음.
// @Accept json
// @Produce json
// @Success 200 {object} elastic.Instance
// @Router /job/instance [get]
func GetAllInstance(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, elastic.GetSingleton().AllInstance(), " ")
}

// @Summary Job 스케줄 실행 
// @Description monitor.id를 입력하세요
// @Accept json
// @Produce json
// @Param monitorId query []string true "Start Cron Job"
// @Success 200 {string} string "job ok"
// @Router /job/start [get]
func StartJob(c echo.Context) error {
	// elastic.CronJob()
	monitorId := strings.Split(c.QueryParams().Get("monitorId"), ",")
	err := elastic.CronJob(monitorId)
	if err != nil {

		return c.JSONPretty(http.StatusBadRequest, err, "\t")	
	}
	return c.JSONPretty(http.StatusOK, fmt.Sprintf("%s start", c.QueryParams().Get("monitorId") ), "\t")
}

// @Summary 관리자 전체 조회
// @Description 관리자 전체 조회
// @Accept json
// @Produce json
// @Success 200 {string} string "job ok"
// @Router /users/list [get]
func GetUserList(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, service.NewUserRepository().FindAdminUser(), "\t")
}

// @Summary 사용자 조회
// @Description username을 입력하세요
// @Accept json
// @Produce json
// @Param username path string true "Get One User"
// @Success 200 {string} string "job ok"
// @Router /users/{username} [get]
func GetUser(c echo.Context) error {
	username := c.Param("username")
	fmt.Println(username)
	return c.JSONPretty(http.StatusOK, service.NewUserRepository().FindUser(username), "\t")
}


// @title           wkms-alert
// @version         1.0
// @description     wkms alert 서버입니다.
// @termsOfService  http://swagger.io/terms/

// @contact.name   CERT팀 김승녕 매니저
// @contact.url    http://stash.wemakeprice.com
// @contact.email  seungnyeong@wemakeprice.com

// @license.name  위메프 CERT팀 제공
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func SwaggerStart() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/job/instance", GetAllInstance)
	e.GET("/api/v1/job/start", StartJob)
	e.GET("/api/v1/users/list", GetUserList)
	e.GET("/api/v1/users/:username", GetUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}