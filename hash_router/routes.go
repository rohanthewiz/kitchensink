package hash_router

import (
	"github.com/valyala/fasthttp"
	"github.com/rohanthewiz/kitchensink/handlers"
	"strings"
	"log"
)

var routes = map[string]func(ctx *fasthttp.RequestCtx) {
	"/": handlers.MainPageHandler,
	"/products": handlers.ProductsHandler,
	"/products-seed": handlers.ProductsSeedHandler,
}

func Route(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	log.Println("URL path:", path)
	arr := strings.Split(path, "/")
	if len(arr) > 2 && (arr[1] == "assets" || arr[1] == "dist") { // Static
		handlers.StaticHandler(ctx)
	} else {
		handler, ok := routes[path] // pull the route out of a hash :-)
		if ok {
			handler(ctx)
		} else {
			handlers.ErrorPageHandler(ctx)
		}
	}
}
