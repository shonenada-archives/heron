package heron

import (
	"net/http"

	arrow "github.com/bmuller/arrow/lib"
	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
)

func IndexController(c web.C, w http.ResponseWriter, r *http.Request) {
	db := GetDatabase()
	events := []models.Event{}
	today := arrow.Now().CFormat("%Y-%m-%d")
	tomorrow := arrow.Tomorrow().CFormat("%Y-%m-%d")
	db.Where("created_at BETWEEN ? AND ?", today, tomorrow).Find(&events)
	RenderTemplate(w, "index.html", "base.html", map[string]interface{}{
		"events": events,
		"db":     db,
	})
}
