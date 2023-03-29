package exec

import (
	"btn_tools/config"
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func Genrate(target string, resPath string, tarPath string) {
	cmd := exec.Command(config.Path+"/bin/"+getBintalk(), "-g", target, "-o", tarPath, resPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

}

func Genrate2(target, resPath, tarPath, option, packageName string) {
	//arg := []string{}
	//cmd := exec.Command(cnf.Path + "/bintalk.exe",  "-g", "go", "-o", cnf.Path + "/src/bintalk/", cnf.Path +"/DataTables.btk")
	//cmd := exec.Command(config.Path+"/bin/"+getBintalk(), "-g", target, "-o", tarPath, option, packageName, resPath)
	cmd := exec.Command(config.Path+"/bin/"+getBintalk(), "-g", target, "-o", tarPath, option, packageName, resPath)
	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	//err := cmd.Start()
	//_, err := cmd.Output()
	//if err != nil {
	//	fmt.Println("Error:", err.Error())
	//	return
	//}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}

func getBintalk() string {
	osType := runtime.GOOS
	if osType == "windows" {
		return "bintalk.exe"
	} else if osType == "darwin" {
		return "bintalk_mac"
	} else {
		return "bintalk"
	}
}
