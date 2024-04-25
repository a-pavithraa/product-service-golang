package domain

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/a-pavithraa/product-service-golang/models"
	"github.com/a-pavithraa/product-service-golang/util"
)

var config = util.LoadAppConfig()

func ConnectToDatabase() (*sql.DB, error) {
	//return sql.Open("postgres", "host=localhost port=15432 user=postgres password=postgres dbname=postgres sslmode=disable")
	return sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.Username, config.DBConfig.Password, config.DBConfig.Database))

}

func CheckProductNameExists(name string) error {

	db, err := ConnectToDatabase()

	if err != nil {
		return err
	}
	defer db.Close()

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM products WHERE name = $1", name).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}
	//If records exists with the same name, return an error
	if count > 0 {
		return NewProductNameFoundError("Product name already exists")
	}
	return nil
}

func GetProductByID(id string) (*models.Product, error) {
	db, err := ConnectToDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM products WHERE id = $1", id)
	var product models.Product
	err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, NewProductNotFoundError("Product not found")
	}
	return &product, nil
}
func GetProducts() ([]models.Product, error) {
	db, err := ConnectToDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func PostProduct(product models.Product) error {
	db, err := ConnectToDatabase()

	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO products (name, description, price) VALUES ($1, $2, $3)", product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}
