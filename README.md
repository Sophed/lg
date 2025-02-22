# lg
![image](preview.png)

# Usage
Add the package to your project
```
go get https://github.com/sophed/lg
```

The following functions will then be available
```go
package main

import "github.com/sophed/lg"

func main() {
	lg.Dbug("a debug message")
	lg.Info("and an info message too??")
	lg.Warn("a warning!?!?")
	lg.Fatl("nooooooooo")
}
```

# Callbacks
Use these to set a function which will be run on every subsequent log
```go
func main {
	lg.SetErroCallback(sendToDiscordWebhook)
	err := functionThatFails()
	if err != nil {
		lg.Erro("@sophed something went wrong in production!!")
	}
}

func sendToDiscordWebhook() {
  	...
}

func functionThatFails() error {
  	...
}
```

# Why not stdlib?
By default, the Go `log` package doesn't have the nicest formatting and while it is possible to edit, it's an (albeit minor) annoyance every time I make a new Go project. [charmbracelet/log](https://github.com/charmbracelet/log) has a bunch of stuff I wouldn't use.

At the end of the day I mainly just like building and using my own tools :)
