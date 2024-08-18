package mypackage

import "log"

func LogInfo(message string) {
	log.Println("Info", message)
}

func LofWarning (message string){
	log.Println("Warn", message)
}
