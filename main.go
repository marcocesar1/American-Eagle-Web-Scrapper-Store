package main

import (
	"fmt"
	"time"
	"webscrapping/config"
	"webscrapping/models"
	"webscrapping/repositories"
	webscraper "webscrapping/web-scraper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	productCategories := []models.CategoryWeb{
		{
			Name:       "Shrots Hombre",
			Url:        "https://www.ae.com/mx/es/c/men/bottoms/shorts/cat5180435?pagetype=plp",
			CategoryId: 7,
		},
		{
			Name:       "Tops Camisas Hombre",
			Url:        "https://www.ae.com/mx/es/c/men/tops/cat10025?pagetype=plp",
			CategoryId: 1,
		},
		{
			Name:       "Polos Hombre",
			Url:        "https://www.ae.com/mx/es/c/men/tops/polos/cat510004?pagetype=plp",
			CategoryId: 8,
		},
		{
			Name:       "Jeans Hombre",
			Url:        "https://www.ae.com/mx/es/c/men/bottoms/jeans/cat6430041?pagetype=plp",
			CategoryId: 3,
		},
		{
			Name:       "Sudaderas Hombre",
			Url:        "https://www.ae.com/mx/es/c/men/tops/hoodies-sweatshirts/cat90020?pagetype=plp",
			CategoryId: 9,
		},
		{
			Name:       "Accesorios",
			Url:        "https://www.ae.com/mx/es/c/men/accessories-socks/cat4840022?pagetype=plp",
			CategoryId: 10,
		},
		{
			Name:       "Tops Mujer",
			Url:        "https://www.ae.com/mx/es/c/women/tops/cat10049?pagetype=plp",
			CategoryId: 6,
		},
		{
			Name:       "Jeans Mujer",
			Url:        "https://www.ae.com/mx/es/c/women/bottoms/jeans/cat6430042?pagetype=plp",
			CategoryId: 4,
		},
		{
			Name:       "Vestidos Mujer",
			Url:        "https://www.ae.com/mx/es/c/women/dresses/cat1320034?pagetype=plp",
			CategoryId: 5,
		},
	}

	db := config.ConnectDb()
	defer db.Close()

	scrapper := webscraper.NewAEScrapper()
	productRepo := repositories.NewProductRepo(db)
	imageRepo := repositories.NewImageRepo(db)

	for _, item := range productCategories {
		products := scrapper.GetProductsByCategory(item.Url)

		for _, product := range products {

			productId, err := productRepo.InsertOne(product, item.CategoryId)
			if err != nil {
				fmt.Println("Error inserting product: " + err.Error())

				continue
			}

			time.Sleep(500 * time.Millisecond)

			fmt.Printf("ProductId: %d\n", productId)

			err = imageRepo.InsertMany(product.Images, productId)
			if err != nil {
				fmt.Println("Error inserting product images: " + err.Error())
			}

			time.Sleep(500 * time.Millisecond)

		}
	}

}
