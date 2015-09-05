package heron

import (
    "os"
    "log"
    "net/http"

    "github.com/zenazn/goji"
)

var logger = log.New(os.Stdout, "[heron]: ", log.LstdFlags)

func StartServer () {
    Init()
    SetupRoutes()
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(Config.StaticPath))))
    goji.Serve()
}
