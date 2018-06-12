
# Debug
A tiny golang debugging utility modelled after go core's debugging technique.
Works in go with any things that implement io.Writer.

## Installation

`go get github.com/alexisvisco/Debug`

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