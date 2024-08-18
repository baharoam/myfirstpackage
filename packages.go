package mypackage

import "log"

func LogInfo(message string) {
	log.Prinln("Info", message)
}

func LofWarning (message string){
	log.Println("Warn", message)
}