package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	methods []string
	reg     *regexp.Regexp
	handler http.Handler
}

func NewRoute(reg *regexp.Regexp, h http.Handler) *Route {
	return &Route{reg: reg, handler: h}
}

func (r *Route) Handle(h http.Handler) {
	r.handler = h
}

func (r *Route) Methods(methods ...string) {
	r.methods = methods
}

func (r *Route) MatchMethods(req *http.Request) bool {
	if len(r.methods) == 0 {
		return true
	}

	for _, m := range r.methods {
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
