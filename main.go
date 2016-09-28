package main

import (
	"log"
	"github.com/valyala/fasthttp"
	"github.com/rohanthewiz/kitchensink/hash_router"
	"github.com/rohanthewiz/kitchensink/models/product"
)

func main() {
	product.InitDB()
	defer product.CloseDB()

	log.Println("Starting the server at http://0.0.0.0:8094...")
	err := fasthttp.ListenAndServe("0.0.0.0:8094", hash_router.Route)
	if err != nil {
		log.Fatal("unexpected error in server:", err)
	}
}
