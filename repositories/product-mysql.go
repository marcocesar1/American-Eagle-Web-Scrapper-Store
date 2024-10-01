package repositories

import (
	"database/sql"
	"webscrapping/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) InsertOne(product models.Product, categoryId int) (int64, error) {

	query := "INSERT INTO Product (name, description, price, compareAtPrice, isFeatured, isArchived, categoryId,slug,inventory, productOptions, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, product.Title, product.Description, product.Price, 0, 0, 0, categoryId, product.Slug, 30, product.ProductOptions, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return 0, err
	}

	productId, _ := result.LastInsertId()

	return productId, nil
}
