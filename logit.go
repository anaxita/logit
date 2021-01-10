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
	if file != nil {
		return nil
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		log.Println("Error open logfile:", err)
		return err
	}

	file = f

	return nil
}

// logit is format log message to ERROR INFO or FATAL style
func logit(level string, color string, message ...interface{}) {
	logline := fmt.Sprintf("%s%s %s %v%s", color, time.Now().Format("02.01.2006 15:04:05"), level, message, colorOff)
	fileline := fmt.Sprintf("%s %s %v\n", time.Now().Format("02.01.2006 15:04:05"), level, message)

	_, err := file.WriteString(fileline)
	if err != nil {
		log.Println("[FILE_ERROR] cannot write into logile:", err)
	}

	if level == "FATAL" {
		log.Fatal(logline)
	} else {
		log.Println(logline)
	}
}

// Log write v into logfile and terminal in blue
func Log(v ...interface{}) {
	logit("ERROR", colorBlue, v)
}

// Info write v into logfile and terminal in green color
func Info(v ...interface{}) {
	logit("INFO", colorGreen, v)
}

// Fatal write v into logfile and terminal in red color
func Fatal(v ...interface{}) {
	logit("FATAL", colorRed, v)
}

// Close close logfile and write about this into terminal
func Close() {
	if file != nil {
		file.Close()
		log.Println("Closed logfile")
	}
}
