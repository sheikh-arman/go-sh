package main

import (
	"fmt"
	"gomodules.xyz/go-sh"
)

func main() {
	fmt.Println("tut")
	sh.NewSession().SetDir("/").Command("pwd")
	sh.Test("dir", "mydir")
	v := map[string]int{}
	err := sh.Command("echo", `{"number": 1}`).UnmarshalJSON(&v)
	fmt.Println(err)
	//buildImage()
	/*sh.Command("echo", "hello\tworld").Run()
	session := sh.NewSession()
	session.SetEnv("BUILD_ID", "123")
	session.SetDir("/")
	//# then call cmd
	session.Command("echo", "hello").Run()
	//# set ShowCMD to true for easily debug
	session.ShowCMD = true*/
}

func buildImage() {
	tag := "arman/pg:ok"
	dockerFilePath := "/home/arman/go/src/github.com/appscode-images/docker-tag-history/docker/"
	args := []interface{}{
		"build",
		"-t",
		tag,
		dockerFilePath,
	}
	// It has another type of usage :
	//v := make(map[string]interface{})
	//sh.Command().Command().UnmarshalJson(&v)
	//for key, val := range v {
	//    fmt.Printf("%+v %+v \n", key, val)
	//}
	data, err := sh.Command("docker", args...).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
