package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FuncIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World, ")
	fmt.Println(r.Method)
	// การเขียน query param http://localhost:8000/?name=mind
	fmt.Println(r.URL.Query().Get("name"))
	if r.Method == "GET" {
		io.WriteString(w, "hello world this is get method")
	} else {
		io.WriteString(w, "hello world this is other method")
	}

}

func main() {
	http.HandleFunc("/", FuncIndex)
	fmt.Println("start server completed")
	error := http.ListenAndServe(":8000", nil)
	if error != nil {
		log.Fatal(error)
	}
}
