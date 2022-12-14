package router

import (
	"github.com/bootcamp-go/desafio-cierre-testing/cmd/handler"
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/bootcamp-go/desafio-cierre-testing/pkg/store"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg)
	}

}

func buildProductsRoutes(r *gin.RouterGroup) {
	db := store.NewFileStore("./products.json")

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := handler.NewHandler(service)

	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}

}
