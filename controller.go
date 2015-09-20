package heron

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/shonenada/heron/models"
)

func RenderJson(w http.ResponseWriter, data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(j))
}

func CurrentUser(r *http.Request) (models.Account, error) {
	store := sessions.NewCookieStore([]byte(Config.SecretKey))
	session, err := store.Get(r, "user")
	if err != nil {
		return models.Account{}, err
	}
	username := session.Values["username"]
	db := GetDatabase()
	user := models.Account{}
	db.Where("username = ?", username).First(&user)
	return user, nil
}
