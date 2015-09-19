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
		}
	}

	if r.Method == "POST" {
	}
}

func AccountSignController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
}

func aaaaa() {
	db := GetDatabase()
	name := c.URLParams["name"]
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	checkUser := models.Account{}
	db.Where("username = ?", username).First(&checkUser)
	if checkUser.Username != "" {
		RenderJson(w, map[string]interface{}{
			"success": false,
			"info":    "username exists",
		})
		return
	}
	user := models.Account{Username: username, Password: password}
	record := db.NewRecord(user)
	if record {
		db.Create(&user)
	}
	RenderJson(w, map[string]interface{}{
		"success": true,
		"info":    "inserted",
	})
}
