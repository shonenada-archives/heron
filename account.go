package heron

import (
	"encoding/json"
	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
	"net/http"
)

func AccountController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	if r.Method == "GET" {
		name := c.URLParams["name"]
		user := models.Account{}
		db.Where("username = ?", name).First(&user)
		if user.Username == "" {
			raw_data := map[string]interface{}{"err": "user not found"}
			data := json.Marshal(raw_data)
			return 404, data
		} else {
			return RenderJson(w, user)
		}
	}

	if r.Method == "POST" {
	}
}

func AccountSignController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	if r.Method == "POST" {
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		user = models.Account{}
		db.Where("username = ? AND password = ?", username, password).First(&user)
		if user.Username == username {
			// session signin
			return RenderJson(w, map[string]interface{}{
				"success": true,
			})
		} else {
			return RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "username or password is incorrect",
			})
		}
	}
	if r.Method == "DELETE" {
		// session signout
	}
}
