package products

import (
	"errors"

	"github.com/bootcamp-go/desafio-cierre-testing/pkg/store"
)

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	if len(sellerID) == 0 {
		return []Product{}, errors.New("debe ingresar sellerId para la busqueda")
	}
	var prodList []Product
	var retProd []Product
	if err := r.db.Read(&prodList); err != nil {
		return nil, err
	}

	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	for _, prod := range prodList {
		if prod.SellerID == sellerID {
			retProd = append(retProd, prod)
		}
	}
	if len(retProd) == 0 {
		return []Product{}, errors.New("No se encontro el producto especificado")
	}

	return retProd, nil
}
