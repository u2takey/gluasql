package gluasql

import (
	mysql "github.com/u2takey/gluasql/mysql"
	sqlite3 "github.com/u2takey/gluasql/sqlite3"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("mysql", mysql.Loader)
	L.PreloadModule("sqlite3", sqlite3.Loader)
}
