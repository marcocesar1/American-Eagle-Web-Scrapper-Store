package repositories

import (
	"database/sql"
	"fmt"
	"time"
)

type ImageRepository struct {
	db *sql.DB
}

func NewImageRepo(db *sql.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func (repo *ImageRepository) InsertMany(images []string, productId int64) error {

	for _, image := range images {
		query := "INSERT INTO Image (url, productId, createdAt, updatedAt) VALUES (?, ?, ?, ?)"
		_, err := repo.db.Exec(query, image, productId, time.Now(), time.Now())
		if err != nil {
			return err
		}

		fmt.Println("Imagen insertada correctamente")
	}
	return nil
}
