package heron

import (
    "github.com/zenazn/goji"
)

func StartServer () {
    SetupRoutes()
    goji.Serve()
}
