//Package logit - small package for simple logging in a file and terminal.
package logit

import (
	"fmt"
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
		os.Stdout.WriteString("Lofile is already open\n")
		return nil
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766)

	if err != nil {
		os.Stdout.WriteString(fmt.Sprintln("[FILE_ERROR] cannot open logile:", err))
		return err
	}

	file = f
	return nil
}

// logit is format log message to ERROR INFO or FATAL style
func logit(level string, color string, messages ...interface{}) {
	logline := fmt.Sprintf("%s%s [%s] %v%s\n", time.Now().Format("02.01.2006 15:04:05"), color, level, colorOff, messages)
	fileline := fmt.Sprintf("%s [%s] %v\n", time.Now().Format("02.01.2006 15:04:05"), level, messages)

	mut.Lock()
	defer mut.Unlock()

	_, err := file.WriteString(fileline)
	if err != nil {
		os.Stdout.WriteString(fmt.Sprintln("[FILE_ERROR] cannot write into logile:", err))
	}

	if level == "FATAL" {
		os.Stdout.WriteString(logline)
		os.Exit(1)
	} else {
		os.Stdout.WriteString(logline)
	}
}

// Log write v into logfile and terminal in blue
func Log(v ...interface{}) {
	logit("ERROR", colorBlue, v...)
}

// Info write v into logfile and terminal in green color
func Info(v ...interface{}) {
	logit("INFO", colorGreen, v...)
}

// Fatal write v into logfile and terminal in red color
func Fatal(v ...interface{}) {
	logit("FATAL", colorRed, v...)
}

// Close close logfile and write about this into terminal
func Close() {
	if file != nil {
		file.Close()
		os.Stdout.WriteString("FILE closed\n")
	}
}
