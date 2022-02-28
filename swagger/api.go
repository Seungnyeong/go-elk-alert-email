package swagger

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"test/crons"
	_ "test/docs"
	"test/elastic"
	"test/keyinfo/service"
	"test/utils"
	"time"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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

// @Summary      Alerting 이 되고 있는 인스턴스 전체 확인.
// @Description  현재 실행되고 있는 잡을 알수있음.
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpResponse
// @Failure      400  {object}  httpResponse
// @Failure      404  {object}  httpResponse
// @Failure      500  {object}  httpResponse
// @Router       /job/instance [get]
// @Tags         스케줄
func findRegisterdInstance(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	result := elastic.GetAllInstance(elastic.GetSingleton())
	return c.JSONPretty(http.StatusOK, &httpResponse{
		Message:    string(Success),
		Result:     result,
		TotalCount: len(result),
	}, indent)
}

// @Summary      Alert Instance 등록
// @Description  ipv4를 입력하세요
// @Accept       json
// @Produce      json
// @Param        ipv4  query     string  true  "Start Cron Job"
// @Success      200   {object}  httpResponse
// @Failure      400   {object}  httpResponse
// @Failure      404   {object}  httpResponse
// @Failure      500   {object}  httpResponse
// @Router       /job/start [get]
// @Tags         스케줄
func createJob(c echo.Context) error {
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

// @Summary      Alert Instance 삭제
// @Description  key를 입력하세요
// @Accept       json
// @Produce      json
// @Param        key    path     string  true  "Remove Instance"
// @Success      200   {object}  httpResponse
// @Failure      400   {object}  httpResponse
// @Failure      404   {object}  httpResponse
// @Failure      500   {object}  httpResponse
// @Router       /job/delete/{key} [delete]
// @Tags         스케줄
func removeJobInstance(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	key := c.Param("key")
	decodedValue, err := url.QueryUnescape(key)
	utils.CheckError(err)
	i, _ := elastic.GetInstance(decodedValue, elastic.GetSingleton())
	if i == nil {
		return c.JSONPretty(http.StatusNotFound, &httpResponse{
			Message: fmt.Sprintf("Cannot find %s instance", decodedValue),
		}, indent)
	}

	result := elastic.GetSingleton().RemoveInstance(decodedValue)
	return c.JSONPretty(http.StatusOK, &httpResponse{
			Message: "Success",
			Result:  result,
		}, indent)
}

// @Summary      WKMS 관리자 전체 조회
// @Description  WKMS 관리자 전체 조회
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpResponse
// @Success      200  {object}  httpResponse
// @Failure      400  {object}  httpResponse
// @Failure      404  {object}  httpResponse
// @Failure      500  {object}  httpResponse
// @Router       /users/list [get]
// @Tags         계정
func findAdminUserList(c echo.Context) error {
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

// @Summary      사용자 조회
// @Description  username을 입력하세요
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "Get One User"
// @Success      200       {object}  httpResponse
// @Router       /users/{username} [get]
// @Tags         계정
func findOneUser(c echo.Context) error {
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

// @title           wkms-alert module
// @version         1.0
// @description     wkms alert moddule
// @termsOfService  https://confluence.wemakeprice.com/pages/viewpage.action?pageId=206230173

// @contact.name   보안기술실 메일 전송
// @contact.url    
// @contact.email  secutech@wemakeprice.com

// @license.name  위메프 CERT팀 제공
// @license.url   https://stash.wemakeprice.com/projects/SECUTECH/repos/wkms-alert/browse

// @host      10.107.12.65:8081
// @BasePath  /api/v1
func SwaggerStart(port int) {
	path := utils.GetBinPath()
	if _ , err := os.Stat(path + "/logs"); err != nil {
		merr := os.MkdirAll(path + "/logs", os.ModePerm)
		utils.CheckError(merr)
	}
	
	e := echo.New()
	logf, err := rotatelogs.New(
    path+"/logs/access.%Y%m%d.log",
		rotatelogs.WithLinkName(path+"logs/access_log.log"),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	log.SetOutput(logf)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`}` + "\n",
		Output: logf,
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  0,
	}))

	e.GET("/api/v1/job/instance", findRegisterdInstance)
	e.DELETE("/api/v1/job/delete/:key", removeJobInstance)
	e.GET("/api/v1/job/start", createJob)
	e.GET("/api/v1/users/list", findAdminUserList)
	e.GET("/api/v1/users/:username", findOneUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
