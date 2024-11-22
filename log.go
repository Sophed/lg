package lg

import (
	"fmt"
	"os"
	"time"
)

const rset = "\033[0m"
const dbug = "\033[35;1m"
const info = "\033[34;1m"
const warn = "\033[33;1m"
const fatl = "\033[31;1m"

func timestamp() string {
	return time.Now().Format("15:04")
}

// Dbug logs debug messages
func Dbug(msg interface{}) {
	fmt.Println(timestamp(), dbug+"DBUG"+rset, msg)
}

// Info logs informational messages
func Info(msg interface{}) {
	fmt.Println(timestamp(), info+"INFO"+rset, msg)
}

// Warn logs errors/warning messages
func Warn(msg interface{}) {
	fmt.Println(timestamp(), warn+"WARN"+rset, msg)
}

// Fatl logs an error message and calls os.Exit(1), exiting the program
func Fatl(msg interface{}) {
	fmt.Println(timestamp(), fatl+"FATL"+rset, msg)
	os.Exit(1)
}
