package webscraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"webscrapping/constans"
	"webscrapping/helpers"
	"webscrapping/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"

	_ "github.com/go-sql-driver/mysql"
)

type AEScrapper struct {
}

func NewAEScrapper() *AEScrapper {
	return &AEScrapper{}
}

func (scrapper *AEScrapper) GetProductsByCategory(categoryUrl string) []models.Product {
	products := []models.Product{}

	res, err := http.Get(categoryUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".results-list .product-tile div.tile-details").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3.product-name").Text()
		productLink, exist := s.Find("a").Attr("href")

		fmt.Printf("---> Producto %d: %s", i, title)
		if exist {
			fullURL := fmt.Sprintf("https://www.ae.com%s", productLink)

			product, _ := scrapper.GetProductInfo(fullURL)

			products = append(products, product)
		}

	})

	return products
}

func (scrapper *AEScrapper) GetProductInfo(productUrl string) (models.Product, error) {
	res, err := http.Get(productUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	productSection := doc.Find(".qa-product-details-page")
	if productSection.Length() == 0 {
		return models.Product{}, fmt.Errorf("product details section not found")
	}

	titleTextElement := productSection.Find("div.product-name-and-flags h1").Text()
	priceTextElement := productSection.Find("div.extras-content div.product-list-price").Text()

	title := strings.TrimSpace(strings.Replace(titleTextElement, "AE ", "", -1))
	slug := slug.Make(title) + "-" + helpers.GenerateRandomString(5)
	price := strings.TrimSpace(strings.Replace(priceTextElement, "$", "", -1))
	images := getProductImages(productSection)

	emptyArr := []interface{}{}
	productOptions, _ := json.Marshal(emptyArr)

	product := models.Product{
		Title:          title,
		Slug:           slug,
		Price:          helpers.ParsePrice(price),
		Images:         images,
		Description:    constans.Lorem,
		ProductOptions: productOptions,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return product, nil
}

func getProductImages(s *goquery.Selection) []string {
	var images []string

	s.Find("div.pdp-main-carousel.carousel div.item-image").Each(func(i int, s *goquery.Selection) {
		image, existImage := s.Find("picture img").Attr("src")

		if existImage {
			urlImage := strings.Replace(image, "//s7d2.scene7.com", "https://s7d2.scene7.com", 1)
			images = append(images, urlImage)
		}
	})

	return images
}
