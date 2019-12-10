package lsp

import (
	echo "github.com/labstack/echo/v4"

	lua "github.com/yuin/gopher-lua"
)

// !TODO

func LResponse_NewUD(L *lua.LState, gs *echo.Response) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("response"))
	return ud
}

func check_LResponse(L *lua.LState, index int) *echo.Response {

	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*echo.Response); ok {

		return v
	}
	L.ArgError(1, "lsp.response object expected")
	return nil
}

func LResponse_Committed(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	L.Push(lua.LBool(resp.Committed))
	return 1
}

func LResponse_Size(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	L.Push(lua.LNumber(resp.Size))
	return 1
}

func LResponse_SetHeader(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	n := L.CheckString(2)
	v := L.CheckString(3)
	resp.Header().Set(n, v)
	return 0
}

func LResponse_DelHeader(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	n := L.CheckString(2)
	resp.Header().Del(n)
	return 0
}

func LResponse_GetHeader(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	n := L.CheckString(2)
	L.Push(lua.LString(resp.Header().Get(n)))
	return 1
}

func LResponse_WriteHeader(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	code := L.CheckInt(2)
	resp.WriteHeader(code)
	return 0
}

func LResponse_Status(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	L.Push(lua.LNumber(resp.Status))
	return 1
}

func LResponse_Flush(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	resp.Flush()
	return 0
}

func LResponse_Write(L *lua.LState) int {
	resp := check_LResponse(L, 1)
	data := L.CheckString(2)
	n, err := resp.Write([]byte(data))
	if err != nil {
		L.Push(lua.LNumber(0))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LNumber(n))
	L.Push(lua.LNil)
	return 2
}
