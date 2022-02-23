package swagger

import (
	"fmt"
	"net/http"
	"test/crons"
	_ "test/docs"
	"test/elastic"
	"test/keyinfo/service"
	"test/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Message string

const indent string = "\t"
const (
	Success    = Message("Sccuess returned response")
	CannotFind = Message("Cannot find this")
	Error      = Message("Server internal Error")
)

type httpResponse struct {
	TotalCount int         `json:"count,omitempty"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result,omitempty"`
}

// @Summary Get All Job
// @Description 현재 실행되고 있는 잡을 알수있음.
// @Accept json
// @Produce json
// @Success 200 {object} elastic.Instance
// @Router /job/instance [get]
// @Tags   스케줄
func GetAllInstance(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	result := elastic.GetAllInstance(elastic.GetSingleton())
	return c.JSONPretty(http.StatusOK, &httpResponse{
		Message:    string(Success),
		Result:     result,
		TotalCount: len(result),
	}, indent)
}

// @Summary Job 스케줄 실행
// @Description ipv4를 입력하세요
// @Accept json
// @Produce json
// @Param ipv4 query string true "Start Cron Job"
// @Success 200 {string} string "job ok"
// @Router /job/start [get]
// @Tags   스케줄
func StartJob(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if c.QueryParams().Get("ipv4") != "" {
		if !utils.CheckIPAddress(c.QueryParams().Get("ipv4")) {
			return c.JSONPretty(http.StatusBadRequest, &httpResponse{
				Message: fmt.Sprintf("%s is not formatted ipv4", c.QueryParams().Get("ipv4")),
			}, indent)
		}

		err := crons.MonitorInstanceJob(c.QueryParams().Get("ipv4"))
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, &httpResponse{
				Message: err.Error(),
			}, indent)
		}
		return c.JSONPretty(http.StatusOK, &httpResponse{
			Message: "Success",
			Result:  fmt.Sprintf("Start the schedule for instances included in the %s server.", c.QueryParams().Get("ipv4")),
		}, indent)
	}
	return c.JSONPretty(http.StatusBadRequest, &httpResponse{
		Message: "ipv4 cannot be null",
	}, indent)
}

// @Summary 관리자 전체 조회
// @Description 관리자 전체 조회
// @Accept json
// @Produce json
// @Success 200 {string} string "job ok"
// @Router /users/list [get]
// @Tags   계정
func GetUserList(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	result, err := service.NewUserRepository().FindAdminUser()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, &httpResponse{
			Message: err.Error(),
		}, indent)
	}

	return c.JSONPretty(http.StatusOK, &httpResponse{
		Message:    string(Success),
		Result:     result,
		TotalCount: len(result),
	}, indent)
}

// @Summary 사용자 조회
// @Description username을 입력하세요
// @Accept json
// @Produce json
// @Param username path string true "Get One User"
// @Success 200 {string} string "job ok"
// @Router /users/{username} [get]
// @Tags   계정
func GetUser(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	username := c.Param("username")
	user, err := service.NewUserRepository().FindUser(username)
	if err != nil {
		return c.JSONPretty(http.StatusOK, &httpResponse{
			Message: err.Error(),
		}, indent)
	}

	return c.JSONPretty(http.StatusOK, &httpResponse{
		Message: string(Success),
		Result:  user,
	}, indent)
}

// @title           wkms-alert
// @version         1.0
// @description     wkms alert 서버입니다.
// @termsOfService  http://swagger.io/terms/

// @contact.name   CERT팀 김승녕 매니저
// @contact.url    https://stash.wemakeprice.com/projects/SECUTECH/repos/wkms-alert/browse
// @contact.email  seungnyeong@wemakeprice.com

// @license.name  위메프 CERT팀 제공
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      10.107.12.65:8081
// @BasePath  /api/v1
func SwaggerStart(port int) {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`}` + "\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.GET("/api/v1/job/instance", GetAllInstance)
	e.GET("/api/v1/job/start", StartJob)
	e.GET("/api/v1/users/list", GetUserList)
	e.GET("/api/v1/users/:username", GetUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
