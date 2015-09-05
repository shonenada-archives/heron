package heron

import (
    "net/http"

    "github.com/zenazn/goji/web"

    "github.com/shonenada/heron/models"
)

func AccountController(c web.C, w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    db := GetDatabase()
    username := r.PostForm.Get("username")
    password := r.PostForm.Get("password")
    user := models.Account{Username: username, Password: password}
    record := db.NewRecord(user)
    if record {
        db.Create(&user)
    }
    RenderJson(w, map[string]interface{}{
        "success": true,
        "info": "inserted",
    })
}

func AccountJoinController(c web.C, w http.ResponseWriter, r *http.Request) {
    RenderTemplate(w, "account_join.html", "base.html", nil)
}
