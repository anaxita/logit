//Package logit - small package for simple logging in a file and terminal.
package logit

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var file *os.File
var mut sync.Mutex

// Color pallete map
var (
	colorOff   = "\033[0m"
	colorRed   = "\033[0;31m"
	colorGreen = "\033[0;32m"
	colorBlue  = "\033[0;34m"
)

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

// Log write v into logfile and terminal in blue
func Log(v ...interface{}) {
	go func() {
		log.Println(colorBlue, "ERROR", v, colorOff)

		mut.Lock()
		_, err := file.WriteString(fmt.Sprintln(time.Now().Format("02.01.2006 15:04:05"), "ERROR", v))
		mut.Unlock()

		if err != nil {
			log.Println(colorBlue, "ERROR write in logifle:", err)
		}

	}()
}

// Info write v into logfile and terminal in green color
func Info(v ...interface{}) {
	go func() {
		log.Println(colorGreen, "INFO", v, colorOff)

		mut.Lock()
		_, err := file.WriteString(fmt.Sprintln(time.Now().Format("02.01.2006 15:04:05"), "INFO", v))
		mut.Unlock()

		if err != nil {
			log.Println(colorBlue, "ERROR write in logifle:", err)
		}
	}()
}

// Fatal write v into logfile and terminal in red color
func Fatal(v ...interface{}) {
	go func() {
		mut.Lock()
		_, err := file.WriteString(fmt.Sprintln(time.Now().Format("02.01.2006 15:04:05"), "FATAL", v))
		mut.Unlock()

		if err != nil {
			log.Println(colorBlue, "ERROR write in logifle:", err)
		}

		log.Fatal(colorRed, "FATAL", v, colorOff)
	}()
}

// Close close logfile and write about this into terminal
func Close() {
	file.Close()
	log.Println("Closed logfile")
}

func main() {
	Info("Hello")
}
