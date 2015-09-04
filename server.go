package onf

import (
    "fmt"
    "github.com/zenazn/goji"
)

func StartServer () {
    SetupRoutes()
    goji.Serve()
}
