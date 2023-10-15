package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"grapefruit/kit/log"
	"io"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func WithLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		var reqString string
		if c.Request.Method == "GET" {
			reqString = c.Request.URL.Query().Encode()
		} else if c.Request.Method == "POST" {
			var reqByte []byte
			if c.Request.Body != nil {
				reqByte, _ = io.ReadAll(c.Request.Body)
				c.Request.Body.Close()
			}
			reqString = string(reqByte)

			// Restore the io.ReadCloser to its original state
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqByte))
		} else {
			log.Info("do not support other method !")
			return
		}

		c.Next()

		latency := time.Since(t)

		log.Info("[grapefruit]"+
			"method: %v",
			"url:%v",
			"requestId:%v",
			"Path:%v",
			"latency:%v",
			"status:%v",
			"req:%v",
			"resp:%v",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Header.Get("Requestid"),
			c.Request.Header.Get("Path"),
			latency,
			c.Writer.Status(),
			reqString,
			w.body.String(),
		)
	}
}
