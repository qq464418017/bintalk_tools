package excel

import (
	bb "btn_tools/btk"
	bFile "btn_tools/bFile"
	"btn_tools/enum"
	"btn_tools/config"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var btk *bFile.BTFile

func ReadExcel() {
	if len(config.Data.ExcelPath) <= 0 {
		return
	}

	bytesPath := config.Path + "/generadir/bytes/"
	err := os.Mkdir(bytesPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	newDataTableFile()
	defer btk.Close()
	defer enum.Close()
	fmt.Print(config.Path + config.Data.ExcelPath)
	files, err := ioutil.ReadDir(config.Path + config.Data.ExcelPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".DS_Store"){
			continue
		}
		xlsx, xe := excelize.OpenFile(config.Path + config.Data.ExcelPath + f.Name())
		if xe != nil {
			fmt.Println(xe, f.Name())
			continue
		}
		rows, re := xlsx.GetRows("Sheet1")
		if re != nil {
			fmt.Println(re)
			continue
		}
		if rows == nil {
			rows, re = xlsx.GetRows("Sheet")
		}
		name := strings.Split(f.Name(), ".")[0]
		writebtk(rows[1], rows[2], name)
		writeEnum(rows, name)
		writeBytes(rows, name)
	}
	fmt.Println("excel end")
}

func writebtk(row2 []string, row3 []string, name string) {

	btk.Write("struct " + name)
	btk.Write("{")
	for i := 0; i < len(row2); i++ {
		btk.WriteTable(row3[i] + " " + row2[i] + ";")
	}
	btk.Write("}")
	btk.WrteNewLine()
	btk.Write("struct " + name + "_table")
	btk.Write("{")
	btk.WriteTable(name + " []record;")
	btk.Write("}")
	btk.WrteNewLine()

}

func writeEnum(rows [][]string, name string) {
	flag := false
	for index := 0; index < len(rows); index++ {
		row := rows[index]
		if index == 1 && row[1] == "codeName" {
			//enum.Write("enum E_" + name)
			//enum.Write("{")
			enum.WriteEnumName(name)
			flag = true
		}
		if index > 2 && flag {
			enum.WriteEnumItem(row[1], row[0], (index + 1) == len(rows))
		}
	}
	if flag {
		enum.WriteEnd()
	}
}

func writeBytes(rows [][]string, name string) {
	file, er := os.Create(config.Path + "/generadir/bytes/" + name + "_table.bytes")
	defer file.Close()
	if er != nil {
		fmt.Println(er)
		return
	}

	//r1 := rows[1]
	r2 := rows[2]
	p := bb.ProtocolWriter{Bytes: make([]byte, 0)}
	v := len(rows) - 3
	p.WriteDynSize(v)
	for index := 0; index < len(rows); index++ {
		if index > 2 {
			row := rows[index]
			for i := 0; i < len(row); i++ {
				value := strings.TrimSpace(row[i])
				//p.WriteData(value)

				switch r2[i] {
				case "string":
					p.WriteString(value)
				case "uint8":
					s, e := strconv.ParseUint(value, 0, 8)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(uint8(s))
					}
				case "uint16":
					s, e := strconv.ParseUint(value, 0, 16)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(uint16(s))
					}
				case "uint32":
					s, e := strconv.ParseUint(value, 0, 32)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(uint32(s))
					}
				case "uint64":
					s, e := strconv.ParseUint(value, 0, 64)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(uint64(s))
					}
				case "int8":
					s, e := strconv.ParseInt(value, 0, 8)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(int8(s))
					}
				case "int16":
					s, e := strconv.ParseInt(value, 0, 16)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(int16(s))
					}
				case "int32":
					s, e := strconv.ParseInt(value, 0, 32)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(int32(s))
					}
				case "int64":
					s, e := strconv.ParseInt(value, 0, 64)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(int64(s))
					}
				case "bool":
					s, e := strconv.ParseBool(value)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(s)
					}
				case "float":
					s, e := strconv.ParseFloat(value, 32)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(float32(s))
					}
					//fmt.Println(fmt.Sprintf("resource value: %d,  parse value: %f", value, s))
				case "double":
					s, e := strconv.ParseFloat(value, 64)
					if e != nil {
						fmt.Println(e)
					} else {
						p.WriteSingle(float64(s))
					}
				default:
					fmt.Println(r2[i])
				}
			}
		}
	}
	file.Write(p.Bytes)
}

func newDataTableFile() {
	btk = bFile.NewBtkFile(config.GenPath + "DataTables.btk")
}
