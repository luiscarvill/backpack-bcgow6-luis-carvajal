odepackage main

import "fmt"

/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura.
Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

type persona struct {
	nombre  string
	edad    int
	correo  string
	contras string
}

func (p *persona) ChangeName(name string) {
	p.nombre = name
}

func (p *persona) ChangeAge(edad int) {
	p.edad = edad
}

func (p *persona) ChangeEmail(email string) {
	p.correo = email
}

func (p *persona) ChangePassw(pwrd string) {
	p.contras = pwrd
}

func main() {
	user := persona{
		nombre:  "Usuario",
		edad:    0,
		correo:  "correo@correo.com",
		contras: "unacontraseña123",
	}
	fmt.Println(user)
	fmt.Println("====================")
	user.ChangeName("Luis Carvajal")
	user.ChangeAge(22)
	user.ChangeEmail("luiscarvajal@correo.com")
	user.ChangePassw("laSuPerContraS007=")
	fmt.Println(user)

}
