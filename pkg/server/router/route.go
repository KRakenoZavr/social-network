package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Rmethods []string
	reg      *regexp.Regexp
	handler  http.Handler
	Path     string
}

func NewRoute(reg *regexp.Regexp, h http.Handler, p string) *Route {
	return &Route{reg: reg, handler: h, Path: p}
}

func (r *Route) Handle(h http.Handler) {
	r.handler = h
}

func (r *Route) Methods(methods ...string) {
	r.Rmethods = methods
}

func (r *Route) MatchMethods(req *http.Request) bool {
	if len(r.Rmethods) == 0 {
		return true
	}

	for _, m := range r.Rmethods {
		if m == req.Method {
			return true
		}
	}

	return false
}

func (r *Route) Match(req *http.Request) bool {
	if r.MatchMethods(req) && r.reg.MatchString(req.URL.Path) {
		return true
	}

	return false
}
