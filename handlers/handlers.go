package handlers

import (
	"github.com/valyala/fasthttp"
	"github.com/rohanthewiz/kitchensink/models/product"
	"github.com/rohanthewiz/kitchensink/templates"
	"encoding/json"
	"strings"
	"io/ioutil"
	"log"
	"net/url"
	"fmt"
)

func MainPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	templates.WriteHome(ctx, "A Presentation of some technologies")
}

func ProductsHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/json; charset=utf-8")
	qargs := ctx.QueryArgs().QueryString()
	log.Printf("QueryArgs - querystring: %v\n", string(qargs))
	values, err := url.ParseQuery(string(qargs))
	if err != nil {
		fmt.Println("Err", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest); return
	}
	pattern, ok := values["q"] // pattern is an array of strings
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusBadRequest); return
	}
	res, err := product.QueryProducts(pattern)
	if err != nil {
		ctx.SetStatusCode(500); return
	}
	str, err := json.Marshal(res)
	log.Printf("products: %v\n", string(str))
	if err != nil {
		ctx.SetStatusCode(500)
	} else {
		ctx.WriteString(string(str))
	}
}

func ProductsSeedHandler(ctx *fasthttp.RequestCtx) {
	err := product.SeedDB()
	var code int
	if err != nil {
		code = 500
	} else {
		code = 200
	}
	log.Println(code)
	ctx.Redirect("/", code) // does code to anything here?
}

// Poor man's static handler :-)
func StaticHandler(ctx *fasthttp.RequestCtx) {
	arr := strings.Split(string(ctx.Path()), "/")
	if len(arr) > 2 {
		file_path := "dist/" + strings.Join(arr[2:], "/")
		log.Println("File path", file_path)
		file, err := ioutil.ReadFile(file_path)
		if err != nil {
			ctx.SetStatusCode(500)
			log.Println("Error reading file", file_path)
			return
		}
		ctx.Write(file)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}

func ErrorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)
}
