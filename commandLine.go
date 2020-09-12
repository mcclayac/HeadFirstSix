package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	fmt.Println("-------------------------------------------------------")
	fmt.Println("OS Arguements")
	fmt.Println("OS Argumenets :", os.Args[1:])
	fmt.Println("Getegid returns the numeric effective group id :", os.Getegid())

	// os - Env
	fmt.Println("-------------------------------------------------------")
	fmt.Println("env")

	os.Setenv("AppNAME", "gopher")
	os.Setenv("AppDir", "/usr/gopher")

	fmt.Printf("%s lives in %s.\n", os.Getenv("AppNAME"), os.Getenv("AppDir"))

	// os - Env
	fmt.Println("-------------------------------------------------------")
	fmt.Println("user-id")

	fmt.Printf("user id:%d , group id:%d\n\n", os.Geteuid(), os.Getgid())
	fmt.Printf("process ID:%d\n\n", os.Getpid())

	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)

}
