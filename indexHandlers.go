package heron

import (
    "net/http"
    "github.com/zenazn/goji/web"
)

func indexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
    who := c.URLParams["who"]
    if who == "" {
        who = "你"
    }
    renderTemplate(w, "index.html", "base.html", map[string]interface{}{
        "Who": who,
    })
}
