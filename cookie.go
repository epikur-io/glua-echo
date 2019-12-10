package lsp

import (
	"net/http"
	"time"

	lua "github.com/yuin/gopher-lua"
)


func LCookie_NewUD(L *lua.LState, gs *http.Cookie) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("cookie"))
	return ud
}

func check_LCookie(L *lua.LState, index int) *http.Cookie {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*http.Cookie); ok {

		return v
	}
	L.ArgError(1, "lsp.cookie object expected")
	return nil
}

func LCookie_Domain(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckString(1)
		c.Domain = d
		return 0
	}
	L.Push(lua.LString(c.Domain))
	return 1
}

func LCookie_HttpOnly(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckBool(1)
		c.HttpOnly = d
		return 0
	}
	L.Push(lua.LBool(c.HttpOnly))
	return 1
}

func LCookie_Name(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckString(1)
		c.Name = d
		return 0
	}
	L.Push(lua.LString(c.Name))
	return 1
}

func LCookie_Path(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckString(1)
		c.Path = d
		return 0
	}
	L.Push(lua.LString(c.Path))
	return 1
}

func LCookie_MaxAge(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckInt(1)
		c.MaxAge = d
		return 0
	}
	L.Push(lua.LNumber(c.MaxAge))
	return 1
}

func LCookie_Secure(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckBool(1)
		c.Secure = d
		return 0
	}
	L.Push(lua.LBool(c.Secure))
	return 1
}

func LCookie_Value(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		d := L.CheckString(1)
		c.Value = d
		return 0
	}
	L.Push(lua.LString(c.Value))
	return 1
}

func LCookie_Expires(L *lua.LState) int {
	c := check_LCookie(L, 1)
	if L.GetTop() > 0 {
		ud := L.CheckUserData(1)
		t, ok := ud.Value.(time.Time)
		if !ok {
			return 0
		}
		c.Expires = t
		return 0
	}
	L.Push(lua.LString(c.Expires.String()))
	return 1
}

func LCookie_Table(L *lua.LState) int {
	c := check_LCookie(L, 1)
	t := L.NewTable()
	t.RawSetH(lua.LString("domain"), lua.LString(c.Domain))
	t.RawSetH(lua.LString("path"), lua.LString(c.Path))
	t.RawSetH(lua.LString("name"), lua.LString(c.Name))
	t.RawSetH(lua.LString("secure"), lua.LBool(c.Secure))
	t.RawSetH(lua.LString("httpOnly"), lua.LBool(c.HttpOnly))
	t.RawSetH(lua.LString("maxAge"), lua.LNumber(c.MaxAge))
	t.RawSetH(lua.LString("expires"), lua.LString(c.Expires.String()))
	t.RawSetH(lua.LString("value"), lua.LString(c.Value))
	L.Push(lua.LString(c.Value))
	return 1
}
