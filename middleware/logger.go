package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logpath := "log/log"
	file, err := os.OpenFile(logpath, os.O_CREATE | os.O_RDWR, 0755)

	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = file


	rotalogs, _ := rotatelogs.New(
		logpath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)


	writerMap := lfshook.WriterMap{
		logrus.DebugLevel: rotalogs,
		logrus.InfoLevel:  rotalogs,
		logrus.WarnLevel:  rotalogs,
		logrus.ErrorLevel: rotalogs,
		logrus.FatalLevel: rotalogs,
	}
	hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	})
	logrus.AddHook(hook)

	fmt.Println("====================")

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		timeCosted0 := time.Since(startTime)
		timeCosted1 := fmt.Sprintf("%f ms", math.Floor(float64(timeCosted0.Nanoseconds()/1000000.0)))

		host := c.Request.Host
		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"timeCosted": timeCosted1,
			"host":       host,
			"statusCode": statusCode,
			"dataSize":   dataSize,
			"clientIp":   clientIp,
			"userAgent":  userAgent,
			"method":     method,
			"path":       path,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode > 500 {
			entry.Error()
		} else if statusCode > 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
