package logger

import "log"

func Error(err error) {
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
func Info(msg string) {
	log.Printf("Info: %s\n", msg)
}
