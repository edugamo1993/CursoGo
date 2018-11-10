package main

import (
	"fmt"
)

func main() {
	fmt.Println("Empieza el programa.")
	//El tamaño del array no puede variar
	//Declarar Arrays
	var miVector [3]int
	fmt.Println(miVector) //Por defecto les pone valor 0

	//De dos dimensiones
	var k [3][2]float64
	fmt.Println(k)

	miVector[1] = 25

	fmt.Println(miVector) //el valor del indice 1 será 25

	fmt.Println(miVector[1])

	y := [5]int{1, 2, 3, 4, 5}

	fmt.Println(y)

	j := [...]int{1, 2, 3, 4} //analiza el tamaño de los elementos que les pasamos.

	fmt.Println(j)

	fmt.Println(len(y)) //Tamaño del array y
	fmt.Println(y[len(y)-2])

	//Comparación

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}

	fmt.Println(a == b)
	fmt.Println(a == c)
}
