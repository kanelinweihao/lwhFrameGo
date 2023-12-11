package funcAttr

import "github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"

// IsKeyOfAttrStr array_key_exists()
func IsKeyOfAttrStr(key string, attrStr typeMap.AttrS1) (ok bool) {
	_, ok = attrStr[key]
	return ok
}

// IsInAttrStr in_array()
func IsInAttrStr(value string, attrStr typeMap.AttrS1) (ok bool) {
	attrStrFlip := ReserveAttrStr(attrStr)
	_, ok = attrStrFlip[value]
	return ok
}

// ReserveAttrStr array_flip
func ReserveAttrStr(attrStrOld typeMap.AttrS1) (attrStrNew typeMap.AttrS1) {
	attrStrNew = make(typeMap.AttrS1, len(attrStrOld))
	for k, v := range attrStrOld {
		attrStrNew[v] = k
	}
	return attrStrNew
}

// MergeAttr array_merge()
func MergeAttr(args ...typeMap.AttrT1) (attrMix typeMap.AttrT1) {
	attrMix = make(typeMap.AttrT1)
	for _, attrOne := range args {
		for k, v := range attrOne {
			attrMix[k] = v
		}
	}
	return attrMix
}
