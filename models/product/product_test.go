package product

import (
	"testing"
	"sync"
)

var mutex = &sync.Mutex{}


func TestAll(t *testing.T) {
	mutex.Lock()
	product_db = NewDB("products_test.db")
	mutex.Unlock()
	defer CloseDB()
	if product_db == nil {
		t.Error()
	}

	product_db.CreateTable()

	products := []Product{
		{1, "John Brown", "john_brown_1"},
		{2, "Mary Sue", "mary_sue_2"},
		{3, "Samantha", "samantha_3"},
	}
	product_db.StoreProducts(products)
	t.Log(product_db.QueryProducts(""))

	more_products := []Product{
		{4, "Big Joe", "big_joe_4"},
		{5, "Banana Head", "banana_head_5"},
		{6, "Fruit n Spice", "fruit_n_spice_6"},
	}
	product_db.StoreProducts(more_products)
	t.Log(product_db.QueryProducts(""))
}
