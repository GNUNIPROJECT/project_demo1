package utilities

func HandleError(err error) {
	if err != nil {
		log1 := GetErrorLogger()
		log1.Panic(err)
	}
}
