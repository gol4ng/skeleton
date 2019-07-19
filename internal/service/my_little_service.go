package service

func (container *Container) GetMyLittleService() func() {
	if container.myLittleService == nil {
		container.myLittleService = myLittleService
	}
	return container.myLittleService
}

// service can be a simple function
func myLittleService() {
	print("print from little service")
}
