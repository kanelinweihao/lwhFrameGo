package res

func initEntityResSuccess(dataRes interface{}) (entityRes *EntityRes) {
	entityRes = new(EntityRes)
	entityRes.Data = dataRes
	return entityRes
}
