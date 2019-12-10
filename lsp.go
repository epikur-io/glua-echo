package lsp

import (
	"errors"

	echo "github.com/labstack/echo/v4"
	lua "github.com/yuin/gopher-lua"
)

func LContext_NewUD(L *lua.LState, gs echo.Context) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("context"))
	return ud
}

func check_LContext(L *lua.LState, index int) echo.Context {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(echo.Context); ok {
		return v
	}
	L.ArgError(1, "lsp.context object expected")
	return nil
}

func LContext_NewFromUD(L *lua.LState) int {
	ud := L.CheckUserData(1)
	ctx, ok := ud.Value.(echo.Context)
	if ok {
		L.Push(LContext_NewUD(L, ctx))
		return 1
	}

	L.Push(lua.LNil)
	return 1
}

func LContext_Blob(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	contentType := L.CheckString(3)
	data := L.CheckString(4)
	err := ctx.Blob(code, contentType, []byte(data))
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

// echo.HTTPError
func LContext_HttpError(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	msg := L.CheckString(3)
	err := &echo.HTTPError{
		Code:    code,
		Message: msg,
	}
	ctx.Error(err)
	return 0
}
func LContext_NoContent(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	err := ctx.NoContent(code)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LContext_Json(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	data := L.CheckString(3)
	err := ctx.JSONBlob(code, []byte(data))
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LContext_Jsonp(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	callback := L.CheckString(3)
	data := L.CheckString(4)
	err := ctx.JSONPBlob(code, callback, []byte(data))
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LContext_Error(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	errMsg := L.CheckString(2)
	ctx.Error(errors.New(errMsg))
	return 0
}

func LContext_FormValue(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	n := L.CheckString(2)

	L.Push(lua.LString(ctx.FormValue(n)))
	return 1
}

func LContext_Param(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	n := L.CheckString(2)
	L.Push(lua.LString(ctx.QueryParam(n)))
	return 1
}

func LContext_Scheme(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	L.Push(lua.LString(ctx.Scheme()))
	return 1
}

func LContext_RealIP(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	L.Push(lua.LString(ctx.RealIP()))
	return 1
}

func LContext_Host(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	L.Push(lua.LString(ctx.Request().Host))
	return 1
}

func LContext_Method(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	L.Push(lua.LString(ctx.Request().Method))
	return 1
}

func LContext_FormFile(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	n := L.CheckString(2)
	rf, err := ctx.FormFile(n)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(LFormFile_NewUD(L, rf))
	L.Push(lua.LNil)
	return 2
}

func LContext_GetCookie(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	n := L.CheckString(2)
	c, err := ctx.Cookie(n)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(LCookie_NewUD(L, c))
	L.Push(lua.LNil)
	return 2
}

func LContext_GetCookies(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	cs := ctx.Cookies()
	t := L.NewTable()
	for _, c := range cs {
		t.Append(LCookie_NewUD(L, c))
	}
	L.Push(t)
	return 1
}

func LContext_IsTLS(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	n := L.CheckString(2)
	c, err := ctx.Cookie(n)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(LCookie_NewUD(L, c))
	L.Push(lua.LNil)
	return 2
}

func LContext_String(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	n := L.CheckString(3)
	err := ctx.String(code, n)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LContext_HTML(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	code := L.CheckInt(2)
	n := L.CheckString(3)
	err := ctx.HTML(code, n)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LContext_Response(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	L.Push(LResponse_NewUD(L, ctx.Response()))
	return 1
}

func LContext_Request(L *lua.LState) int {
	ctx := check_LContext(L, 1)
	L.Push(LRequest_NewUD(L, ctx.Request()))
	return 1
}

func LContext_MultipartForm(L *lua.LState) int {
	ctx := check_LContext(L, 1)

	mf, err := ctx.MultipartForm()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(LMultipartForm_NewUD(L, mf))
	L.Push(lua.LNil)
	return 2
}
