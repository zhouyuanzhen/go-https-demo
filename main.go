package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "A tiny https server Demo!\n")
	})
	if e := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}
