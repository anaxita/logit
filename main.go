//Package logit - small package for simple logging in a file and terminal.
package logit

import (
	"fmt"
	"log"
	"os"
	"time"
)

var file *os.File

// New create a *os.File for logging and return non-nil error, if something went wrong.
func New(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		log.Println("Error open logfile:", err)
		return err
	}

	file = f

	return nil
}

// Log write v into logfile and terminal
func Log(v ...interface{}) {
	log.Println(v...)

	_, err := file.WriteString(fmt.Sprintln(time.Now().Format("02.01.2006 15:04:05"), v))

	if err != nil {
		log.Println("Error write in logifle:", err)
	}
}

// Close close logfile and write about this into terminal
func Close() {
	file.Close()
	log.Println("Closed logfile")
}
