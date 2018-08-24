package main

import "fmt"

func main(){
	var numero int
	numero = 1
	fmt.Printf("%d\n",numero) // Printf como en C
	numero = 35 // Utiliza el mismo espacio en memoria asignado para la variable.
	fmt.Printf("%d\n",numero) 
	
	nombre := "Eduardo" //Declara e inicializa la variable a la vez. Automaticamente nombre es Stirng
	fmt.Println(nombre)
	nombre = "Otro valor"
	fmt.Println(nombre)

	nombre, numero = "Lucía", 25

	fmt.Println(nombre,numero)

	numero2 := 40
	numero, numero2 = numero2, numero //Intercambio de variables sin necesidad de variable auxiliar
	fmt.Println(numero,numero2)
	
	nombre2, nombre := "Nuevo Nombre", "Edugamo" 
	fmt.Println(nombre2,nombre)

	multiplicador := 25
	fmt.Println(numero * multiplicador)
}