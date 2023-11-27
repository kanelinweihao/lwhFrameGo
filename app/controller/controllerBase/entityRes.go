package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"net/http"
)

type EntityRes struct {
	Code int
	Msg  string
	Data interface{}
}

func getResSuccess(dataRes interface{}) (jsonRes string) {
	entityRes := makeEntityResSuccess(dataRes)
	jsonRes = conv.ToJsonFromEntity(entityRes)
	return jsonRes
}

func makeEntityResSuccess(dataRes interface{}) (entityRes *EntityRes) {
	entityRes = new(EntityRes)
	entityRes.Code = http.StatusOK
	entityRes.Msg = "OK"
	entityRes.Data = dataRes
	return entityRes
}
