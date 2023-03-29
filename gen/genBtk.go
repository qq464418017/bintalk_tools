package gen

import (
	"btn_tools/config"
	"btn_tools/exec"
	db "btn_tools/orm"
	"btn_tools/util"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Gen() {
	if len(config.Data.DataBase) > 0 {
		generaDatabase()
	}
	generaGO()
	generaCS()
	generaJava()
	generaLua()
	generaErl()
	generaDart()

	fmt.Println("genera end")
}

func Copy() {
	copyGo()
	copyCs()
	copyJava()
	copyErl()
	copyLua()
	copyDart()
	fmt.Println("copy end")
}

func generaDatabase() {
	for _, v := range config.Data.DataBase {
		d := db.NewDBClient(v.Host, v.User, v.Password, v.Database)
		d.Start()
		d.WriteBtk()
		if config.Data.Go.NeeData {
			d.WriteGoStruct()
		}
	}
}

func generaGO() {
	if len(config.Data.Go.GenList) <= 0 || !config.Data.Go.IsOpen {
		return
	}
	if len(config.Data.ExcelPath) > 0 {
		//exec.Genrate2("go", config.GenPath+"/DataTableEnums.btk", config.GenPath, "-gp", config.Data.Go.PackageName)
		exec.Genrate2("go", config.GenPath+"/DataTables.btk", config.GenPath, "-gp", config.Data.Go.PackageName)
	}
	for _, v := range config.Data.Go.GenList {
		exec.Genrate2("go", config.Path+"/"+v, config.GenPath, "-gp", config.Data.Go.PackageName)
		//exec.Genrate("go", config.Path+"/"+v, config.GenPath)
	}
}

func generaDart() {
	if len(config.Data.Dart.GenList) <= 0 || !config.Data.Dart.IsOpen {
		return
	}
	if len(config.Data.ExcelPath) > 0 {
		println("Gen dart excel")
		exec.Genrate2("dart", config.GenPath+"/DataTables.btk", config.GenPath, "-dp", config.Data.Dart.PackageName)
	}
	for _, v := range config.Data.Dart.GenList {
		println("dart: " + v)
		exec.Genrate2("dart", config.Path+"/"+v, config.GenPath, "-dp", config.Data.Dart.PackageName)
	}
}

func generaCS() {
	if len(config.Data.Cs.GenList) <= 0 || !config.Data.Cs.IsOpen {
		return
	}
	if len(config.Data.ExcelPath) > 0 {
		//exec.Genrate("cs", config.GenPath+"/DataTableEnums.btk", config.GenPath)
		exec.Genrate("cs", config.GenPath+"/DataTables.btk", config.GenPath)
	}
	for _, v := range config.Data.Cs.GenList {
		exec.Genrate("cs", config.Path+"/"+v, config.GenPath)
	}
}

func generaJava() {
	if len(config.Data.Java.GenList) <= 0 || !config.Data.Java.IsOpen {
		return
	}
	if len(config.Data.ExcelPath) > 0 {
		//exec.Genrate2("java", config.GenPath+"/DataTableEnums.btk", config.GenPath, "-jp",config.Data.Java.PackageName)
		exec.Genrate2("java", config.GenPath+"/DataTables.btk", config.GenPath, "-jp", config.Data.Java.PackageName)
	}
	for _, v := range config.Data.Java.GenList {
		exec.Genrate2("java", config.Path+"/"+v, config.GenPath, "-jp", config.Data.Java.PackageName)
	}
}

func generaLua() {
	if len(config.Data.Lua.GenList) <= 0 || !config.Data.Lua.IsOpen {
		return
	}
	if len(config.Data.ExcelPath) > 0 {
		//exec.Genrate("lua", config.GenPath+"/DataTableEnums.btk", config.GenPath)
		exec.Genrate("lua", config.GenPath+"/DataTables.btk", config.GenPath)
	}
	for _, v := range config.Data.Lua.GenList {
		exec.Genrate("lua", config.Path+"/"+v, config.GenPath)
	}
}

func generaErl() {
	//if len(config.Data.E.GenList) <= 0 {
	//	return
	//}
	//if len(config.Data.ExcelPath) > 0 {
	//	exec.Genrate("erl", config.GenPath+"/DataTableEnums.btk", config.GenPath)
	//	exec.Genrate("erl", config.GenPath+"/DataTables.btk", config.GenPath)
	//}
	//
	//if config.Data.ErlGen.NeedData {
	//	for _, v := range config.Data.DataBase {
	//		exec.Genrate("erl", config.GenPath+"/"+v.Database+".btk", config.GenPath)
	//	}
	//}
	//
	//for _, v := range config.Data.ErlGen.GenList {
	//	exec.Genrate("erl", config.Path+"/"+v, config.GenPath)
	//}
}

func copyGo() {
	if len(config.Data.Go.GenList) <= 0 || !config.Data.Go.IsOpen {
		return
	}
	if len(config.Data.Go.BintalkPath) > 0 {
		b := config.Path + config.Data.Go.BintalkPath
		createPath(b)
		if len(config.Data.ExcelPath) > 0 {
			copyFile(b+"DataTables.go", config.GenPath+"DataTables.go")
			copyFile(b+"DataTableEnums.go", config.GenPath+"DataTableEnums.go")
		}

		if config.Data.Go.NeeData {
			for _, v := range config.Data.DataBase {
				copyFile(b+v.Database+".go", config.GenPath+v.Database+".go")
			}
		}
		for _, v := range config.Data.Go.GenList {
			f := strings.ReplaceAll(v, "btk", "go")
			copyFile(b+f, config.GenPath+f)
		}
	}

	if len(config.Data.Go.BytePath) > 0 {
		b := config.Path + config.Data.Go.BytePath
		createPath(b)
		file := getFiles(config.GenPath + "bytes/")
		for _, v := range file {
			f := strings.Split(v.Name(), "_")
			if len(f) > 0 && (f[0] == "g" || f[0] == "s") {
				copyFile(config.Path+config.Data.Go.BytePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
			}
		}
	}
}

func copyCs() {
	if len(config.Data.Cs.GenList) <= 0 || !config.Data.Cs.IsOpen {
		return
	}

	if len(config.Data.ExcelPath) > 0 {
		b := config.Path + config.Data.Cs.TablePath
		createPath(b)
		copyFile(b+"DataTables.cs", config.GenPath+"DataTables.cs")
		copyFile(b+"DataTableEnums.cs", config.GenPath+"DataTableEnums.cs")

		if len(config.Data.Cs.ResTableName) > 0 {
			b = config.Path + config.Data.Cs.ResTablePath
			createPath(b)
			file := getFiles(config.GenPath + "bytes/")
			for _, v := range file {
				if contiansName(v.Name(), config.Data.Cs.ResTableName) {
					copyFile(config.Path+config.Data.Cs.ResTablePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
				}
			}
		}
		if len(config.Data.Cs.BytePath) > 0 {
			b = config.Path + config.Data.Cs.BytePath
			createPath(b)
			file := getFiles(config.GenPath + "bytes/")
			for _, v := range file {
				f := strings.Split(v.Name(), "_")
				if len(f) > 0 && (f[0] == "g" || f[0] == "c") && !contiansName(v.Name(), config.Data.Cs.ResTableName) {
					copyFile(config.Path+config.Data.Cs.BytePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
				}
			}
		}
	}

	if len(config.Data.Cs.BintalkPath) > 0 {
		b := config.Path + config.Data.Cs.BintalkPath
		createPath(b)
		for _, v := range config.Data.Cs.GenList {
			f := strings.ReplaceAll(v, "btk", "cs")
			copyFile(b+f, config.GenPath+f)
		}
	}
}

func copyDart() {
	if len(config.Data.Dart.GenList) <= 0 || !config.Data.Dart.IsOpen {
		return
	}

	if len(config.Data.ExcelPath) > 0 {
		b := config.Path + config.Data.Dart.TablePath
		createPath(b)
		copyFile(b+"DataTables.dart", config.GenPath+"DataTables.dart")
		copyFile(b+"DataTableEnums.dart", config.GenPath+"DataTableEnums.dart")

		if len(config.Data.Dart.ResTableName) > 0 {
			b = config.Path + config.Data.Dart.ResTablePath
			createPath(b)
			file := getFiles(config.GenPath + "bytes/")
			for _, v := range file {
				if contiansName(v.Name(), config.Data.Dart.ResTableName) {
					copyFile(config.Path+config.Data.Dart.ResTablePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
				}
			}
		}

		if len(config.Data.Dart.BytePath) > 0 {
			b = config.Path + config.Data.Dart.BytePath
			createPath(b)
			file := getFiles(config.GenPath + "bytes/")
			for _, v := range file {
				f := strings.Split(v.Name(), "_")
				if len(f) > 0 && (f[0] == "g" || f[0] == "c") && !contiansName(v.Name(), config.Data.Dart.ResTableName) {
					copyFile(config.Path+config.Data.Dart.BytePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
				}
			}
		}
	}

	if len(config.Data.Dart.BintalkPath) > 0 {
		b := config.Path + config.Data.Dart.BintalkPath
		createPath(b)
		for _, v := range config.Data.Dart.GenList {
			f := strings.ReplaceAll(v, "btk", "dart")
			copyFile(b+f, config.GenPath+f)
		}
	}
}

func copyLua() {
	if len(config.Data.Lua.GenList) <= 0 || !config.Data.Lua.IsOpen {
		return
	}

	if len(config.Data.ExcelPath) > 0 {
		b := config.Path + config.Data.Lua.TablePath
		createPath(b)
		copyFile(b+"DataTables.lua", config.GenPath+"DataTables.lua")
		copyFile(b+"DataTableEnums.lua", config.GenPath+"DataTableEnums.lua")

		if len(config.Data.Lua.ResTableName) > 0 {
			b = config.Path + config.Data.Lua.ResTablePath
			createPath(b)
			file := getFiles(config.GenPath + "bytes/")
			for _, v := range file {
				if contiansName(v.Name(), config.Data.Lua.ResTableName) {
					copyFile(config.Path+config.Data.Lua.ResTablePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
				}
			}
		}
		if len(config.Data.Lua.BytePath) > 0 {
			b = config.Path + config.Data.Lua.BytePath
			createPath(b)
			file := getFiles(config.GenPath + "bytes/")
			for _, v := range file {
				f := strings.Split(v.Name(), "_")
				if len(f) > 0 && (f[0] == "g" || f[0] == "c") && !contiansName(v.Name(), config.Data.Lua.ResTableName) {
					copyFile(config.Path+config.Data.Lua.BytePath+v.Name(), config.GenPath+"/bytes/"+v.Name())
				}
			}
		}
	}

	if len(config.Data.Lua.BintalkPath) > 0 {
		b := config.Path + config.Data.Lua.BintalkPath
		//createPath(b)
		for _, v := range config.Data.Lua.GenList {
			f := strings.ReplaceAll(v, "btk", "lua")
			copyFile(b+f, config.GenPath+f)
		}
	}

}

func copyJava() {
	if len(config.Data.Java.GenList) <= 0 || !config.Data.Java.IsOpen {
		return
	}
	if len(config.Data.Java.BintalkPath) > 0 {
		b := config.Path + config.Data.Java.BintalkPath
		createPath(b)
		if len(config.Data.ExcelPath) > 0 {
			copyFile(b+"DataTables.java", config.GenPath+"DataTables.java")
			copyFile(b+"DataTableEnums.java", config.GenPath+"DataTableEnums.java")
		}
		for _, v := range config.Data.Java.GenList {
			f := strings.ReplaceAll(v, "btk", "java")
			copyFile(b+util.StrFirstToUpper(f), config.GenPath+util.StrFirstToUpper(f))
		}
	}
}

func copyErl() {
	//if len(config.Data.ErlGen.GenList) <= 0 {
	//	return
	//}
	//src := config.Path + config.Data.ErlGen.Src
	//createPath(src)
	//
	//include := config.Path + config.Data.ErlGen.Include
	//createPath(include)
	//
	//files, err := ioutil.ReadDir(config.GenPath)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, v := range files {
	//	if strings.Contains(v.Name(), ".erl") {
	//		copyFile(src+v.Name(), config.GenPath+v.Name())
	//	} else if strings.Contains(v.Name(), ".hrl") {
	//		copyFile(include+v.Name(), config.GenPath+v.Name())
	//	}
	//
	//}
	//
	//if len(config.Data.ExcelPath) > 0 {
	//	priv := config.Path + config.Data.ErlGen.Priv
	//	createPath(priv)
	//	file := getFiles(config.GenPath + "bytes/")
	//	for _, v := range file {
	//		f := strings.Split(v.Name(), "_")
	//		if len(f) > 0 && (f[0] == "g" || f[0] == "s") {
	//			copyFile(priv+v.Name(), config.GenPath+"/bytes/"+v.Name())
	//		}
	//	}
	//}

}

func createPath(path string) {
	_, e := os.Stat(path)
	if e == nil {
		os.RemoveAll(path)
	}
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Println("make path error:  " + err.Error())
		return
	}
}

func copyFile(dstFilePath string, srcFilePath string) {
	sourceFileStat, err := os.Stat(srcFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		fmt.Println(fmt.Errorf("%s is not a regular file", srcFilePath))
		return
	}
	source, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer source.Close()

	destination, err := os.Create(dstFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}

func contiansName(name string, names []string) bool {
	flag := false
	for _, v := range names {
		if strings.Contains(name, v) {
			flag = true
			break
		}
	}
	return flag
}
