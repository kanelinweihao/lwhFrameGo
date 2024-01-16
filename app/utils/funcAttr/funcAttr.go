package funcAttr

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"sort"
)

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

// SortAttr sort()
func SortAttr(attrT1 typeMap.AttrT1) (attrT1Sorted typeMap.AttrT1) {
	keys := make([]string, 0)
	for k, _ := range attrT1 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	attrT1Sorted = make(typeMap.AttrT1, 0)
	for _, key := range keys {
		attrT1Sorted[key] = attrT1[key]
	}
	return attrT1Sorted
}
