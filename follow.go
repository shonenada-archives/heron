package heron

import (
	"net/http"

	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
)

func FollowController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	currentUser, err := CurrentUser(r)
	if err != nil || currentUser.Username == "" {
		RenderJson(w, map[string]interface{}{
			"success": false,
			"message": "user not login",
		})
		return
	}
	if r.Method == "POST" {
		follow_id := r.PostForm.Get("follow_id")
		account := models.Account{}
		db.Where("id = ?", follow_id).First(&account)
		if account.Username == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "user not found",
			})
			return
		}
		return
	}
}
