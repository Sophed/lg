package lg

import (
	"fmt"
	"os"
	"time"
)

const (
	rset = "\033[0m"
	dbug = "\033[34;1m"
	info = "\033[32;1m"
	warn = "\033[33;1m"
	erro = "\033[31;1m"
	fatl = "\033[35;1m"
)

type callback func()

var (
	dbugCallback callback
	infoCallback callback
	warnCallback callback
	erroCallback callback
	fatlCallback callback
)

func timestamp() string {
	return time.Now().Format("15:04")
}

// Dbug logs debug messages
func Dbug(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), dbug+"DBUG"+rset, msg)
	if dbugCallback != nil {
		dbugCallback()
	}
}

// Info logs informational messages
func Info(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), info+"INFO"+rset, msg)
	if infoCallback != nil {
		infoCallback()
	}
}

// Warn logs warning messages
func Warn(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), warn+"WARN"+rset, msg)
	if warnCallback != nil {
		warnCallback()
	}
}

// Erro logs error messages
func Erro(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), erro+"ERRO"+rset, msg)
	if erroCallback != nil {
		erroCallback()
	}
}

// Fatl logs a message and calls os.Exit(1), exiting the program
func Fatl(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), fatl+"FATL"+rset, msg)
	if fatlCallback != nil {
		fatlCallback()
	}
	os.Exit(1)
}

// SetDbugCallback sets a callback function to run on every Dbug() use
func SetDbugCallback(c func()) {
	dbugCallback = c
}

// SetInfoCallback sets a callback function to run on every Info() use
func SetInfoCallback(c func()) {
	infoCallback = c
}

// SetWarnCallback sets a callback function to run on every Warn() use
func SetWarnCallback(c func()) {
	warnCallback = c
}

// SetErroCallback sets a callback function to run on every Erro() use
func SetErroCallback(c func()) {
	erroCallback = c
}

// SetFatlCallback sets a callback function to run on every Fatl() use
func SetFatlCallback(c func()) {
	fatlCallback = c
}
