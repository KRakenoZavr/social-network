package ctrl

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler
type Middlewares []Middleware

func (mws Middlewares) Apply(hdlr http.Handler) http.Handler {
	if len(mws) == 0 {
		return hdlr
	}
	return mws[1:].Apply(mws[0](hdlr))
}

type Controller struct {
	Logger        *log.Logger
	NextRequestID func() string
}

func (c *Controller) Logging(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func(start time.Time) {
			c.Logger.Println("\033[0m", req.Method, req.URL.Path, time.Since(start))
		}(time.Now())
		hdlr.ServeHTTP(w, req)
	})
}

func (c *Controller) Tracing(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := req.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = c.NextRequestID()
		}
		w.Header().Set("X-Request-Id", requestID)
		hdlr.ServeHTTP(w, req)
	})
}
