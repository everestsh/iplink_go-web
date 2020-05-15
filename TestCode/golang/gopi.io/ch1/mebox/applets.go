package main

//Applet import
import (
	"mebox/applets/echo"
	"mebox/applets/ls"
	"mebox/applets/mecmd"
)

//
var Applets map[string]Applet = map[string]Applet{
	"echo":  echo.Echo,
	"ls":    ls.Ls,
	"mecmd": mecmd.Mecmd,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
