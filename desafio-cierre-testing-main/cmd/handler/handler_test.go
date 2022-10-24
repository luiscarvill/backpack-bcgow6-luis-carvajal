package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/bootcamp-go/desafio-cierre-testing/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockService mocks.MockService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := NewHandler(&mockService)
	r := gin.Default()
	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetAllStatusOk(t *testing.T) {
	// arrange
	mockService := mocks.MockService{
		DataMock: []products.Product{},
		Error:    "",
	}
	var resp []products.Product
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=jaimito", "")

	r.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, mockService.DataMock, resp)
	assert.Nil(t, err)
}

func TestGetAllErrorParamRequired(t *testing.T) {
	// arrange
	mockService := mocks.MockService{
		DataMock: []products.Product{},
		Error:    "error ocurred",
	}
	var resp map[string]string
	expected := map[string]string{
		"error": "seller_id query param is required",
	}
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 400, rr.Code)
	assert.Equal(t, expected, resp)
}

func TestGetAllErrorNotFound(t *testing.T) {
	// arrange
	mockService := mocks.MockService{
		DataMock: []products.Product{},
		Error:    "No se encontro el producto especificado",
	}
	var resp map[string]string
	expected := map[string]string{
		"error": "No se encontro el producto especificado",
	}
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=20", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 404, rr.Code)
	assert.Equal(t, expected, resp)
}
