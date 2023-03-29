package main

import (
	_ "btn_tools/config"
	_ "btn_tools/enum"
	"btn_tools/excel"
	"btn_tools/gen"
)
func main() {
	excel.ReadExcel()
	gen.Gen()
	gen.Copy()
//readBtk.Read("")
//	_, e := os.Stat(config.GenPath)
//	if e == nil {
//		os.RemoveAll(config.GenPath)
//	}
}
