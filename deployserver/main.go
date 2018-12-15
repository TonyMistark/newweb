package main

import (
    "fmt"
    "io"
    "net/http"
    "os/exec"
    "log"
)

func reLaunch() {
    cmd := exec.Command("sh", "./deploy.sh")
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
    err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1>Hello, this is my deploy page!</h1>")
    reLaunch()
}

func main() {
    fmt.Println("Start deploy server 0.0.0.0:8988...")
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":8988", nil)
}
