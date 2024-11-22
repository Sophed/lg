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

func Dbug(msg interface{}) {
	fmt.Println(timestamp(), dbug+"DBUG"+rset, msg)
}
func Info(msg interface{}) {
	fmt.Println(timestamp(), info+"INFO"+rset, msg)
}
func Warn(msg interface{}) {
	fmt.Println(timestamp(), warn+"WARN"+rset, msg)
}
func Fatl(msg interface{}) {
	fmt.Println(timestamp(), fatl+"FATL"+rset, msg)
	os.Exit(1)
}
