package lsp

import (
	purl "net/url"

	lua "github.com/yuin/gopher-lua"
)

// !TODO

func LUrl_NewUD(L *lua.LState, gs *purl.URL) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("url"))
	return ud
}

func check_LUrl(L *lua.LState, index int) *purl.URL {

	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*purl.URL); ok {

		return v
	}
	L.ArgError(1, "lsp.url object expected")
	return nil
}

func LUrl_New(L *lua.LState) int {
	var u *purl.URL
	var v string
	var err error
	v = L.CheckString(1)
	u, err = purl.Parse(v)
	if err != nil {
		L.Push(LUrl_NewUD(L, &purl.URL{}))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(LUrl_NewUD(L, u))
	L.Push(lua.LNil)
	return 2
}

func LUrl_EscapedPath(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.EscapedPath()))
	return 1
}

func LUrl_Fragment(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.Fragment))
	return 1
}

func LUrl_Param(L *lua.LState) int {
	url := check_LUrl(L, 1)
	n := L.CheckString(2)
	L.Push(lua.LString(url.Query().Get(n)))
	return 1
}

func LUrl_DelParam(L *lua.LState) int {
	url := check_LUrl(L, 1)
	n := L.CheckString(2)
	url.Query().Del(n)
	return 0
}

func LUrl_SetParam(L *lua.LState) int {
	url := check_LUrl(L, 1)
	n := L.CheckString(2)
	v := L.CheckString(3)
	url.Query().Set(n, v)
	return 0
}

func LUrl_Username(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.User.Username()))
	return 1
}

func LUrl_Password(L *lua.LState) int {
	url := check_LUrl(L, 1)
	pw, _ := url.User.Password()
	L.Push(lua.LString(pw))
	return 1
}

func LUrl_Scheme(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.Scheme))
	return 1
}

func LUrl_Opaque(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.Opaque))
	return 1
}

func LUrl_Path(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.Path))
	return 1
}

func LUrl_IsAbs(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LBool(url.IsAbs()))
	return 1
}

func LUrl_Host(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.Host))
	return 1
}

func LUrl_Query(L *lua.LState) int {
	url := check_LUrl(L, 1)
	L.Push(lua.LString(url.Query().Encode()))
	return 1
}
