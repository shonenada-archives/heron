package heron

import (
    "net/http"
    "github.com/zenazn/goji/web"
)

func AccountJoinController(c web.C, w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        RenderTemplate(w, "account_join.html", "base.html", map[string]interface{}{
        })
        break
    case "POST":
        break
    }
}
