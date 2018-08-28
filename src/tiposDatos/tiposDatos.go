package main

import "fmt"
import "unsafe"
import "strconv"

func main(){
	fmt.Println("Empieza el programa.")
	/*
	4 TIPOS DE DATOS
	 - Basicos: Números, Cadenas, Booleanos
	 - Conjuntos: Arrays y Struct
	 - Referencias: Punteros, Segmentos, Maps, Funciones, Canales
	 - Interfaces
	*/

	/*Tipos de datos enteros*/
	// Enteros con signo
	// var entero8 int8
	// var entero16 int16
	var entero32 int32
	var entero64 int64

	// // Enteros sin signo
	// var uentero8 uint8
	// var uentero16 uint8
	// var uentero32 uint32
	// var uentero64 uint64

	// //Alias
	// var enteroByte byte //alias para uint8
	// var enteroRune rune // alias para int32

	// //Tipos dependiente de la plataforma
	// var enteroUint uint //32 o 64 bits
	// var enteroInt int //32 o 64 bits
	// var enteroUintptr uintptr //Entero sin signo que puede contener el valor de un puntero

	/************************************/
	//Conversión entre tipos

	// Suma int32 y int64
	entero32 = 25
	entero64 = 85

	//fmt.Println(entero32 + entero64)  No funcionaría
	fmt.Println(entero32 + int32(entero64))
	fmt.Println(unsafe.Sizeof(entero32), unsafe.Sizeof(entero64))


	fmt.Println("********************************************************")

	// TIPOS FLOAT
	var f32 float32 // 6 digitos decimales de precision
	var f64 float64 //15 digitos decimales de precision
	var c64 complex64 //Numero complejo para float32
	var c128 complex128 //Numero complejo para float64

	f32 = 0.156
	f64 = 0.156

	c64 = complex(5, 6)
	c128 = complex (5,6)
	fmt.Println("f32, f64, c64, c128 =", f32, f64, c64, c128)

	fmt.Println("********************************************************")

	// TIPOS BOOL

	var resultado bool 

	resultado = 5 < 6
	fmt.Println("5 < 6 = ", resultado)

	fmt.Println("********************************************************")

	// TIPO STRING
	// STRING SON INMUTABLES
	// STRING SIEMPRE CON COMILLAS DOBLES O ACENTOS INVERTIDOS ``
	var nombre string
	nombre = "Eduardo"
	fmt.Println(nombre)

	fmt.Println("Longitud nombre =", len(nombre))

	fmt.Println("Imprimir un caracter =", string(nombre[2]))

	fmt.Println("Imprimir substring = ", nombre[0:3])
	fmt.Println("Imprimir desde el principio = ", nombre[:4])
	fmt.Println("Imprimir hasta el final = ", nombre[2:])
	nombre=nombre+" Romero"
	fmt.Println("Concatenada = ", nombre)
	nombre+=" Garrido"
	fmt.Println("Concatenada =", nombre)
	//Con acentos invertidos es como <pre></pre>, pilla todo tal cual
	cadena := `
	<html>
		<head>
			<meta chartset="utf8">
			<title></title>
		</head>
		<body>
		</body>
	</html>
	`
	fmt.Println(cadena)
	edad := 29
	nombre+="Con edad "+strconv.Itoa(edad)
	fmt.Println(nombre)
	
}