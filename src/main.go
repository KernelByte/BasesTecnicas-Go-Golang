package main

// Se importa la libreria automaticamente
import (
	"fmt"
)

// Funcion principal Main
func main() {

	// Declaracion de una variable constante
	const variable float64 = 3.14
	const variable2 = 3.14345
	fmt.Println("El valor de la variables es: ", variable, "Y el valor de la segunda es: ", variable2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Print(base, altura, area, "/")

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// realizar el area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area es: ", areaCuadrado)
}
