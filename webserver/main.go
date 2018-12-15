package main

import (
    "fmt"
    "io"
    "net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1>Hello World, i am God!</h1>")
}

func main() {
    fmt.Println("Start Server 0.0.0.0:8989...")
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":8989", nil)
}
