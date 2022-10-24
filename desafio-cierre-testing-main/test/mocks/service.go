package mocks

import (
	"fmt"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
)

type MockService struct {
	DataMock []products.Product
	Error    string
}

func (m *MockService) GetAllBySeller(sellerID string) ([]products.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil

}
