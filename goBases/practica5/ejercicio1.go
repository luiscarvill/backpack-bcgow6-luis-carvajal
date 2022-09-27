package main

import (
	"os"
)

type Producto interface {
	ConcatProduct() string
}
type producto struct {
	id       string
	price    string
	quantity string
}
type Empresa interface {
	CrearArchivo() string
}
type empresa struct {
	products []Producto
}

func (t *empresa) Agregar(p Producto) {
	t.products = append(t.products, p)
}
func (p producto) ConcatProduct() string {
	//fmt.Print("aqyu")
	return p.id + ";" + p.price + ";" + p.quantity + "\n"
}

func (e empresa) CrearArchivo() string {
	var total string
	//var valid int = 0
	for _, product := range e.products {
		total += product.ConcatProduct()
	}

	os.WriteFile("./productos.csv", []byte(total), 0644)
	//fmt.Println("aqui total----", total)
	//dataArchivo := ConcatProduct()
	return "ok"
}
func main() {
	prod := producto{
		id:       "1",
		price:    "2_000",
		quantity: "1",
	}
	inv := empresa{
		products: []Producto{prod},
	}
	inv.Agregar(producto{
		id:       "2",
		price:    "30_000_000",
		quantity: "3",
	})
	inv.CrearArchivo()
}
