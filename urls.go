package heron

import (
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

type Route struct {
    url string
    method string
    handler web.HandlerType
}

var routes = []Route {
    {"/", "get", indexHandler},
    {"/hello/:name", "get", helloHandler},
}

func SetupRoutes () {
    for _, route := range routes {
        switch route.method {
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
