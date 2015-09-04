package onf

import (
    "fmt"
    "net/http"
    "github.com/zenazn/goji/web"
)

func indexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func helloHandler(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}