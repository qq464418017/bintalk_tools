package orm

import (
	File "btn_tools/bFile"
	"btn_tools/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type DbClient struct {
	host     string
	user     string
	pass     string
	database string
	db       *gorm.DB
}

func NewDBClient(host, user, pass, database string) *DbClient {
	d := new(DbClient)
	d.host = host
	d.user = user
	d.pass = pass
	d.database = database
	return d
}

func (d *DbClient) Start() {
	dsn := d.user+":"+d.pass+"@tcp("+d.host+":3306)/"+d.database+"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	})
	if err != nil {
		fmt.Println("databases connect failed")
		return
	} else {
		fmt.Println("databases connect success")
	}
	sqldb, err := db.DB()
	if err != nil {
		println("sql db error ", err)
		return
	}
	err = sqldb.Ping()
	if err != nil {
		println("db ping error")
	} else {
		println("db ping success")
		d.db = db
	}
	//defer sqldb.Close()
	//d.selectDatabase()
}

func (d *DbClient) WriteBtk() {
	btk := File.NewBtkFile(config.GenPath + d.database + ".btk")
	//result, err := d.db.CommonDB().Query("SHOW TABLES")
	sqldb, err := d.db.DB()
	if err != nil {
		println("sql db error ", err)
		return
	}
	result, err := sqldb.Query("SHOW TABLES")
	//result, err = d.db.Statement.
	if err != nil {
		fmt.Print(err)
		panic(err)
		return
	}
	for result.Next(){
		var v string
		result.Scan(&v)
		btk.Write("struct " + v)
		btk.Write("{")
		var v1, v2, v3, v4, v5, v6 string
		r, e := sqldb.Query("desc " + v)
		if e != nil {
			panic(err)
			return
		}
		for r.Next() {
			r.Scan(&v1, &v2, &v3, &v4, &v5, &v6)
			//fmt.Println(v1, v2, v3, v4, v5, v6)
			btk.WriteTable(getParamData(v2) + "   " + v1 + ";")
		}
		btk.Write("}")
		btk.WrteNewLine()
	}
	btk.Close()
}

func getParamData(v2 string) string {
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
		} else if strings.Contains(s[0], "smaill") {
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
		pType = "uint8[]"
	} else if strings.Contains(v2, "float") {
		pType = "float"
	} else if strings.Contains(v2, "enum"){
		pType = "string"
	}
	return pType
}
