package cacheSet

func SetCache(arrPathFileExcel []string) {
	entityCacheSet := MakeEntityCacheSet(arrPathFileExcel)
	entityCacheSet.SetCache()
	return
}
