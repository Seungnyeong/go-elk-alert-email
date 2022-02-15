package swagger

import (
	"net/http"
	_ "test/docs"
	"test/elastic"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// User struct
type User struct {
	Name string `json:"name"`
	Age	 int	`json:"age"`
}

type UserMap map[string]User

var (
	userMap UserMap
)


// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 {object} User
// @Router /user/{name} [get]``
func getUser(c echo.Context) error {
    userName := c.Param("name")
	if val, ok := userMap[userName]; ok {
		return c.JSONPretty(http.StatusOK, val, " ")
	}
    
	defaultUser := &User{
		Name : "default",
		Age : 0,
	}
	return c.JSONPretty(http.StatusBadRequest, defaultUser, " ")
}



// @Summary Get All Job
// @Description Get ALL Jonbs
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} instance
// @Router /user [get]
func getAllInstance(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, elastic.GetSingleton().AllInstance(), " ")
}

// @Summary Get All Job
// @Description Get ALL Jonbs
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} instance
// @Router /user [get]
func startJob(c echo.Context) error {
	elastic.CronJob()
	return c.JSONPretty(http.StatusOK, "test", " ")
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} User
// @Router /user [post]
func createUser(c echo.Context) error {
    user := new(User)
    if err := c.Bind(&user); err != nil {
		return err
	}
    userMap[user.Name] = *user
	return c.JSONPretty(http.StatusOK, userMap[user.Name], " ")
}

// @title Wookiist Sample Swagger API
// @version 1.0
// @host localhost:30000
// @BasePath /api/v1
func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/user/:name", getUser)
	e.POST("/api/v1/user", createUser)
	e.GET("/api/job/instance", getAllInstance)
	e.GET("/api/job/start", startJob)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":30000"))
}