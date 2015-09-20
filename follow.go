package heron

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func FollowController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		return
	}
}
