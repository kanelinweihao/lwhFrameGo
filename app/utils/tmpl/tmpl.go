package tmpl

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
	"html/template"
	"net/http"
	"net/url"
)

var paramsRoute = make(base.AttrT3)
var entityBackEnd backEndExecer

type backEndExecer interface {
	ExecBackEnd(paramsIn base.AttrT1) (paramsFromBackend base.AttrT1)
}

/*
Html
*/

func SetTmplAndServer(paramsRouteToSet base.AttrT3, entityBackEndToSet backEndExecer) {
	if len(paramsRouteToSet) == 0 {
		return
	}
	entityBackEnd = entityBackEndToSet
	paramsRoute = paramsRouteToSet
	setTmplToHandle()
	time.Sleep(100, "ms")
	setServerToListen()
	return
}

/*
Tmpl
*/

func setTmplToHandle() {
	for route, _ := range paramsRoute {
		http.HandleFunc(
			route,
			setTmpl)
	}
	return
}

func setTmpl(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	route, ok := getRoute(req)
	if !ok {
		return
	}
	paramsIn, ok := execFromTmpl(route, req)
	paramsOutAppend := make(base.AttrT1)
	if ok {
		paramsOutAppend = entityBackEnd.ExecBackEnd(paramsIn)
	}
	ok = execSetToTmpl(route, resp, paramsOutAppend)
	if !ok {
		return
	}
	return
}

/*
FromTmpl
*/

func execFromTmpl(route string, req *http.Request) (paramsIn base.AttrT1, ok bool) {
	paramsFromTmpl := req.URL.Query()
	if len(paramsFromTmpl) == 0 {
		return nil, false
	}
	paramsInDefault, ok := getParamsInDefault(route)
	if !ok {
		return nil, false
	}
	paramsIn = getParamsFromTmpl(paramsFromTmpl, paramsInDefault)
	return paramsIn, true
}

func getParamsInDefault(route string) (paramsInDefault base.AttrT1, ok bool) {
	paramsInDefault, ok = paramsRoute[route]["paramsInDefault"]
	return paramsInDefault, ok
}

func getParamsFromTmpl(paramsFromTmpl url.Values, paramsInDefault base.AttrT1) (paramsIn base.AttrT1) {
	var valueFromTmpl interface{}
	var valueIn interface{}
	paramsIn = make(base.AttrT1)
	for field, valueDefault := range paramsInDefault {
		valueFromTmpl = paramsFromTmpl.Get(field)
		if valueFromTmpl == "" {
			valueIn = valueDefault
		} else {
			valueIn = valueFromTmpl
		}
		paramsIn[field] = valueIn
	}
	return paramsIn
}

/*
ToTmpl
*/

func getRoute(req *http.Request) (route string, ok bool) {
	route = req.URL.Path
	ok = false
	for routeValid, _ := range paramsRoute {
		if route == routeValid {
			ok = true
		}
	}
	if !ok {
		return "", false
	}
	return route, true
}

func execSetToTmpl(route string, resp http.ResponseWriter, paramsOutAppend base.AttrT1) (ok bool) {
	pathTmpl, ok := getPathTmplByRoute(route)
	if !ok {
		return false
	}
	tmpl := getTmpl(pathTmpl)
	paramsOutDefault, ok := getParamsOutDefault(route)
	if !ok {
		return false
	}
	paramsOutAppendAll := getParamsOutAppendAll(paramsOutAppend)
	paramsOut := getParamsToTmpl(paramsOutDefault, paramsOutAppendAll)
	errTmplExec := tmpl.Execute(resp, paramsOut)
	err.ErrCheck(errTmplExec)
	return true
}

func getPathTmplByRoute(route string) (pathTmpl string, ok bool) {
	pathTmpl, ok = paramsRoute[route]["pathTmpl"]["pathTmpl"].(string)
	if !ok {
		return "", false
	}
	return pathTmpl, true
}

func getTmpl(pathTmpl string) (tmpl *template.Template) {
	fsResource, patternsTmpl := pack.GetFSAndPath(pathTmpl)
	tmpl, errTmplParse := template.ParseFS(fsResource, patternsTmpl)
	err.ErrCheck(errTmplParse)
	return tmpl
}

func getParamsOutDefault(route string) (paramsOutDefault base.AttrT1, ok bool) {
	paramsOutDefault, ok = paramsRoute[route]["paramsOutDefault"]
	return paramsOutDefault, ok
}

func getParamsOutAppendAll(paramsOutAppend base.AttrT1) (paramsOutAppendAll base.AttrT1) {
	paramsOutFromENV := conf.GetParamsTmpl()
	paramsOutAppendAll = base.MergeAttr(paramsOutFromENV, paramsOutAppend)
	return paramsOutAppendAll
}

func getParamsToTmpl(paramsOutDefault base.AttrT1, paramsOutAppendAll base.AttrT1) (paramsOut base.AttrT1) {
	paramsOut = make(base.AttrT1)
	for field, valueDefault := range paramsOutDefault {
		valueOut := valueDefault
		valueOutAppend, ok := paramsOutAppendAll[field].(string)
		if ok {
			valueOut = valueOutAppend
		}
		paramsOut[field] = valueOut
	}
	return paramsOut
}

/*
Listen
*/

func setServerToListen() {
	address := conf.GetDomain()
	server := http.Server{
		Addr: address,
	}
	errServerListen := server.ListenAndServe()
	err.ErrCheck(errServerListen)
	return
}
