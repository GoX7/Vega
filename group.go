package vega

import "strings"

type (
	Group struct {
		name     string
		handlers map[string]func(*Context)
	}
)

func (eng *Engine) Group(name string) *Group {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	if !strings.HasSuffix(name, "/") {
		name = name + "/"
	}

	var group Group
	group.name = name
	group.handlers = make(map[string]func(*Context))
	eng.group = group
	return &group
}

func checkPattern(pattern *string) {
	data := *pattern
	if strings.HasPrefix(data, "/") {
		*pattern = data[1:]
	}
}

func (grp *Group) Get(pattern string, handler func(*Context)) {
	checkPattern(&pattern)
	grp.handlers["GET||"+grp.name+pattern] = handler
}

func (grp *Group) Post(pattern string, handler func(*Context)) {
	checkPattern(&pattern)
	grp.handlers["POST||"+grp.name+pattern] = handler
}

func (grp *Group) Put(pattern string, handler func(*Context)) {
	checkPattern(&pattern)
	grp.handlers["PUT||"+grp.name+pattern] = handler
}

func (grp *Group) Patch(pattern string, handler func(*Context)) {
	checkPattern(&pattern)
	grp.handlers["PATCH||"+grp.name+pattern] = handler
}

func (grp *Group) Delete(pattern string, handler func(*Context)) {
	checkPattern(&pattern)
	grp.handlers["DELETE||"+grp.name+pattern] = handler
}
