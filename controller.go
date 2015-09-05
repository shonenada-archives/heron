package heron

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func RenderJson(w http.ResponseWriter, data interface{}) {
    j, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(j))
}
