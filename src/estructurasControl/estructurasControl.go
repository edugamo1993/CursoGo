package main

import (
	"fmt"
)

func main() {
	// Estructura de Control: IF

	i := 6

	if i < 6 {
		fmt.Println("Es menor")
	} else if i > 6 {
		fmt.Println("Es mayor")
	} else {
		fmt.Println("Es igual")
	}

	if i = i - 1; i < 6 {
		fmt.Println("i es menor que 6")
	}

	// Estructura de control: Switch

	var dia int

	fmt.Printf("Teclea el núemro de día de la semana: ")

	fmt.Scanln(&dia)

	switch dia { //Con variable
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Día erroneo")
	}
	numero := 23
	switch { //Sin variable
	case numero > 25:
		fmt.Println("Es mayor que 25")
	case numero < 50:
		fmt.Println("Es menor que 50")
		fallthrough // no se sale del switch si ponemos esto.
	case numero < 25:
		fmt.Println("Es menor que 25")
	default:
		fmt.Println("Por defecto")
	}
}
