package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
	func main() {
		data, _ := os.ReadFile("./productos.csv")
		for _, dat := range strings.Split(string(data), "\n") {
			cod := dat.Split
			fmt.Println(dat)
		}
	}
*/
func main() {
	data, _ := os.ReadFile("./productos.csv")
	re1 := regexp.MustCompile(`;`)
	fmt.Println("ID\t\t\t\tPRECIO\t\t\t\tCANTIDAD")

	dataTabs := re1.ReplaceAllString(string(data), "\t\t\t\t")
	fmt.Println(dataTabs)
}
