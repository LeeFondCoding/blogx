package middleware

import (
	"blogx/service/log_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriter) Header() http.Header {
	return w.Head
}

func Log(c *gin.Context) {
	log := log_service.NewActionLogByGin(c)
	log.SetRequest(c)
	c.Set("log", log)

	res := &ResponseWriter{
		ResponseWriter: c.Writer,
		Head:           make(http.Header),
	}
	c.Writer = res
	c.Next()

	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.Save()
}
