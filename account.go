package heron

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
)

func AccountController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	if r.Method == "GET" {
		name := c.URLParams["name"]
		user := models.Account{}
		db.Where("username = ?", name).First(&user)
		if user.Username == "" {
			raw_data := map[string]string{"err": "user not found"}
			data, err := json.Marshal(raw_data)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			http.Error(w, string(data), 404)
			return
		} else {
			RenderJson(w, user)
			return
		}
	}

	if r.Method == "POST" {
	}
}

func AccountSignController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	store := sessions.NewCookieStore([]byte(Config.SecretKey))
	if r.Method == "POST" {
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		user := models.Account{}
		db.Where("username = ? AND password = ?", username, password).First(&user)
		if user.Username == username {
			session, err := store.Get(r, "user")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			session.Values["username"] = user.Username
			session.Save(r, w)
			RenderJson(w, map[string]interface{}{
				"success": true,
			})
			return
		} else {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "username or password is incorrect",
			})
			return
		}
	}
	if r.Method == "DELETE" {
		session, err := store.Get(r, "user")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		session.Values["username"] = ""
		session.Save(r, w)
	}
}
