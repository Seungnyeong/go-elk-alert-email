package swagger

import (
	"fmt"
	"net/http"
	"strings"
	_ "test/docs"
	"test/elastic"

	"github.com/labstack/echo/v4"
)

// @Summary Get All Job
// @Description Get ALL Jonbs
// @Accept json
// @Produce json
// @Success 200 {object} elastic.Instance
// @Router /job/instance [get]
func GetAllInstance(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, elastic.GetSingleton().AllInstance(), " ")
}

// @Summary Get All Job
// @Description Get ALL Jonbs
// @Accept json
// @Produce json
// @Param agentId query []string true "Start Cron Job"
// @Success 200 {string} string "job ok"
// @Router /job/start [get]
func StartJob(c echo.Context) error {
	// elastic.CronJob()
	agentId := strings.Split(c.QueryParams().Get("agentId"), ",")
	err := elastic.CronJob(agentId)
	if err != nil {

		return c.JSONPretty(http.StatusBadRequest, err, "\t")	
	}
	return c.JSONPretty(http.StatusOK, fmt.Sprintf("%s start", c.QueryParams().Get("agentId") ), "\t")
}