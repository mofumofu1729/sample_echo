package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    startWebServer()
}

func startWebServer() error {
    r := mux.NewRouter().StrictSlash(true)

    r.HandleFunc("/", rootPage)

    return http.ListenAndServe(fmt.Sprintf(":%d", 3000), r)
}


func rootPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to my go api server!\n")
    fmt.Println("Root endpoint is hooked!")
}
