package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "strconv"
	"strings"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	thezeHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Theze Nutz, MF!\n")
		io.WriteString(w, "-------------------------------------------------------\n")
		io.WriteString(w, "OS Arguements\n")
		aSlice := os.Args[1:]
		aSliceString := strings.Join(aSlice, ",")
		s := "OS Argumenets :" + aSliceString
		io.WriteString(w, s)
		s = fmt.Sprintf("Getegid group id :%d\n", os.Getegid())
		io.WriteString(w, s)

		// os - Env
		io.WriteString(w, "-------------------------------------------------------\n")
		io.WriteString(w, "env\n")

		os.Setenv("AppNAME", "gopher")
		os.Setenv("AppDir", "/usr/gopher")

		outputString := fmt.Sprintf("%s lives in %s.\n\n", os.Getenv("AppNAME"), os.Getenv("AppDir"))
		io.WriteString(w, outputString)

		// os - Env
		io.WriteString(w, "-------------------------------------------------------\n")
		io.WriteString(w, "user-id\n")

		outputString = fmt.Sprintf("user id:%d , group id:%d\n\n", os.Geteuid(), os.Getgid())
		io.WriteString(w, outputString)

		outputString = fmt.Sprintf("process ID:%d\n\n", os.Getpid())
		io.WriteString(w, outputString)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/these", thezeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
