package product

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"errors"
	"strings"
	"fmt"
)

var product_db *DB // DB singleton
const query_limit = 8 // For limit, but also cache activation // Must match limit set in Selector

// PRODUCT DB

type DB struct {
	path string
	db *sql.DB
}

func NewDB(path string) *DB {
	d := new(DB)
	if path == "" { path = "products.db" }
	d.path = path
	db, err := sql.Open("sqlite3", d.path)
	if err != nil { panic(err) }
	if db == nil { panic("db is nil") }
	d.db = db
	return d
}

func (d DB) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

func (d DB) CreateTable() {
	query :=
`		CREATE TABLE IF NOT EXISTS products(
			Id TEXT NOT NULL PRIMARY KEY,
			Label TEXT,
			Value TEXT,
			CreatedAt DATETIME);`
	_, err := d.db.Exec(query)
	if err != nil {
		log.Println("Error creating table.")
		panic(err)
	}
}

func (d DB) DropTable() {
	query :=
`		DROP TABLE IF EXISTS products;`
	_, err := d.db.Exec(query)
	if err != nil {
		log.Println("Error dropping table.")
		panic(err)
	}
}

func (d DB) StoreProducts(products []Product) error {
	if d.db == nil {
		return errors.New("In StoreProducts: db is nil")
	}
	query :=
		`	INSERT OR REPLACE INTO products(
				Id, Label, Value, CreatedAt)
				values(?, ?, ?, CURRENT_TIMESTAMP)`
	stmt, err := d.db.Prepare(query)
	if err != nil { return err }
	defer stmt.Close()

	for _, product := range products {
		_, err2 := stmt.Exec(product.Id, product.Label, product.Value)
		if err2 != nil { panic(err2) }
	}
	return nil
}

func (d DB) QueryProducts(tokens []string) ([]Product, error) {
	products := []Product{}
	if d.db == nil {
		return products, errors.New("Database is not initialized")
	}
	query := "SELECT Id, Label, Value FROM products"
	if len(tokens) > 0 {
		where_clause := " WHERE "
		clauses := []string{}
		for _, token := range tokens {
			clauses = append(clauses, "Label LIKE '%" + token + "%'")
		}
		where_clause += strings.Join(clauses, " AND ")
		query += where_clause
	}

	query += " ORDER BY Label DESC"
	query += fmt.Sprintf(" LIMIT %d", query_limit)
	log.Println(query)

	rows, err := d.db.Query(query)
	if err != nil {
		log.Println("Err in db query:", err)
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err2 := rows.Scan(&product.Id, &product.Label, &product.Value)
		if err2 != nil {
			return products, err2 // could probably add our say here too
		}
		products = append(products, product)
	}
	return products, nil
}
