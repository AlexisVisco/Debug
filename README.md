
# Debug 

[![CircleCI](https://img.shields.io/circleci/project/github/AlexisVisco/debug.svg)](https://circleci.com/gh/AlexisVisco/debug)
[![Gocover](https://img.shields.io/badge/coverage-100.0%25-brightgreen.svg?style=flat)](https://gocover.io/github.com/AlexisVisco/debug)

A tiny golang debugging utility based on https://github.com/visionmedia/debug principles.

## Installation

`go get github.com/AlexisVisco/Debug`

## Usage

debug expose some simple functions like Register, Get, Delete to manage debug.

Example [http_debug.go](https://github.com/AlexisVisco/debug/blob/master/examples/http_debug.go)
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	debug "github.com/AlexisVisco/debug"
)

var httpdeb, _ = debug.Register("http")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	httpdeb.Log(fmt.Sprintf("%s %s", r.Method, r.URL.String()))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
![pic](https://i.imgur.com/CWJkgrv.jpg)

Example [infinite_debug.go](https://github.com/AlexisVisco/debug/blob/master/examples/http_debug.go)

```go
package main

import (
	debug "github.com/AlexisVisco/debug"
	"time"
	"strconv"
	"sync"
)

var fivesec, _ = debug.Register("5 times")
var nivesec, _ = debug.Register("9 times")
var wait sync.WaitGroup

var five = 0
var nine = 0

func main() {
	wait.Add(1)
	go doEvery(5 * time.Second, func(i time.Time) {
		fivesec.Log("5 = " + strconv.Itoa(five))
		five++
	})
	go doEvery(9 * time.Second, func(i time.Time) {
		nivesec.Log("9 = " + strconv.Itoa(nine))
		nine++
	})
	wait.Wait()
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
```
![pic](https://i.imgur.com/YX5lyQw.jpg)

The DEBUG environment variable is then used to enable these based on space or comma-delimited names.

## Wildcards
The `*` character may be used as a wildcard. Suppose for example your library has debuggers named "connect:bodyParser", "connect:compress", "connect:session", instead of listing all three with DEBUG=connect:bodyParser,connect:compress,connect:session, you may simply do DEBUG=connect:*.

## Exclusion

The `-` prefix character may be used to exclude a debugger.<br>
Example `DEBUG=*,-test` => atest OK, hello OK, test NOT OK<br>
You can combine with wildcard obviously !

## Environment Variables

You can set a few environment variables that will change the behavior of the debug logging:

| Name      | Purpose                                         |
|-----------|-------------------------------------------------|
| `DEBUG`   | Enables/disables specific debugging namespaces. |
| `DEBUG_HIDE_DATE` | Hide date from debug output (non-TTY).  |
| `DEBUG_COLORS`| Whether or not to use colors in the debug output. |
| `DEBUG_HIDE_LATENCY` | Hide latency at the end of a tty output. |

## Documentation

Functions:
* [`NewDebug(name string) *Debug`](#newdebug)
* [`Register(name string) (*Debug, Err)`](#register)
* [`Get(name string) (*Debug, Err)`](#get)
* [`Delete(name string) Err`](#delete)
* [`Enable()`](#enable)
* [`Disable()`](#disable)

Methods:
* [`(d *Debug) Log(message string)`](#log)
* [`(d *Debug) Sprint(message string)`](#sprint)
* [`(d *Debug) SetWriter(writer io.Writer, tty bool) *Debug`](#setwriter)
* [`(d *Debug) SetFdWriter(file *os.File) *Debug`](#setfdwriter)

## Functions

### NewDebug

__Prototype__: `NewDebug(name string) *Debug`<br/>

__Description__:<br/>
Create a debug structure without registering it. Cannot be accessible with [`Get`](#get).
Generate a random color from 31 to 37 and 91 to 97 as ainsi code.

```go
debug := debug.NewDebug("woaw") 
```

### Register

__Prototype__: `Register(name string) (*Debug, Err)` <br/>

__Description__:<br/>
Create a debug and registering it. Can be accessible with [`Get`](#get).
[`NewDebug`](#newdebug) is used to create the structure.

__Error__:<br/>
Return an error if name is already in the registry.

```go
debug, err := debug.Create("woaw")
if err {
  fmt.Printf("name %s already used !", "woaw")
}
```

### Get

__Prototype__: `Get(name string) (*Debug, Err)`<br/>

__Description__:<br/>
Get a debug structure from it name.

__Error__:<br/>
Return an error if name is not in the registry.

```go
debug, err := debug.Get("woaw")
if err {
  fmt.Printf("name %s has not been created !", "woaw")
}
```

### Delete

__Prototype__: `Delete(name string) Err` <br/>

__Description__:<br/>
Delete a debug structure from the registry.

__Error__:<br/>
Return an error if name is not in the registry.

```go
err := debug.Delete("woaw")
if err {
  fmt.Printf("name %s has not been created !", "woaw")
}
```

### Enable

__Prototype__: `Enable()` <br/>

__Description__:<br/>
Enable printing with debug.

```go
debug.Enable()
```

### Disable

__Prototype__: `Disable()`<br/>

__Description__:<br/>
Disable printing with debug.

```go
debug.Disable()
```

## Methods

### Log

__Prototype__: `(d *Debug) Log(message string`<br/>

__Description__:<br/>
Print if debug is active the message with the name of the debug and the latency between the last call if it was activated.

```go
woaw, _ := debug.Create("woaw")

woaw.Log("Hola !")
woaw.Log("Hola 2 !")
```

### Sprint

__Prototype__: `(d *Debug) Sprint(message string)` <br/>

__Description__:<br/>
Return the full string that should be printed.

```go
woaw, _ := debug.Create("woaw")

str := waw.Sprint("Hola !")
```

### SetWriter

__Prototype__: `(d *Debug) SetWriter(writer io.Writer, tty bool) *Debug` <br/>

__Description__:<br/>
Set the writer, if it's a terminal set to true the next parameter.

```go
woaw, _ := debug.Create("woaw")

woaw.SetWriter(os.Stdout, true)
```

### SetFdWriter

__Prototype__: `(d *Debug) SetFdWriter(file *os.File) *Debug` <br/>

__Description__:<br/>
This function will set the writer and determine if the `file.Fd()` is a terminal.

```go
woaw, _ := debug.Create("woaw")

woaw.SetFdWriter(os.Stdout)
```
