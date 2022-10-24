package mocks

import (
	"fmt"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
)

type MockStorage struct {
	DataMock []products.Product
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]products.Product)
	*a = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]products.Product)
	m.DataMock = append(m.DataMock, a[len(a)-1])
	return nil
}
