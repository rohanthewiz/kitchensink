package handlers

import (
	"github.com/valyala/fasthttp"
	"github.com/rohanthewiz/kitchensink/models"
	"github.com/rohanthewiz/kitchensink/templates"
	"encoding/json"
	"strings"
	"io/ioutil"
	"log"
)

func MainPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	templates.WriteHome(ctx, "A Presentation of some technologies")
}

func ProductsHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/json; charset=utf-8")
	str, err := json.Marshal(models.Products)
	if err != nil {
		ctx.SetStatusCode(500)
	} else {
		ctx.WriteString(string(str))
	}
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
