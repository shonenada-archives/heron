package heron

import (
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

type Route struct {
    url string
    methods []string
    handler web.HandlerType
}

var routes = []Route {
    {"/", []string{"get"}, IndexController},
    {"/p/:who", []string{"get"}, IndexController},
    {"/account/join", []string{"get", "post"}, AccountJoinController},
}

func SetupRoutes() {
    for _, route := range routes {
        for _, method := range route.methods {
            switch method {
            case "get":
                goji.Get(route.url, route.handler)
                break
            case "post":
                goji.Post(route.url, route.handler)
                break
            case "put":
                goji.Put(route.url, route.handler)
                break
            case "patch":
                goji.Patch(route.url, route.handler)
                break
            case "delete":
                goji.Delete(route.url, route.handler)
                break
            default:
                goji.Handle(route.url, route.handler)
            }
        }
    }
}
