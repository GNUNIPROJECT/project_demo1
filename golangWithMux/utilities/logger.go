package utilities

import "log"

func GetErrorLogger() *log.Logger {
	log1 := log.Default()
	log1.SetPrefix("ERROR : ")
	return log1
}

func GetInfoLogger() *log.Logger {
	log1 := log.Default()
	log1.SetPrefix("INFO : ")
	return log1
}
