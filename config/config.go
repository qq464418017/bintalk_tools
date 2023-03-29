package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Path string
var GenPath string
var Data configJson

type configJson struct {
	GeneradPath string	//生成代码路径
	ExcelPath   string	//配置表路径
	DataBase    []dataBase
	Lua			luaGen
	Cs			csGen
	Go			goGen
	Java		javaGen
	Dart		dartGen
}


type baseGen struct {
	IsOpen		bool		//是否开启
	BintalkPath string		//btk文件路径
	BytePath   string		//配置表bytes路径
	TablePath   string		//配置表结构和枚举路径
	GenList     []string	//要生成的btk文件名
	DataBasePath	string	//数据库文件路径
}

type luaGen struct {
	baseGen
	ResTableName []string
	ResTablePath string
}

type csGen struct {
	baseGen
	ResTableName []string	//单独存放的配置表名
	ResTablePath string		//单独存放的配置表结构和枚举路径
}

type goGen struct {
	baseGen
	ResTableName []string
	ResTablePath string
	PackageName string		//包名
	EnumName	string		//枚举名（go枚举特殊性）
	NeeData		bool		//是否需要数据库
}

type dartGen struct {
	baseGen
	ResTableName []string
	ResTablePath string
	PackageName string		//包名
}

type javaGen struct {
	baseGen
	ResTableName []string
	ResTablePath string
	PackageName string		//包名
	Enums  map[string]string		//所有枚举列表
}

type dataBase struct {
	Host     string
	User     string
	Password string
	Database string
}


func init()  {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Path = p
	fmt.Println(p)
	bytes, err := ioutil.ReadFile(p + "/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	Data = configJson{}
	err = json.Unmarshal(bytes, &Data)
	if err != nil {
		fmt.Println(err)
		return
	}

	path := p + "/generadir/"
	err = createPath(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	GenPath = path
}



func createPath(p string) error {
	_, err := os.Stat(p)
	if err == nil {
		os.RemoveAll(p)
	}
	return os.Mkdir(p, os.ModePerm)
}