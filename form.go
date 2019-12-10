package lsp

import (
	//"github.com/labstack/echo"
	"mime/multipart"

	lua "github.com/yuin/gopher-lua"
)

func LMultipartForm_NewUD(L *lua.LState, gs *multipart.Form) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("multipartForm"))
	return ud
}

func check_LMultipartForm(L *lua.LState, index int) *multipart.Form {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*multipart.Form); ok {
		return v
	}
	L.ArgError(1, "lsp.formFile object expected")
	return nil
}

func LFormFile_NewUD(L *lua.LState, gs *multipart.FileHeader) *lua.LUserData {
	ud := L.NewUserData()
	f := &LFormFile{
		FileHeader: gs,
	}
	ud.Value = f
	L.SetMetatable(ud, L.GetTypeMetatable("formFile"))
	return ud
}

type LFormFile struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
}

func check_LFormFile(L *lua.LState, index int) *LFormFile {
	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*LFormFile); ok {
		return v
	}
	L.ArgError(1, "lsp.formFile object expected")
	return nil
}

func LMultipartForm_Files(L *lua.LState) int {
	mf := check_LMultipartForm(L, 1)
	t := L.NewTable()
	for field, fileList := range mf.File {
		t.RawSetH(lua.LString(field), L.NewTable())
		fl := t.RawGetH(lua.LString(field))
		if fl.Type() != lua.LTTable {
			continue
		}
		files, _ := fl.(*lua.LTable)
		for _, f := range fileList {
			files.Append(LFormFile_NewUD(L, f))
		}
	}
	L.Push(t)
	return 1
}

func LMultipartForm_RemoveAll(L *lua.LState) int {
	mf := check_LMultipartForm(L, 1)
	err := mf.RemoveAll()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LFormFile_Read(L *lua.LState) int {
	lf := check_LFormFile(L, 1)
	size := L.CheckInt64(2)
	if lf.File == nil {
		var err error
		lf.File, err = lf.FileHeader.Open()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
	}
	buf := make([]byte, size)
	n, err := lf.File.Read(buf)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LNumber(n))
	L.Push(lua.LNil)
	return 2
}

func LFormFile_ReadAt(L *lua.LState) int {
	lf := check_LFormFile(L, 1)
	size := L.CheckInt64(2)
	offset := L.CheckInt64(3)
	if lf.File == nil {
		var err error
		lf.File, err = lf.FileHeader.Open()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
	}
	buf := make([]byte, size)
	n, err := lf.File.ReadAt(buf, offset)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LNumber(n))
	L.Push(lua.LNil)
	return 2
}

func LFormFile_Seek(L *lua.LState) int {
	lf := check_LFormFile(L, 1)
	offset := L.CheckInt64(2)
	whence := L.CheckInt(3)
	if lf.File == nil {
		var err error
		lf.File, err = lf.FileHeader.Open()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
	}
	n, err := lf.File.Seek(offset, whence)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LNumber(n))
	L.Push(lua.LNil)
	return 2
}

func LFormFile_Close(L *lua.LState) int {
	lf := check_LFormFile(L, 1)
	if lf.File == nil {
		var err error
		lf.File, err = lf.FileHeader.Open()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}
	}
	err := lf.File.Close()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LNil)
	return 1
}

func LFormFile_FileName(L *lua.LState) int {
	lf := check_LFormFile(L, 1)
	L.Push(lua.LString(lf.FileHeader.Filename))
	return 1
}

func LFormFile_Size(L *lua.LState) int {
	lf := check_LFormFile(L, 1)
	L.Push(lua.LNumber(lf.FileHeader.Size))
	return 1
}
