package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/cmd/server/handler"
	cnn "github.com/backpack-bcgow6-luis-carvajal/storage/storageI/db"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/domains"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var s = createServer()

func createServer() *gin.Engine {
	os.Setenv("USERNAME", "root")
	os.Setenv("PASSWORD", "")
	os.Setenv("DATABASE", "storage")

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)

	p := handler.NewProduct(serv)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetByName())
	pr.POST("/", p.Store())
	pr.GET("/all/", p.GetAll())
	pr.DELETE("/", p.Delete())

	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

var ID = 0

// Ejercicio 3 - Ejecutar Store()
// Diseñar un Test que permita ejecutar Store y comprobar la correcta inserción del registro en la tabla.

func TestStoreProduct_Ok(t *testing.T) {
	new := domains.Product{
		Name:  "cafe",
		Type:  "cafe polvo",
		Count: 10,
		Price: 8400.4,
	}

	product, err := json.Marshal(new)
	assert.Nil(t, err)

	req, rr := createRequest(http.MethodPost, "/api/v1/products/", string(product))
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// struct for assertion
	p := struct {
		Data domains.Product
	}{}
	err = json.Unmarshal(rr.Body.Bytes(), &p)
	assert.Nil(t, err)

	new.ID = p.Data.ID
	ID = new.ID
	assert.Equal(t, new, p.Data)
}

// Ejercicio 4 - Ejecutar GetByName()
// Diseñar un Test que permita ejecutar GetByName y comprobar que retorne el registro buscado en caso de que exista.

func TestGetByNameProduct_Ok(t *testing.T) {
	req, rr := createRequest(http.MethodGet, "/api/v1/products/", `{"nombre":"cafe"}`)
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Extra  - Diseñar test para las funcionalidades()
// 	1- Test GetAll()
// 	2- Test Delete()

func TestGetAll_Ok(t *testing.T) {
	req, rr := createRequest(http.MethodGet, "/api/v1/products/all/", `{}`)
	s.ServeHTTP(rr, req)

	// assert code
	fmt.Println(rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestDelete_Ok(t *testing.T) {
	s2 := strconv.Itoa(ID)
	data := `{"id":` + s2 + `}`
	req, rr := createRequest(http.MethodDelete, "/api/v1/products/", data)
	s.ServeHTTP(rr, req)

	// assert code
	fmt.Println(rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDelete_Fail(t *testing.T) {
	req, rr := createRequest(http.MethodDelete, "/api/v1/products/", `{"id":7}`)
	s.ServeHTTP(rr, req)

	// assert code
	fmt.Println(rr.Body.String())
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
