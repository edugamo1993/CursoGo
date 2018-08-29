package main

import (
	"fmt"
)

func main() {
	//Contador de numeros impares
	encabezado := `
	*******************************
	Contador de Numeros impares
	*******************************
	`
	//Imprimimos encabezado
	fmt.Println(encabezado)

	fmt.Println("Teclea el primer número")
	var numero1 int
	var numero2 int
	//Leemos el número
	fmt.Scanln(&numero1)
	fmt.Println("Teclea el segundo número")
	fmt.Scanln(&numero2)

	var contador int

	for i := numero1; i <= numero2; i++ {
		if i%2 != 0 {
			contador++
			fmt.Printf("%d es un numero impar \n", i)
		}
	}
	encabezado = `
	****************************
	Muestra de resultados
	****************************
	`
	fmt.Println(encabezado)
	fmt.Printf("Hay %d números impares entre %d y %d", contador, numero1, numero2)

}
