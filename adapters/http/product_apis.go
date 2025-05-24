package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products/app"
	"products/domain"
)

type ProductHandler struct {
	service app.ProductService
}

func NewProductHandler(service app.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	users, err := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
