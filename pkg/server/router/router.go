package router

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	routes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Match(req) {
			route.handler.ServeHTTP(w, req)
			break
		}
	}
}

func (r *Router) HandleFunc(path string, f http.HandlerFunc) *Route {
	reg := toRegexp(path)
	route := NewRoute(reg, f)
	r.routes = append(r.routes, route)

	context.append(route.reg, reg.SubexpNames()...)
	return route
}

// PathPrefix returns new router which handles all pathes with given prefix
func (r *Router) PathPrefix(prefix string) *Route {
	reg, err := regexp.Compile(fmt.Sprintf("%v\\w*", prefix))
	if err != nil {
		log.Fatal(err)
	}

	route := NewRoute(reg, nil)
	r.routes = append(r.routes, route)

	return route
}

func toRegexp(path string) *regexp.Regexp {
	var newPath string
	peaces := strings.Split(path, "/")
	for _, peace := range peaces {
		if peace == "" {
			continue
		}

		if strings.Contains(peace, ":") {
			newPath += fmt.Sprintf("\\/(?P<%v>[0-9a-zA-Z-]+)", peace[1:])
		} else {
			newPath += fmt.Sprintf("\\/%v", peace)
		}
	}

	if newPath == "" {
		newPath = "/"
	}

	newPath = fmt.Sprintf("^%v$", newPath)
	reg, err := regexp.Compile(newPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return reg
}
