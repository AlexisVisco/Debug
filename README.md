
# Debug
A tiny golang debugging utility modelled after go core's debugging technique.
Works in go with any things that implement io.Writer.

## Installation

`go get github.com/alexisvisco/Debug`

## Documentation

Functions:
* [`NewDebug(name string) *Debug`](#NewDebug)
* [`Register(name string) (*Debug, Err)`](#Register)
* [`Get(name string) (*Debug, Err)`](#Get)
* [`Delete(name string) Err`](#Delete)
* [`Enable()`](#Enable)
* [`Disable()`](#Disable)

Methods:
* [`(d *Debug) Print(message string)`](#Print)
* [`(d *Debug) Sprint(message string)`](#Sprint)

## Functions

### NewDebug

__Prototype__: `NewDebug(name string) *Debug`<br/>

__Description__:<br/>
Create a debug structure without registering it. Cannot be accessible with [`Get`](#Get).
Generate a random color from 31 to 37 and 91 to 97 as ainsi code.

```go
debug := debug.NewDebug("woaw") 
```

### Register

__Prototype__: `Register(name string) (*Debug, Err)` <br/>

__Description__:<br/>
Create a debug and registering it. Can be accessible with [`Get`](#Get).
[`NewDebug`](#NewDebug) is used to create the structure.

__Error__:<br/>
Return an error if name is already in the registery.

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
Return an error if name is not in the registery.

```go
debug, err := debug.Get("woaw")
if err {
  fmt.Printf("name %s has not been created !", "woaw")
}
```

### Delete

__Prototype__: `Delete(name string) Err` <br/>

__Description__:<br/>
Delete a debug structure from the registery.

__Error__:<br/>
Return an error if name is not in the registery.

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

### Print

__Prototype__: `(d *Debug) Print(message string`<br/>

__Description__:<br/>
Print if debug is active the message with the name of the debug and the latency between the last call if it was activated.

```go
woaw, _ := debug.Create("woaw")

woaw.Print("Hola !")
woaw.Print("Hola 2 !")
```

### Sprint

__Prototype__: `(d *Debug) Sprint(message string)` <br/>

__Description__:<br/>

```go
woaw, _ := debug.Create("woaw")

str := waw.Sprint("Hola !")
```
