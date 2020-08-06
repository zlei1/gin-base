package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"

	"gin-base/pkg/log"
)

type LogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (l LogWriter) Write(b []byte) (int, error) {
	l.body.Write(b)
	return l.ResponseWriter.Write(b)
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqStartAt := time.Now().UTC()
		reqPath := c.Request.URL.Path

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		reqMethod := c.Request.Method
		reqIp := c.ClientIP()

		lWriter := &LogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = lWriter

		c.Next()

		reqEndAt := time.Now().UTC()
		execTime := reqEndAt.Sub(reqStartAt)

		log.Infof(
			"exec time: %s | req ip: %s | method: %s | path: %s | request: %s | response: %s",
			execTime,
			reqIp,
			reqMethod,
			reqPath,
			c.Request.Body,
			lWriter.body.Bytes(),
		)
	}
}
