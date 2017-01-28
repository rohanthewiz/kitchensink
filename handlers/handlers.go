package handlers

import (
	"github.com/valyala/fasthttp"
	"github.com/rohanthewiz/kitchensink/models/product"
	"github.com/rohanthewiz/kitchensink/templates"
	"encoding/json"
	"log"
	"net/url"
	"fmt"
	"github.com/kataras/iris"
)

func MainPageHandler(context *iris.Context) {
	ctx := context.GetRequestCtx()  // Get the FastHTTP context
	ctx.SetContentType("text/html; charset=utf-8")
	p := &templates.MainPage{
		//CTX: ctx, // or any other param we are setup to pass
	}
	templates.WritePageTemplate(ctx, p)
}

func ProductsHandler(context *iris.Context) {
	ctx := context.GetRequestCtx()
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

func ProductsSeedHandler(context *iris.Context) {
	ctx := context.GetRequestCtx()
	err := product.SeedDB()
	var code int
	if err != nil {
		code = 500
	} else {
		code = 200
	}
	log.Println(code)
	ctx.Redirect("/", code) // does code do anything here?
}

// We'll use Iris static handler from here on
// Poor man's static handler :-)
//func StaticHandler(ctx *fasthttp.RequestCtx) {
//	arr := strings.Split(string(ctx.Path()), "/")
//	if len(arr) > 2 {
//		file_path := "dist/" + strings.Join(arr[2:], "/")
//		log.Println("File path", file_path)
//		file, err := ioutil.ReadFile(file_path)
//		if err != nil {
//			ctx.SetStatusCode(500)
//			log.Println("Error reading file", file_path)
//			return
//		}
//		ctx.Write(file)
//	} else {
//		ctx.SetStatusCode(fasthttp.StatusNotFound)
//	}
//}

func ErrorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)
}
