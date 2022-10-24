package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/pkg/store"
	"github.com/stretchr/testify/assert"
)

/* Test with Mock */
func TestGetAll(t *testing.T) {
	id1, id2 := "1", "2"

	products := []*Product{{ID: "1", SellerID: "20", Description: "cola&pola", Price: 1300.0},
		{ID: "2", SellerID: "20", Description: "ponny", Price: 1300.0},
		{ID: "3", SellerID: "1", Description: "malta", Price: 1300.0},
	}
	data, _ := json.Marshal(products)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &mock,
	}
	r := NewRepository(&stubStore)

	prodGet, err := r.GetAllBySeller(id1)
	fmt.Println("====== test read invoque ========")
	assert.True(t, mock.ReadInvoked)
	fmt.Println("====== test error nil ========")
	assert.Nil(t, err)
	fmt.Println("====== test return by seller equal ========")
	assert.Equal(t, *&products[2].Description, prodGet[0].Description)
	fmt.Println("====== test not equal results ========")
	assert.NotEqual(t, products, prodGet)
	prodGet2, err := r.GetAllBySeller(id2)
	fmt.Println("====== test response without results ========")

	assert.Equal(t, []Product{}, prodGet2)
	fmt.Println("====== test2 error response ========")

	assert.Equal(t, err.Error(), "No se encontro el producto especificado", err)
}

/* Test with Stub */

func TestGetAllSt(t *testing.T) {
	products := []Product{
		{
			ID:          "first",
			SellerID:    "jaimito",
			Description: "pastel de pollo",
			Price:       2300.0,
		},
		{
			ID:          "second",
			SellerID:    "jaimito",
			Description: "chococono",
			Price:       2300.0,
		}, {
			ID:          "third",
			SellerID:    "josue",
			Description: "chocorramo",
			Price:       2300.0,
		},
	}
	data, _ := json.Marshal(products)
	dbMock := store.Mock{
		Data:  data,
		Error: nil,
	}
	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	r := NewRepository(&stubStore)
	prodGetAll, err := r.GetAllBySeller("josue")
	fmt.Println("====== test2 error equal ========")
	assert.Equal(t, products[2], prodGetAll[0])
	assert.Nil(t, err)
	prodGetAll2, err2 := r.GetAllBySeller("jaimito")
	products = products[:len(products)-1]
	fmt.Println("====== test2 products search ========")
	assert.Equal(t, products, prodGetAll2)
	assert.Nil(t, err2)

}

func TestGetAllError(t *testing.T) {
	errorExpected := errors.New("debe ingresar sellerId para la busqueda")
	dbMock := store.Mock{ // No le pasamos información
		Error: errorExpected, // Deberia fallar
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	products, err := r.GetAllBySeller("")
	fmt.Println("====== test3 without id ========")
	assert.NotNil(t, err)
	assert.Equal(t, errorExpected, err)
	fmt.Println("====== test3 error equal ========")
	assert.Equal(t, []Product{}, products)
}
