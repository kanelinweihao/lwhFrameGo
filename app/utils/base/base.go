package base

import (
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

// interface
type EntityBase interface{}
type EntityDBData interface {
	EntityBase
}

// map
type AttrT1 map[string]interface{}
type AttrT2 map[string]AttrT1
type AttrT3 map[string]AttrT2
type AttrS1 map[string]string
type AttrS2 map[string]AttrS1
type AttrS3 map[string]AttrS2
type BoxData map[string]interface{}

// array_key_exists()
func IsKeyOfAttrStr(key string, attrStr AttrS1) (ok bool) {
	_, ok = attrStr[key]
	return ok
}

// in_array()
func IsInAttrStr(value string, attrStr AttrS1) (ok bool) {
	attrStrFlip := ReserveAttrStr(attrStr)
	_, ok = attrStrFlip[value]
	return ok
}

// array_flip
func ReserveAttrStr(attrStrOld AttrS1) (attrStrNew AttrS1) {
	attrStrNew = make(AttrS1, len(attrStrOld))
	for k, v := range attrStrOld {
		attrStrNew[v] = k
	}
	return attrStrNew
}

// array_merge()
func MergeAttr(args ...AttrT1) (attrMix AttrT1) {
	attrMix = make(AttrT1)
	for _, attrOne := range args {
		for k, v := range attrOne {
			attrMix[k] = v
		}
	}
	return attrMix
}
