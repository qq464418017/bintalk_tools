package orm

import (
	File "btn_tools/bFile"
	"btn_tools/config"
	"btn_tools/util"
	"strings"
)

func (d *DbClient)WriteGoStruct(){
	btk := File.NewBtkFile(config.GenPath + d.database + ".go")
	btk.Write("package bintalk")
	btk.WrteNewLine()
	sqldb, err := d.db.DB()
	if err != nil {
		println("sql db error ", err)
		return
	}
	result, err := sqldb.Query("SHOW TABLES")
	if err != nil {
		panic(err)
		return
	}
	for result.Next() {
		var v string
		result.Scan(&v)
		btk.Write("type " + util.StrFirstToUpper(v) + " struct {")
		var v1, v2, v3, v4, v5, v6 string
		r, e := sqldb.Query("desc " + v)
		if e != nil {
			panic(err)
			return
		}
		index := 0
		for r.Next() {
			r.Scan(&v1, &v2, &v3, &v4, &v5, &v6)
			//fmt.Println(v1, v2, v3, v4, v5, v6)

			btk.WriteTable(getParam(v1, v2, index))
			index ++
		}
		btk.Write("}")
		btk.WrteNewLine()
	}
	btk.Close()
}

func getParam(v1, v2 string, index int) string {
	pType := getParamDataByGo(v2)
	if index == 0 {
		return util.StrFirstToUpper(v1) + " " + pType + "  `" + "gorm:\"primary_key\"" + "`"
	} else {
		return util.StrFirstToUpper(v1) + " " + pType
	}
}

func getParamDataByGo(v2 string) string {
	var pType string
	if strings.Contains(v2, "int") {
		s := strings.Split(v2, " ")
		var isUint bool = false
		if len(s) > 1 {
			if s[1] == "unsigned" {
				isUint = true
			}
		}
		var value string
		if strings.Contains(s[0], "big") {
			value = "64"
		} else if strings.Contains(s[0], "tiny") {
			value = "8"
		} else if strings.Contains(s[0], "small") {
			value = "16"
		} else {
			value = "32"
		}
		if isUint {
			pType = "uint" + value
		} else {
			pType = "int" + value
		}

	} else if strings.Contains(v2, "varchar") {
		pType = "string"
	} else if strings.Contains(v2, "text") {
		pType = "string"
	} else if strings.Contains(v2, "datetime") {
		pType = "string"
	} else if strings.Contains(v2, "date") {
		pType = "string"
	} else if strings.Contains(v2, "blob") {
		pType = "[]uint8"
	} else if strings.Contains(v2, "float") {
		pType = "float32"
	} else if strings.Contains(v2, "enum"){
		pType = "string"
	}
	return pType
}
