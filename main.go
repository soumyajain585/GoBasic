package main

import (
	"io"
	"os"
	"path"
	"root/Routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file ", err)
	}
	router := Routes.SetupRouter()
	// Controllers.OBSlaveConnectionPool()
	// Controllers.OBSlaveConnectionPool()
	// Controllers.MFBackBoneConnection()
	// Controllers.RedisDB()
	router.Static("/assets", "./assets")
	router.Run(":" + os.Getenv("APP_PORT"))
}

func init() {
	now := time.Now() //or time.Now().UTC()
	logFileName := now.Format("2006-01-02") + ".log"
	file, err := os.OpenFile(path.Join("./storage/logs", logFileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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
}
