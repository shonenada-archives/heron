package heron

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func AccountSignInViewController(c web.C, w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "account/signin.html", "base.html", nil)
	return
}
