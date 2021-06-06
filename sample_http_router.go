package main

import (
    "net/http"
    "fmt"
    "log"

    "github.com/julienschmidt/httprouter"
)

func main() {
    router := httprouter.New()

    router.GET("/", rootPageHandler)
    router.GET("/login", login)


    log.Fatal(http.ListenAndServe(":3000", router))
}

func rootPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "welcome to my go api server\n")
    fmt.Println("Root endpoint is hooked!")

    cookie, err := r.Cookie("hoge")

    if err != nil {
        fmt.Fprintf(w, "cookie hoge is empty\n")
        fmt.Println("cookie hoge is empty")
        return
    }

    v := cookie.Value
    fmt.Println(v)
    fmt.Fprintf(w, "read cookie: %s\n", v)
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    cookie := &http.Cookie{
        Name: "hoge",
        Value: "bar",
    }
    http.SetCookie(w, cookie)

    fmt.Fprintf(w, "set cookie\n")
    fmt.Println("set cookie!")
}
