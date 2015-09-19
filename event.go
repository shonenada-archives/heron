package heron

import (
	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
	"net/http"
)

func EventController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
}
