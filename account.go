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
		username := c.URLParams["username"]
		user := models.Account{}
		db.Where("username = ?", username).First(&user)
		if user.Username == "" {
			raw_data := map[string]string{"message": "user not found"}
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
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		repeatPassword := r.PostForm.Get("repeat-password")
		if username == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "username cannot be empty",
			})
			return
		}
		if password == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "password cannot be empty",
			})
			return
		}
		if repeatPassword == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "repeat password cannot be empty",
			})
			return
		}
		if password != repeatPassword {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "repeat password is not match with passwrord",
			})
			return
		}
		account := models.Account{}
		db.Where("username = ?", username).First(&account)
		if account.Username != "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "username exists",
			})
			return
		}
		account.Username = username
		account.Password = password
		account.Actived = true
		if db.NewRecord(account) {
			db.Create(&account)
			RenderJson(w, map[string]interface{}{
				"success": true,
			})
			return
		} else {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "record exist",
			})
			return
		}
	}
}

func AccountSignController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	store := sessions.NewCookieStore([]byte(Config.SecretKey))
	if r.Method == "POST" {
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		if username == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "usernmae cannot be empty",
			})
			return
		}
		if password == "" {
			RenderJson(w, map[string]interface{}{
				"succcess": false,
				"message":  "password cannot be empty",
			})
			return
		}
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
