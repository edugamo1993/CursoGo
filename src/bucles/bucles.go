package main

import(
	"fmt"
	"time"
)

func main(){

	//BUCLE FOR
	// Se puede escribir de dos formas


	i:=0
	for i<10 { 
		fmt.Println(i) 
		i++ //Aumenta la variable
		time.Sleep(100 * time.Millisecond) //Espera 100 milisegundos
	}

	for j:=0; j<10; j++ { //Otra manera de declararlo
		fmt.Println(j) 
		time.Sleep(1000 * time.Millisecond) //Espera un segundo
	}

	for{
		fmt.Println("555")
		break
	}
	
}