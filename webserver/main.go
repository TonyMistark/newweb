package main

import (
    "fmt"
    "io"
    "net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1>Hello, this is my first page!</h1>")
}

func main() {
	fmt.Println("Start Server...")
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":8989", nil)
}
