package heron

import (
	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
	"net/http"
)

func FollowController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()

	if r.Method == "POST" {
		return
	}
}
