package handler

import (
	"net/http"
	"strings"

	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/domains"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/products"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre" binding:"required"`
	Type  string  `json:"tipo" binding:"required"`
	Count int     `json:"cantidad" binding:"required"`
	Price float64 `json:"precio" binding:"required"`
}

type requestName struct {
	Name string `json:"nombre" binding:"required"`
}
type requestID struct {
	ID int `json:"id" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (s *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product := domains.Product(req)
		id, err := s.service.Store(product)
		if err != nil {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		product.ID = id
		c.JSON(http.StatusCreated, web.NewResponse(product, "", http.StatusCreated))
	}
}

func (s *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestName
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product, err := s.service.GetByName(req.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(product, "", http.StatusOK))
	}
}
func (s *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		products, err := s.service.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(products, "", http.StatusOK))
	}
}
func (s *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestID
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}
		err := s.service.Delete(req.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse("ok", "", http.StatusOK))
	}
}
