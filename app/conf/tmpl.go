package conf

import (
	"fmt"
	"html"
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

func divDisplay(isDivDisplay string) (msgDivDisplay string) {
	msgDivDisplay = "block"
	if isDivDisplay == "FALSE" {
		msgDivDisplay = "none"
	}
	return msgDivDisplay
}

func inputDisabled(isInputDisabled string) (msgInputDisabled string) {
	msgInputDisabled = "false"
	if isInputDisabled == "TRUE" {
		msgInputDisabled = "disabled"
	}
	return msgInputDisabled
}

func inputOnkeyup(isInputOnkeyup string) (msgInputOnkeyup string) {
	msgInputOnkeyup = ""
	if isInputOnkeyup == "onlyNumber" {
		msgInputOnkeyup = `onkeyup="value=value.replace(/[^\d]/g,'')"`
		msgInputOnkeyup = html.EscapeString(msgInputOnkeyup)
	}
	return msgInputOnkeyup
}

func GetTmplFuncMap() (tmplFuncMap template.FuncMap) {
	tmplFuncMap = template.FuncMap{
		"funcDivDisplay":    divDisplay,
		"funcInputDisabled": inputDisabled,
		"funcInputOnkeyup":  inputOnkeyup,
	}
	return tmplFuncMap
}
