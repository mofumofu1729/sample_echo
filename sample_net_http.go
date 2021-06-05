package main

import (
    "net/http"
    "fmt"
    "log"
)

func main() {
    http.HandleFunc("/", rootPageHandler)
    http.HandleFunc("/login", login)


    http.ListenAndServe(":3000", nil)
}

func rootPageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to my go api server\n")
    fmt.Println("Root endpoint is hooked!")

    cookie, err := r.Cookie("hoge")

    if err != nil {
        log.Fatal("Cookie: ", err)
    }

    v := cookie.Value
    fmt.Println(v)
    fmt.Fprintf(w, "read cookie\n")
}

func login(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name: "hoge",
        Value: "bar",
    }
    http.SetCookie(w, cookie)

    fmt.Fprintf(w, "set cookie\n")
    fmt.Println("set cookie!")
}
