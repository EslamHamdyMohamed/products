package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	httphandler "products/adapters/http" // Alias to avoid naming conflict
	"products/adapters/storage"
	"products/app"
	"products/domain"
)

func setupRouter() *gin.Engine {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	//	db, err := gorm.Open(sqlite.Open("memory.db"), &gorm.Config{})
	if err != nil {

		panic("failed to connect to test db")
	}
	_ = db.AutoMigrate(&domain.Product{})

	repo := storage.NewSqliteProductRepository(db)
	usecase := app.NewProductService(repo)
	handler := httphandler.NewProductHandler(usecase)

	r := gin.Default()
	g := r.Group("/products")
	g.POST("", handler.CreateProduct)
	g.GET("/:id", handler.GetProduct)
	g.GET("/", handler.ListProducts)
	g.DELETE("/:id", handler.DeleteProduct)

	return r
}

func TestCreateProduct(t *testing.T) {
	fmt.Println("Run TestCreateProduct")
	r := setupRouter()

	product := map[string]interface{}{
		"id":    "p1",
		"name":  "Test Product",
		"price": 100,
	}
	body, _ := json.Marshal(product)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 but got %d", w.Code)
	}
	fmt.Println("Run TestCreateProduct Passed")

}

func TestGetProduct(t *testing.T) {
	fmt.Println("Run TestGetProduct")

	r := setupRouter()

	product := map[string]interface{}{
		"id":    "p2",
		"name":  "Another Product",
		"price": 200,
	}
	body, _ := json.Marshal(product)

	createW := httptest.NewRecorder()
	createReq, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))
	createReq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(createW, createReq)

	getW := httptest.NewRecorder()
	getReq, _ := http.NewRequest("GET", "/products/p2", nil)
	r.ServeHTTP(getW, getReq)

	if getW.Code != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", getW.Code)
	}
	fmt.Println("Run TestGetProduct Passed")
}

func TestDeleteProduct(t *testing.T) {
	fmt.Println("Run TestDeleteProduct")

	r := setupRouter()

	product := map[string]interface{}{
		"id":    "p3",
		"name":  "ToDelete",
		"price": 300,
	}
	body, _ := json.Marshal(product)

	createW := httptest.NewRecorder()
	createReq, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))
	createReq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(createW, createReq)

	delW := httptest.NewRecorder()
	delReq, _ := http.NewRequest("DELETE", "/products/p3", nil)
	r.ServeHTTP(delW, delReq)

	if delW.Code != http.StatusNoContent {
		t.Fatalf("Expected status 204 but got %d", delW.Code)
	}
	fmt.Println("Run TestDeleteProduct Passed")

}

func TestListProducts(t *testing.T) {
	fmt.Println("Run TestListProducts")

	r := setupRouter()

	product := map[string]interface{}{
		"id":    "p4",
		"name":  "Listed",
		"price": 400,
	}
	body, _ := json.Marshal(product)

	createW := httptest.NewRecorder()
	createReq, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))
	createReq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(createW, createReq)

	listW := httptest.NewRecorder()
	listReq, _ := http.NewRequest("GET", "/products/", nil)
	r.ServeHTTP(listW, listReq)

	if listW.Code != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", listW.Code)
	}

	var products []domain.Product
	if err := json.Unmarshal(listW.Body.Bytes(), &products); err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}
	if len(products) == 0 {
		t.Fatal("Expected at least one product")
	}
	fmt.Println("Run TestListProducts Passed")

}
