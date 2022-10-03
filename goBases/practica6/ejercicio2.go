package main

/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.

	Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.

Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/
import "fmt"

type usuario struct {
	nombre   string
	apellido string
	correo   string
	producto []producto
}

type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func (p *producto) CreateProduct(nombre string, precio float64) {
	p.nombre = nombre
	p.precio = precio
	p.cantidad = 1
	//return p
}
func ChargeProduct(u *usuario, producto *producto, cantidad int) {
	u.producto += append(u.producto, []producto)
	u.cantidad = cantidad
}
func main() {
	prod := producto{"Zucaritas", 13_000, 3}
	user := &usuario{
		"luis",
		"carvajal",
		"correo@prueba.com",
		[]producto{prod},
	}
	fmt.Println(user)
	user.ChargeProduct(user)
}
