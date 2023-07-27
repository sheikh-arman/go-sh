package main

import (
	"fmt"
	"gomodules.xyz/go-sh"
)

func main() {
	/*	output, err := sh.Command("echo", "Hello, World!").Output()
		if err != nil {
			fmt.Println("Error running command:", err)
			return
		}

		fmt.Println("Command output:", string(output))*/
	//installTrivy()
	/*fmt.Println("tut")
	sh.NewSession().SetDir("/").Command("pwd")
	sh.Test("dir", "mydir")
	v := map[string]int{}
	err := sh.Command("echo", `{"number": 1}`).UnmarshalJSON(&v)
	fmt.Println(err)*/
	//buildImage()
	/*sh.Command("echo", "hello\tworld").Run()
	session := sh.NewSession()
	session.SetEnv("BUILD_ID", "123")
	session.SetDir("/")
	//# then call cmd
	session.Command("echo", "hello").Run()
	//# set ShowCMD to true for easily debug
	session.ShowCMD = true*/
	//sh.Command("cat", "mariadb:11.1.1-rc-jammy.json").Run()
	args := []interface{}{
		"ayugsda",
		">>",
		"arman.txt",
	}
	sh.Command("echo", args...).Run()

	/*textToAppend := "arman\n"

	// Execute the shell command using go-sh
	cmd := sh.Command("echo", textToAppend)

	// Open the file with the "os" package and set the file to append mode (os.O_APPEND)
	file, err := os.OpenFile("arman.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	cmd.Stdout = file
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	fmt.Println("Text appended to arman.txt successfully.")*/
}

func installTrivy() {
	args1 := []interface{}{
		"apt-get", "install", "-y", "wget", "apt-transport-https", "gnupg", "lsb-release",
	}
	data, err := sh.Command("sudo", args1...).Output()
	if err != nil {
		fmt.Println(" arm1 ", err)
		return
	}
	fmt.Println(string(data))

	args2 := []interface{}{
		"-qO", "-", "https://aquasecurity.github.io/trivy-repo/deb/public.key",
	}

	args3 := []interface{}{
		"apt-key", "add", "-",
	}
	data, err = sh.Command("wget", args2...).Command("sudo", args3).Output()
	if err != nil {
		fmt.Println(" arm2 ", err)
		return
	}
	fmt.Println(string(data))
	args4 := []interface{}{
		"deb", "https://aquasecurity.github.io/trivy-repo/deb", "$(lsb_release", "-sc)", "main",
	}

	args5 := []interface{}{
		"tee", "-a", "/etc/apt/sources.list.d/trivy.list",
	}
	data, err = sh.Command("echo", args4...).Command("sudo", args5).Output()
	if err != nil {
		fmt.Println(" arm3 ", err)
		return
	}
	fmt.Println(string(data))
	args6 := []interface{}{
		"apt-get", "update",
	}
	data, err = sh.Command("sudo", args6...).Output()
	if err != nil {
		fmt.Println(" arm4 ", err)
		return
	}
	fmt.Println(string(data))
	args7 := []interface{}{
		"apt-get", "install", "-y", "trivy",
	}
	data, err = sh.Command("sudo", args7...).Output()
	if err != nil {
		fmt.Println(" arm5 ", err)
		return
	}
	fmt.Println(string(data))
	fmt.Println("Trivy Installation completed successfully.")
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
