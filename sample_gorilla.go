package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

var (
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
)

func main() {
    startWebServer()
}

func startWebServer() error {
    r := mux.NewRouter().StrictSlash(true)

    r.HandleFunc("/", rootPage)
    r.HandleFunc("/login", login)
    r.HandleFunc("/secret", secret)

    return http.ListenAndServe(fmt.Sprintf(":%d", 3000), r)
}


func rootPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to my go api server!\n")

    fmt.Println("Root endpoint is hooked!")
}

func login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    fmt.Fprintf(w, "session authenticated!\n")
    fmt.Println("session authenticated!")

    session.Values["authenticated"] = true
    session.Save(r, w)
}

func secret(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    fmt.Fprintln(w, "the cake is a lie!")
}
