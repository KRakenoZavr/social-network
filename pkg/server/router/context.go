package router

import (
	"net/http"
	"regexp"
)

type Context struct {
	vars map[*regexp.Regexp][]string
}

var context = Context{make(map[*regexp.Regexp][]string)}

func (c *Context) append(reg *regexp.Regexp, vars ...string) {
	c.vars[reg] = vars
}

func Vars(r *http.Request) map[string]string {
	url := r.URL.String()

	for reg, keys := range context.vars {
		if reg.MatchString(url) {
			return getVars(reg, url, keys)
		}
	}

	return nil
}

func getVars(reg *regexp.Regexp, url string, keys []string) map[string]string {
	vars := make(map[string]string)

	for i := 1; i < len(keys); i++ {
		key := keys[i]
		for _, value := range reg.FindAllStringSubmatch(url, -1) {
			vars[key] = value[1]
		}
	}

	return vars
}
