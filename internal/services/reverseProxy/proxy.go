package reverseProxy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New() gin.HandlerFunc {
	target, _ := url.Parse("http://koel.test/")
	proxy := httputil.NewSingleHostReverseProxy(target)

	dire := proxy.Director

	proxy.Director = func(req *http.Request) {
		dire(req)
		req.Host = target.Host
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Header.Set("Host", target.Host)
	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
