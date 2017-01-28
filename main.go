package main

import (
	"log"
	"github.com/rohanthewiz/kitchensink/handlers"
	"github.com/rohanthewiz/kitchensink/models/product"
	"github.com/kataras/iris"
)

func main() {
	product.InitDB()
	defer product.CloseDB()

	log.Println("Starting the server at http://0.0.0.0:8094...")
	iris.Get("/", handlers.MainPageHandler)
	iris.Get("/products", handlers.ProductsHandler)
	iris.Get("/seed", handlers.ProductsSeedHandler)
	iris.StaticServe("./dist", "/dist")
	iris.Listen(":8094")
}
