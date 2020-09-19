package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/"))

	http.Handle("/", fileServer)
	http.HandleFunc("/cgi-bin/printenv", printEnv)

	err := http.ListenAndServe("127.0.0.1:1234", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
