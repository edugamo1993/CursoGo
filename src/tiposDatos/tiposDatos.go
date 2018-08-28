package main

import "fmt"
import "unsafe"

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

}