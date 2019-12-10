# Lua bindings to labstack's echo library

[Gopher-lua](https://github.com/yuin/gopher-lua) bindings to labstack's [echo](https://echo.labstack.com/) library.

## Notes

These bindings don't support echo's routing or middlerware functionallity.
You have to register your routes and middlewares from inside go.
This library is intended to provide the ability to run custom Lua scripts on top of echo.

## Example

```go

    import (
        gecho "github.com/epikur-io/glua-echo"
    )

    e := echo.New()

    // Echo setup. Bind routes and middlerwares...

    e.GET("/", func(c echo.Context) error {
        L := lua.NewState()
        L.PreloadModule("echo", gecho.Loader)
        L.DoString(`
            local echo = require("echo")
        `)

        // !TODO: Add example. Show how to bind echo's context to lua

        return
    })

```

```lua
    -- !TODO
```
