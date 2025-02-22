package lg

import (
	"fmt"
	"os"
	"time"
)

const rset = "\033[0m"
const dbug = "\033[34;1m"
const info = "\033[32;1m"
const warn = "\033[33;1m"
const erro = "\033[31;1m"
const fatl = "\033[35;1m"

func timestamp() string {
	return time.Now().Format("15:04")
}

// Dbug logs debug messages
func Dbug(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), dbug+"DBUG"+rset, msg)
}

// Info logs informational messages
func Info(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), info+"INFO"+rset, msg)
}

// Warn logs warning messages
func Warn(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), warn+"WARN"+rset, msg)
}

// Erro logs error messages
func Erro(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), erro+"ERRO"+rset, msg)
}

// Fatl logs a message and calls os.Exit(1), exiting the program
func Fatl(msg interface{}) {
	fmt.Fprintln(os.Stderr, timestamp(), fatl+"FATL"+rset, msg)
	os.Exit(1)
}
