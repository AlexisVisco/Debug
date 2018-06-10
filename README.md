
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

Methods (todo):
* [`(d *Debug) Print(message string)`](#Print)
* [`(d *Debug) Sprint(message string)`](#Print)

<a name="NewDebug" />
NewDebug(name string) *Debug

---

__Description__:<br/>
Create a debug structure without registering it. Cannot be accessible with [`Get`](#Get).
Generate a random color from 31 to 37 and 91 to 97 as ainsi code.

```go
debug := debug.NewDebug("woaw") 
```

<a name="Register" />
Register(name string) (*Debug, Err)

---

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

<a name="Get" />
Get(name string) (*Debug, Err)

---

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

<a name="Delete" />
Delete(name string) Err

---

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

<a name="Enable" />
Enable()

---

__Description__:<br/>
Enable printing with debug.

```go
debug.Enable()
```

<a name="Disable" />
Disable()

---

__Description__:<br/>
Disable printing with debug.

```go
debug.Disable()
```
