[![GoDoc](https://godoc.org/github.com/mtabini/goluago?status.png)](https://godoc.org/github.com/mtabini/goluago)

Lua wrappers for the Go standard library
========================================

Wraps Go's standard libraries so they can be used with the [go-lua](https://github.com/mtabini/go-lua) Lua VM implementation.

Most of the packages under `pkg` expose a single function `Open(l *lua.State)` that makes the corresponding Go package available to Lua scripts. For example:
```go
import "github.com/mtabini/goluago/pkg/strings"
...
strings.Open(l)
...
```
allows Lua scripts loaded by `l` to:
```lua
require("goluago/strings")
strings.trim("loll ")
strings.split("cat,dog,elephant,walrus", ",")
strings.replace("oink oink oink", "k", "ky", 2)
```

To make all supported APIs available to Lua:
```go
import "github.com/mtabini/goluago"
...
goluago.Open(l)
...
```

The `"github.com/mtabini/goluago/util"` package provides helper functions to push Go values onto the Lua stack, pull tables of string->string or string->value from the Lua stack to Go, and support variadic arguments.

License
-------

goluago is licensed under the [MIT license](https://github.com/mtabini/goluago/blob/master/LICENSE.md).
