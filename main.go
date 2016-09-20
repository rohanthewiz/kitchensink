package main

import (
	"log"
	"github.com/valyala/fasthttp"
	"github.com/rohanthewiz/kitchensink/templates"
	"encoding/json"
)
type Product struct {
	Label string
	Value string
}

var products = []Product{
	{Label: "John Brown", Value: "john_brown"},
	{Label: "Mary Jane", Value: "mary_jane"},
	{Label: "JoAnn Smith", Value: "joann"},
	{Label: "Mike Jackson", Value: "mickey"},
	{Label: "Little John", Value: "little_john"},
}

func main() {
	log.Println("Starting the server at http://localhost:8080...")
	err := fasthttp.ListenAndServe("localhost:8080", requestHandler)
	if err != nil {
		log.Fatal("unexpected error in server: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		mainPageHandler(ctx)
	case "/products":
		productsHandler(ctx)
	default:
		errorPageHandler(ctx)
	}
}

func mainPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	templates.WriteHome(ctx, "A Presentation of some technologies")
}

func productsHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/json; charset=utf-8")
	str, err := json.Marshal(products)
	if err != nil {
		ctx.SetStatusCode(500)
	} else {
		ctx.WriteString(string(str))
	}
}

func errorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}
