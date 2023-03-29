package readBtk

import (
	"btn_tools/config"
	"fmt"
	"io/ioutil"
	"strings"
)

func Read(name string)  {
	bytesPath := config.Path + "/platform.btk"
	byteData, err := ioutil.ReadFile(bytesPath)
	if err != nil {
		fmt.Println("dont read file: " + name)
		return
	}
	strData := string(byteData)

	have := false
	values := strings.Split(strData, "\r\n")
	for _, v := range values {
		if !have {
			if strings.Contains(v, "struct"){
				have = true
				fmt.Println(v)
			} else if strings.Contains(v, "enum"){
				have = true
				fmt.Println(v)
			}
		} else {
			if strings.Contains(v, "}"){
				have = false
			} else if !strings.Contains(v, "{") {
				fmt.Println(v)
			}
		}
	}
}
