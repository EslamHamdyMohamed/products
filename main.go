package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"products/adapters/http" // aliased import simplified
	"products/adapters/storage"
	"products/app"
	"products/domain"
)

func main() {
	fmt.Println("Start Main")
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	_ = db.AutoMigrate(&domain.Product{})

	log.Println("Database connected successfully")
	repo := storage.NewSqliteProductRepository(db)
	useCase := app.NewProductService(repo)
	handler := http.NewProductHandler(useCase)

	r := gin.Default()
	g := r.Group("/products")
	g.POST("", handler.CreateProduct)
	g.DELETE("/:id", handler.DeleteProduct)
	g.GET("/:id", handler.GetProduct)
	g.GET("/", handler.ListProducts)

	log.Println("Server started at :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
