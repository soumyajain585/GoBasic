package Routes

import (
	"io"
	"net/http"
	"os"
	"path"
	"root/Controllers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/State", Controllers.ConnectionGet)
	router.GET("/StateId", Controllers.ConnectionGetById)
	router.POST("/State", Controllers.ConnectionPost)
	router.PUT("/State", Controllers.ConnectionUpdate)
	router.DELETE("/State", Controllers.ConnectionDelete)
	return router

}

type MiddlewareResponseBody struct {
	StatusCode int
	Message    string
	Body       struct{}
}

func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		errEnv := godotenv.Load(".env")
		if errEnv != nil {
			log.Error("failed to load env file")
		}
		ApiKey := c.Request.Header.Get("x-api-key")
		generateDailyLogFile()
		if ApiKey != "" {
			if (os.Getenv("X_API_KEY")) != ApiKey {
				RespondWithError(c, http.StatusUnauthorized, "Invalid Key.")
			}
			c.Next()
		} else {
			RespondWithError(c, http.StatusUnauthorized, "Unauthorized Access.")
			return
		}
	}
}
func generateDailyLogFile() {
	now := time.Now() //or time.Now().UTC()
	logFileName := now.Format("2006-01-02") + ".log"

	file, err := os.OpenFile(path.Join("./Logs", logFileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Error(err)
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{
		DisableHTMLEscape: true,
		PrettyPrint:       true,
		TimestampFormat:   "2006-01-02 15:04:05",
	})
	gin.DefaultWriter = io.MultiWriter(file)
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.Debug("Current Time: ", now)
	log.Debug("logFileName: ", logFileName)
}
func RespondWithError(c *gin.Context, code int, message string) {
	response := MiddlewareResponseBody{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
	c.AbortWithStatusJSON(code, response)
}
