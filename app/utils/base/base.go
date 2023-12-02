package base

type AttrT1 map[string]interface{}
type AttrT2 map[string]AttrT1
type AttrT3 map[string]AttrT2
type AttrS1 map[string]string
type AttrS2 map[string]AttrS1
type AttrS3 map[string]AttrS2
type BoxData map[string]interface{}

type TEntityBase interface{}

type TEntityDBData interface {
	TEntityBase
}

type TEntityParams interface {
	Init(paramsIn AttrT1) TEntityParams
	Validate() TEntityParams
	SetPropertiesAppend(paramsAppend AttrT1) TEntityParams
	ToAttr() (paramsOut AttrT1)
}

type TEntityService interface {
	Init(serviceName string, paramsIn AttrT1) TEntityService
	Exec() (paramsOut AttrT1)
}
