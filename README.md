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
* [`(d *Debug) Sprint(message string)`](#Print)

<a name="NewDebug"/>
### NewDebug(name string) *Debug

*Description*:
Create a debug structure without registering it. Cannot be accessible with [`Get`](#Get).
Generate a random color from 31 to 37 and 91 to 97 as ainsi code.

```go
debug := debug.NewDebug("woaw") 
```

<a name="Register"/>
### Register(name string) (*Debug, Err)

*Description*:
Create a debug and registering it. Can be accessible with [`Get`](#Get).
[`NewDebug`](#NewDebug) is used to create the structure.

*Error*:
Return an error if name is already in the registery.

```go
debug, err := debug.Create("woaw")
if err {
  fmt.Printf("name %s already used !", "woaw")
}
```

<a name="Get"/>
### Get(name string) (*Debug, Err)

*Description*:
Get a debug structure from it name.

*Error*:
Return an error if name is not in the registery.

```go
debug, err := debug.Get("woaw")
if err {
  fmt.Printf("name %s has not been created !", "woaw")
}
```

<a name="Delete"/>
### Delete(name string) Err

*Description*:
Delete a debug structure from the registery.

*Error*:
Return an error if name is not in the registery.

```go
err := debug.Delete("woaw")
if err {
  fmt.Printf("name %s has not been created !", "woaw")
}
```
