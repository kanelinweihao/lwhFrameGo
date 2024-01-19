package conf

import (
	"fmt"
	"html/template"
)

func GetDomain() (domain string) {
	host := getEnvValue("HtmlHost")
	port := getEnvValue("HtmlPort")
	domain = fmt.Sprintf(
		"%s:%s",
		host,
		port)
	return domain
}

func divDisplay(isDivDisplay bool) (msgDivDisplay string) {
	msgDivDisplay = "block"
	if !isDivDisplay {
		msgDivDisplay = "none"
	}
	return msgDivDisplay
}

func GetTmplFuncMap() (tmplFuncMap template.FuncMap) {
	tmplFuncMap = template.FuncMap{
		"funcDivDisplay": divDisplay,
	}
	return tmplFuncMap
}
