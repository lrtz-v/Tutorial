package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	// http://127.0.0.1:8080/debug/pprof/
	// go tool pprof http://localhost:8080/debug/pprof/profile
	// go tool pprof http://localhost:8080/debug/pprof/heap
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer handle request
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
