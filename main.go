package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("call, %s\n", req.RemoteAddr)
		w.Write([]byte("hello"))
	})

	http.ListenAndServe(":8088", nil)
}
