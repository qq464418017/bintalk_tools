package enum

import (
	bFile "btn_tools/bFile"
	"btn_tools/config"
	"btn_tools/util"
)

var luaFile *bFile.BTFile
var goFile *bFile.BTFile
var csFile *bFile.BTFile
var javaFile *bFile.BTFile
var erlFile *bFile.BTFile
var dartFile *bFile.BTFile

func init() {
	if config.Data.Lua.IsOpen {
		luaFile = bFile.NewBtkFile(config.GenPath + "/DataTableEnums.lua")
	}
	if config.Data.Cs.IsOpen {
		csFile = bFile.NewBtkFile(config.GenPath + "/DataTableEnums.cs")
	}
	if config.Data.Go.IsOpen {
		goFile = bFile.NewBtkFile(config.GenPath + "/DataTableEnums.go")
		goFile.Write("package bintalk")
		goFile.WrteNewLine()
	}
	if config.Data.Java.IsOpen {
		javaFile = bFile.NewBtkFile(config.GenPath + "/DataTableEnums.java")
		javaFile.Write("public class DataTableEnums")
		javaFile.Write("{")
	}
	if config.Data.Dart.IsOpen {
		dartFile = bFile.NewBtkFile(config.GenPath + "/DataTableEnums.dart")
		dartFile.Write("import 'dart:typed_data';")
		dartFile.Write("import 'package:" + config.Data.Dart.PackageName + "/bintalk/i_bintalk.dart';")
		dartFile.Write("import 'package:" + config.Data.Dart.PackageName + "/bintalk/protocol_reader.dart';")
		dartFile.Write("import 'package:" + config.Data.Dart.PackageName + "/bintalk/protocol_writer.dart';")
		dartFile.WrteNewLine()
	}
}

func WriteEnumName(enumName string) {
	if luaFile != nil {
		luaFile.Write("E_" + enumName + " =")
		luaFile.Write("{")
	}
	if csFile != nil {
		csFile.Write("public enum E_" + enumName)
		csFile.Write("{")
	}
	if goFile != nil {
		config.Data.Go.EnumName = enumName
		goFile.Write("const (")
	}
	if javaFile != nil {
		javaFile.WriteTable("public enum E_" + enumName + "{")
		config.Data.Java.Enums = make(map[string]string)
	}
	if dartFile != nil {
		dartFile.Write("enum E_" + enumName + "{")
		//config.Data.Dart.Enums = make(map[string]string)
	}
}

func WriteEnumItem(itemName, value string, isLast bool) {
	//index,err := strconv.Atoi(row)
	//value := row
	//if err == nil {
	//	index -= 1
	//	value = strconv.Itoa(index)
	//}

	if luaFile != nil {
		luaFile.WriteTable(itemName + " = " + value + ",")
	}
	if csFile != nil {
		csFile.WriteTable(itemName + " = " + value + ",")
	}
	if goFile != nil {
		goFile.WriteTable("E_" + config.Data.Go.EnumName + itemName + " = " + value)
	}
	if javaFile != nil {
		if isLast {
			javaFile.WriteTwoTable(util.StrToUpper(itemName) + ";")
			javaFile.WrteNewLine()
		} else {
			javaFile.WriteTwoTable(util.StrToUpper(itemName) + ",")
		}
		config.Data.Java.Enums[itemName] = value
	}
	if dartFile != nil {
		dartFile.WriteTable(itemName + ",")
	}
}

func WriteEnd() {
	if luaFile != nil {
		luaFile.Write("}")
		luaFile.WrteNewLine()
	}
	if csFile != nil {
		csFile.Write("}")
		csFile.WrteNewLine()
	}
	if goFile != nil {
		goFile.Write(")")
		goFile.WrteNewLine()
	}
	if javaFile != nil {
		for i, v := range config.Data.Java.Enums {
			javaFile.WriteTwoTable("public static final int " +
				util.StrToUpper(i) + "_VALUE = " + v + ";")
		}
		javaFile.WriteTable("}")
		javaFile.WrteNewLine()
	}
	if dartFile != nil {
		dartFile.Write("}")
		dartFile.WrteNewLine()
	}
}

func Close() {
	if luaFile != nil {
		luaFile.Close()
	}
	if goFile != nil {
		goFile.Close()
	}
	if csFile != nil {
		csFile.Close()
	}
	if javaFile != nil {
		javaFile.Write("}")
		javaFile.Close()
	}
	if erlFile != nil {
		erlFile.Close()
	}
	if dartFile != nil {
		dartFile.Close()
	}
}
