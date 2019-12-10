package lsp

import (
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

// !TODO

func LRequest_NewUD(L *lua.LState, gs *http.Request) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("request"))
	return ud
}

func check_LRequest(L *lua.LState, index int) *http.Request {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*http.Request); ok {

		return v
	}
	L.ArgError(1, "lsp.request object expected")
	return nil
}

func LRequest_GetHeader(L *lua.LState) int {
	req := check_LRequest(L, 1)
	n := L.CheckString(2)
	r := req.Header.Get(n)
	L.Push(lua.LString(r))
	return 1
}

func LRequest_SetHeader(L *lua.LState) int {
	req := check_LRequest(L, 1)
	n := L.CheckString(2)
	v := L.CheckString(3)
	req.Header.Set(n, v)
	return 0
}

func LRequest_DelHeader(L *lua.LState) int {
	req := check_LRequest(L, 1)
	n := L.CheckString(2)
	req.Header.Del(n)
	return 0
}

func LRequest_ReadBody(L *lua.LState) int {
	req := check_LRequest(L, 1)
	size := L.CheckInt64(2)
	buf := make([]byte, size)
	_, err := req.Body.Read(buf)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(string(buf)))
	L.Push(lua.LNil)
	return 2
}

func LRequest_CloseBody(L *lua.LState) int {
	req := check_LRequest(L, 1)
	defer req.Body.Close()
	err := req.Body.Close()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LRequest_Host(L *lua.LState) int {
	req := check_LRequest(L, 1)
	L.Push(lua.LString(req.Host))
	return 1
}

func LRequest_URI(L *lua.LState) int {
	req := check_LRequest(L, 1)
	L.Push(lua.LString(req.RequestURI))
	return 1
}

func LRequest_UserAgent(L *lua.LState) int {
	req := check_LRequest(L, 1)
	L.Push(lua.LString(req.UserAgent()))
	return 1
}

func LRequest_Method(L *lua.LState) int {
	req := check_LRequest(L, 1)
	L.Push(lua.LString(req.Method))
	return 1
}

func LRequest_URL(L *lua.LState) int {
	req := check_LRequest(L, 1)
	L.Push(LUrl_NewUD(L, req.URL))
	return 1
}

func LRequest_RemoteAddr(L *lua.LState) int {
	req := check_LRequest(L, 1)
	L.Push(lua.LString(req.RemoteAddr))
	return 1
}
