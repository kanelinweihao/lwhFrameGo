package regex

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"regexp"
)

var (
	strRegexEmail    string = "^[A-Za-z0-9]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$"
	strRegexMobileNo string = "^1[0-9]{10}$"
	strRegexIdCard   string = "^(([0-9]{18})|([0-9]{17}x)|([0-9]{17}X))$"
	strRegexIP       string = "^([0-9]{1,3}\\.){3}[0-9]{1,3}$"
	strRegexDate     string = "^[0-9]{4}-[0-1][0-9]-[0-3][0-9]$"
	strRegexTime     string = "^[0-9]{4}-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9]$"
)

func isStrNeed(strRegex string, strToCheck string) (isValid bool) {
	isValid, errRegexp := regexp.MatchString(strRegex, strToCheck)
	err.ErrCheck(errRegexp)
	return isValid
}

func IsEmail(strToCheck string) (isValid bool) {
	strRegex := strRegexEmail
	isValid = isStrNeed(strRegex, strToCheck)
	return isValid
}

func IsMobileNo(strToCheck string) (isValid bool) {
	strRegex := strRegexMobileNo
	isValid = isStrNeed(strRegex, strToCheck)
	return isValid
}

func IsIdCard(strToCheck string) (isValid bool) {
	strRegex := strRegexIdCard
	isValid = isStrNeed(strRegex, strToCheck)
	return isValid
}

func IsIP(strToCheck string) (isValid bool) {
	strRegex := strRegexIP
	isValid = isStrNeed(strRegex, strToCheck)
	return isValid
}

func IsDate(strToCheck string) (isValid bool) {
	strRegex := strRegexDate
	isValid = isStrNeed(strRegex, strToCheck)
	return isValid
}

func IsTime(strToCheck string) (isValid bool) {
	strRegex := strRegexTime
	isValid = isStrNeed(strRegex, strToCheck)
	return isValid
}
