package heron

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Route struct {
	url     string
	methods []string
	handler web.HandlerType
}

var routes = []Route{
	{"/", []string{"GET"}, IndexController},
	{"/apis/accounts", []string{"POST"}, AccountController},
	{"/apis/accounts/:name", []string{"GET"}, AccountController},
	{"/apis/accounts/sign", []string{"POST", "DELETE"}, AccountSignController},
	{"/apis/events", []string{"GET", "POST"}, EventController},
	{"/apis/follows", []string{"POST"}, FollowController},
}

func SetupRoutes() {
	for _, route := range routes {
		for _, method := range route.methods {
			switch method {
			case "GET":
				goji.Get(route.url, route.handler)
				break
			case "POST":
				goji.Post(route.url, route.handler)
				break
			case "PUT":
				goji.Put(route.url, route.handler)
				break
			case "PATCH":
				goji.Patch(route.url, route.handler)
				break
			case "DELETE":
				goji.Delete(route.url, route.handler)
				break
			default:
				goji.Handle(route.url, route.handler)
			}
		}
	}
}
