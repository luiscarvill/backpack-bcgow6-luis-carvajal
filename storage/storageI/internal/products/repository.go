package products

import (
	"database/sql"
	"fmt"

	"github.com/backpack-bcgow6-luis-carvajal/storage/storageI/internal/domains"
)

type Repository interface {
	Store(p domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
	GetAll() ([]domains.Product, error)
	Delete(id int) error
	GetById(id int) bool
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

// Ejercicio 1 - Implementar GetByName()
// Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:
// - Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorne un objeto del tipo Product.
// - Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.

func (r *repository) GetByName(name string) (domains.Product, error) {
	getQuery := "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	stmt, err := r.db.Prepare(getQuery)
	if err != nil {
		return domains.Product{}, fmt.Errorf("Error en la consulta:  %v", err)
	}
	defer stmt.Close()
	var product domains.Product
	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domains.Product{}, fmt.Errorf("Sin registros %s --> %v", name, err)
	}

	return product, nil
}

// Ejercicio 2 - Replicar Store()
// Tomar el ejemplo visto en la clase y diseñar el método Store():
// - Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
// - Implementar el método Store.

func (r *repository) Store(p domains.Product) (int, error) {
	storeQuery := "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(storeQuery)
	if err != nil {
		return 0, fmt.Errorf("Error en la consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return 0, fmt.Errorf("Error en la consulta: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener ultimo ID: %v", err)
	}

	return int(id), nil
}

// PRACTICA 2 ==> Ejercicio 2 - Implementar GetAll()
// Diseñar un método GetAll.
// Dentro del archivo repository desarrollar el método GetAll().
// Comprobar el correcto funcionamiento.

func (r *repository) GetAll() ([]domains.Product, error) {
	var products []domains.Product
	getQuery := "SELECT id, name, type, count, price FROM products"
	rows, err := r.db.Query(getQuery)
	if err != nil {
		return []domains.Product{}, fmt.Errorf("Error en la consulta:  %v", err)
	}
	//se recorren todas las filas
	for rows.Next() {
		//por fila se obtiene un objeto del tipo Product
		var product domains.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return []domains.Product{}, fmt.Errorf("Error en  registros %s ", err)
		}
		// se añade el objeto al slice
		products = append(products, product)
	}
	return products, nil
}

// PRACTICA 2 ==> Ejercicio 3 - Implementar Delete()
// Diseñar un método para eliminar un recurso de la base de datos.
// Dentro del archivo repository desarrollar el método Delete().
// Comprobar el correcto funcionamiento.
func (r *repository) Delete(id int) error {
	deleteQuery := "DELETE FROM products WHERE id = ?"
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return fmt.Errorf("Error en delete: %v", err)
	}
	//cierra la ejecucion para evitar consumo de memoria
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("Error en delete: %v", err)
	}
	return nil

}

func (r *repository) GetById(id int) bool {
	getQuery := "SELECT id, name, type, count, price FROM products WHERE id = ?"
	stmt, err := r.db.Prepare(getQuery)
	if err != nil {
		return false
	}
	defer stmt.Close()
	var product domains.Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return false
	}
	return true
}
