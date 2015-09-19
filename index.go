package heron

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func IndexController(c web.C, w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", "base.html", map[string]interface{}{})
}
