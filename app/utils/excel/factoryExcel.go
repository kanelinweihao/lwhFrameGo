package excel

func MakeEntityExcel() (entityExcel *EntityExcel) {
	entityExcel = new(EntityExcel)
	entityExcel.Init()
	return entityExcel
}

func (self *EntityExcel) Init() *EntityExcel {
	return self
}
