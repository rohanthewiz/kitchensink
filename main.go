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

	log.Println("Starting the server at http://localhost:8080...")
	err := fasthttp.ListenAndServe("localhost:8080", hash_router.Route)
	if err != nil {
		log.Fatal("unexpected error in server:", err)
	}
}
