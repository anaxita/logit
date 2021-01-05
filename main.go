package logit

import (
	"fmt"
	"log"
	"os"
	"time"
)

var file *os.File

// New ...
func New(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		log.Println("Error open logfile:", err)
	}

	file = f
}

// Log ...
func Log(v ...interface{}) {
	log.Println(v...)

	_, err := file.WriteString(fmt.Sprintln(time.Now().Format("02.01.2006 15:04:05"), v))

	if err != nil {
		log.Println("Error write in logifle:", err)
	}
}

// Close ...
func Close() {
	file.Close()
	log.Println("Closed logfile")
}
