package heron

import (
    "os"
    "log"
    "github.com/zenazn/goji"
)

var logger = log.New(os.Stdout, "[heron]: ", log.LstdFlags)

func StartServer () {
    Init()
    SetupRoutes()
    goji.Serve()
}
