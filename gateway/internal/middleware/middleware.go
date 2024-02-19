package middleware

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Reverse(next http.Handler) http.Handler {
	userServiceURL, err := url.Parse("http://user:8081")
	if err != nil {
		log.Println("err 8081", err)
	}
	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)

	portfolioServiceURL, err := url.Parse("http://portfolio:8082")
	if err != nil {
		log.Println("err 8082", err)
	}
	portfolioProxy := httputil.NewSingleHostReverseProxy(portfolioServiceURL)

	cartServiceURL, err := url.Parse("http://cart:8083")
	if err != nil {
		log.Println("err 8083", err)
	}
	cartProxy := httputil.NewSingleHostReverseProxy(cartServiceURL)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/user") {
			userProxy.ServeHTTP(w, r)
		}
		if strings.HasPrefix(r.URL.Path, "/portfolio") {
			portfolioProxy.ServeHTTP(w, r)
		}
		if strings.HasPrefix(r.URL.Path, "/cart") {
			cartProxy.ServeHTTP(w, r)
		}
	})
}

