package products

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/db"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/domains"
	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// Ejercicio 1 - Testear Store() y GetOne()
// Suponiendo la correcta implementación de los métodos Store y GetOne.
// Escribir un test utilizando go-txdb en donde se guarde un User y
// luego se obtenga, en el caso de que no exista GetOne debe devolver un resultado nulo.
// No hay límite de cantidad de tests ni casos, crear los que considere necesarios para asegurar el correcto funcionamiento de las distintas variaciones.

func TestSaveAndGetByNameTXDB(t *testing.T) {
	db, err := utils.InitTxDB()
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		Name:  "chocorramo",
		Type:  "pastel",
		Count: 12,
		Price: 2200,
	}

	repo := NewRepository(db)
	id, err := repo.Store(*p)

	assert.NoError(t, err)
	p.ID = id
	product, err := repo.GetByName("chocorramo")
	assert.NoError(t, err)
	assert.Equal(t, p, &product)

}

// Ejercicio 2 - Testear Update() y Delete()
// Generar tests para delete para verificar que un registro fue borrado
// correctamente y ya no se puede obtener ni utilizando GetOne ni al llamar a GetAll.
func TestSaveAndDeleteTXDB(t *testing.T) {
	db, err := utils.InitTxDB()
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		Name:  "arepa",
		Type:  "harina",
		Count: 120,
		Price: 1200,
	}

	repo := NewRepository(db)
	id, err := repo.Store(*p)

	assert.NoError(t, err)
	assert.NotZero(t, id)
	err = repo.Delete(id)
	assert.NoError(t, err)

}

// Ejercicio 4 - Testear que sucede en el caso de que falle una query
// Seleccione alguno de los métodos
// (Store, GetOne, GetAll, Update o Delete) y simule un error al ejecutar la consulta.
// El test debe pasar si ocurre un error. (assert.Error).
// El objetivo de este ejercicio es conocer la forma de ver como reacciona nuestra implementación frente a una falla.

func TestGetByNameTXDB_NotFound(t *testing.T) {
	db, err := utils.InitTxDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	product, err := repo.GetByName("chocorramito")
	assert.Error(t, err)
	assert.Zero(t, product)

}

func TestDeleteTXDB_NotFound(t *testing.T) {
	db, err := utils.InitTxDB()
	assert.NoError(t, err)
	defer db.Close()
	repo := NewRepository(db)
	id := -1

	err = repo.Delete(id)
	assert.Error(t, err)

}

// Ejercicio 3 - Replicar tests anteriores utilizando mocks
// Tomar alguno de los tests realizados en los ejercicios anteriores
// (o todos) y replicarlos utilizando go-sqlmock.

func TestGetByNameMock(t *testing.T) {
	db, mock, err := utils.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		ID:    1,
		Name:  "test mock",
		Type:  "type test",
		Count: 1,
		Price: 1000,
	}

	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT_QUERY)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Store failed
	id, err := repo.Store(*p)
	assert.Equal(t, p.ID, id)
	assert.NoError(t, err)

	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(p.ID, p.Name, p.Type, p.Count, p.Price)
	mock.ExpectPrepare(regexp.QuoteMeta(GET_BY_NAME_QUERY)).
		ExpectQuery().
		WithArgs(p.Name).
		WillReturnRows(rows)

	product, err := repo.GetByName(p.Name)
	assert.NoError(t, err)
	assert.Equal(t, p, &product)
}

func TestGetByNameMock_NotFound(t *testing.T) {
	db, mock, err := utils.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{}

	repo := NewRepository(db)
	mock.ExpectPrepare(regexp.QuoteMeta(GET_BY_NAME_QUERY)).
		ExpectQuery().
		WithArgs(p.Name).
		WillReturnError(
			errors.New("Received unexpected error"))

	product, err := repo.GetByName("chocorramo")
	assert.Error(t, err)
	assert.Equal(t, p, &product)
}
func TestGetAllMock(t *testing.T) {
	db := db.MySQLConnection()
	repo := NewRepository(db)
	products, err := repo.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(products) > 0)

}

func TestSaveMock(t *testing.T) {
	db, mock, err := utils.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		ID:    1,
		Name:  "test mock",
		Type:  "type test",
		Count: 1,
		Price: 1000,
	}

	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT_QUERY)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Store failed
	id, err := repo.Store(*p)
	assert.Equal(t, p.ID, id)
	assert.NoError(t, err)
}

func TestSaveMock_Error(t *testing.T) {
	db, mock, err := utils.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{}

	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT_QUERY)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price).
		WillReturnError(errors.New("you have not provided the necessary fields to insert"))

	// Store failed
	id, err := repo.Store(*p)
	assert.Equal(t, 0, id)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestDeleteMock(t *testing.T) {
	db, mock, err := utils.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		ID:    1,
		Name:  "test mock",
		Type:  "type test",
		Count: 1,
		Price: 1000,
	}

	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT_QUERY)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Store failed
	id, err := repo.Store(*p)
	assert.Equal(t, p.ID, id)
	assert.NoError(t, err)

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_QUERY)).
		ExpectExec().
		WithArgs(p.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(p.ID)
	assert.NoError(t, err)
}

func TestDeleteMock_NotFound(t *testing.T) {
	db, mock, err := utils.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_QUERY)).
		ExpectExec().
		WithArgs(15).
		WillReturnError(errors.New("you have not provided the necessary fields to insert"))

	err = repo.Delete(15)
	assert.Error(t, err)
}

// TODO create repo to update data.
// Generar tests para update, en donde se verifique que luego de modificarse un modelo,
// al obtenerlo el mismo posea los cambios realizados.
