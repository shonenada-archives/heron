package heron

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
)

func EventsController(c web.C, w http.ResponseWriter, r *http.Request) {
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
		name := r.PostForm.Get("name")
		content := r.PostForm.Get("content")
		dueString := r.PostForm.Get("due_date")
		due, err := time.Parse(time.RFC3339, dueString)
		if err != nil {
			RenderJson(w, map[string]interface{}{
				"success": true,
				"message": "fail to parse due date",
			})
			return
		}
		event := models.Event{Name: name, Content: content, Due: due, UserId: currentUser.ID}
		if db.NewRecord(event) {
			db.Create(&event)
			RenderJson(w, map[string]interface{}{
				"success": true,
			})
			return
		} else {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "record exists",
			})
			return
		}
	}
}

func EventController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	currentUser, err := CurrentUser(r)
	if err != nil {
		RenderJson(w, map[string]interface{}{
			"success": false,
			"message": "user not login",
		})
		return
	}

	eventId := c.URLParams["eid"]
	event := models.Event{}
	db.Where("id = ?", eventId).First(&event)
	if event.ID <= 0 {
		raw_data := map[string]string{"message": "event not found"}
		data, err := json.Marshal(raw_data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Error(w, string(data), 404)
		return
	}

	if r.Method == "GET" {
		resp, err := json.Marshal(event)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		RenderJson(w, resp)
		return
	}

	if r.Method == "DELETE" {
		if event.UserId != currentUser.ID {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "permission denied",
			})
			return
		}
		db.Delete(&event)
		RenderJson(w, map[string]interface{}{
			"success": true,
		})
		return
	}
}
