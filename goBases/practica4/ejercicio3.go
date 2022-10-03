package main

import "fmt"

/*
Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/
const (
	small     string  = "small"
	medium    string  = "medium2"
	big       string  = "big"
	mantThree float64 = 0.03
	mantSix   float64 = 0.06
)

type Producto interface {
	CostCalculate() float64
}

type producto struct {
	prodType string
	price    float64
	name     string
}

type Ecomerce interface {
	Total() float64
	Agregar(Producto)
}

type tienda struct {
	products []Producto
}

func (t tienda) Total() float64 {
	var total float64

	for _, product := range t.products {
		total += product.CostCalculate()
	}

	return total
}

func (t *tienda) Agregar(p Producto) {
	t.products = append(t.products, p)
}

func (p producto) CostCalculate() float64 {
	switch p.prodType {
	case small:
		return p.price
	case medium:
		return p.price + (p.price * mantThree)
	case big:
		return p.price + (p.price * mantSix)
	default:
		return 0
	}
}

func main() {
	prod := producto{
		prodType: big,
		name:     "chocorramo",
		price:    2_000,
	}

	tid := tienda{
		products: []Producto{prod},
	}
	tid.Agregar(producto{
		prodType: small,
		name:     "ponymalta",
		price:    3_100,
	})

	fmt.Printf("%+v\n", tid)
	fmt.Print("+===================")
	fmt.Printf("%+v\n", tid.Total())
}
