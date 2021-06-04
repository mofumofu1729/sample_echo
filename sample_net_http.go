package main

import (
    "net/http"
    "fmt"
)

func main() {
    http.HandleFunc("/", rootPageHandler)


    http.ListenAndServe(":3000", nil)
}

func rootPageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to my go api server\n")

    fmt.Println("Root endpoint is hooked!")
}

func login(w http.ResponseWriter, r *http.Request) {

