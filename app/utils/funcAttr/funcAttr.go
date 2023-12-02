package funcAttr

import "github.com/kanelinweihao/lwhFrameGo/app/utils/base"

// IsKeyOfAttrStr array_key_exists()
func IsKeyOfAttrStr(key string, attrStr base.AttrS1) (ok bool) {
	_, ok = attrStr[key]
	return ok
}

// IsInAttrStr in_array()
func IsInAttrStr(value string, attrStr base.AttrS1) (ok bool) {
	attrStrFlip := ReserveAttrStr(attrStr)
	_, ok = attrStrFlip[value]
	return ok
}

// ReserveAttrStr array_flip
func ReserveAttrStr(attrStrOld base.AttrS1) (attrStrNew base.AttrS1) {
	attrStrNew = make(base.AttrS1, len(attrStrOld))
	for k, v := range attrStrOld {
		attrStrNew[v] = k
	}
	return attrStrNew
}

// MergeAttr array_merge()
func MergeAttr(args ...base.AttrT1) (attrMix base.AttrT1) {
	attrMix = make(base.AttrT1)
	for _, attrOne := range args {
		for k, v := range attrOne {
			attrMix[k] = v
		}
	}
	return attrMix
}
