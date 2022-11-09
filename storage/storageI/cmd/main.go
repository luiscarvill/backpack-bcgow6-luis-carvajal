package main

import (
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/cmd/server/handler"
	cnn "github.com/backpack-bcgow6-luis-carvajal/storage/storageI/db"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetByName())
	pr.GET("/all/", p.GetAll())
	pr.DELETE("/", p.Delete())

	r.Run()
}
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
