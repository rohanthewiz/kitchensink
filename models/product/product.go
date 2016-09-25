package product

import (
	"sync"
	"log"
)

type Product struct {
	Id uint64 `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

// Yep still visible within this package
var	products_seed = []Product{
	{1, "John Brown", "john_brown_1"},
	{2, "Mary Sue", "mary_sue_2"},
	{3, "Samantha", "samantha_3"},
	{4, "Big Joe", "big_joe_4"},
	{5, "Banana Head", "banana_head_5"},
	{6, "Fruit n Spice", "fruit_n_spice_6"},
}

var mutex = &sync.Mutex{}

func InitDB() {
	mutex.Lock()
	product_db = NewDB("products.db")
	mutex.Unlock()
}

func CloseDB() {
	product_db.Close()
}

func QueryProducts(pattern []string) ([]Product, error) {
	if product_db == nil { InitDB()	}
	// unescape URL
	// clean up, escape, build regex
	log.Println("Query string:", pattern)
	return product_db.QueryProducts(pattern)
}

func SeedDB() error {
	if product_db == nil { InitDB()	}
	mutex.Lock()
	product_db.DropTable()
	product_db.CreateTable()
	err := product_db.StoreProducts(products_seed)
	mutex.Unlock()
	return err
}
