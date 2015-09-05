package heron

import (
    "net/http"
    "github.com/zenazn/goji/web"
)

func IndexController(c web.C, w http.ResponseWriter, r *http.Request) {
    who := c.URLParams["who"]
    if who == "" {
        who = "你"
    }
    RenderTemplate(w, "index.html", "base.html", map[string]interface{}{
        "Who": who,
    })
}
