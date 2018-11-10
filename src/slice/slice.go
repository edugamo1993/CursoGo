package main

import (
	"fmt"
)

func main() {
	fmt.Println("Empieza el programa. SLICES")
	//Declaración
	var j []int
	fmt.Println("Slice j: ", j)
	x := []int{1, 4}
	fmt.Println("Slice x: ", x)

	y := make([]int, 5) //make inicializa un slice con el número de elementos que le pasamos
	fmt.Println(y)
	fmt.Println("Longitud: ", len(y))
	fmt.Println("Capacidad: ", cap(y))

	k := make([]int, 5, 10) // (tipo, logitud, capacidad)
	fmt.Println(k)
	fmt.Println("Longitud: ", len(k))
	fmt.Println("Capacidad: ", cap(k))

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var a, b []int
	a = ar[1:5]
	b = ar[6:10] //Apuntan a un array
	fmt.Println(a)
	fmt.Println(b)

	b[0] = 25
	fmt.Println("Slice b", b)

	fmt.Println("Array ar: ", ar) //Se cambia en "ar" también
	fmt.Println("Capacidad a :", cap(a))
	fmt.Println("Capacidad b :", cap(b)) //Desde donde empieza de ar, hasta el final de ar

	nuevoSlice := make([]byte, 4, 10)
	fmt.Println(nuevoSlice)

	nuevoSlice = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(nuevoSlice)

	fmt.Printf("Slice x: %q \n", nuevoSlice)

	for i := 0; i < len(nuevoSlice); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, nuevoSlice[i])
	}

	// nuevoSlice[5] = ' '    Esto daría error
	nuevoSlice = append(nuevoSlice, ' ')
	nuevoSlice = append(nuevoSlice, 'M', 'U', 'N', 'D', 'O')
	fmt.Printf("Slice : %q \n", nuevoSlice)
	fmt.Println(cap(nuevoSlice))

	/// COPY

	origen := []int{1, 2, 3}
	destino := []int{3, 4, 5}
	copy(destino, origen)
	fmt.Println(origen, destino) //Modifica destino

	origen2 := []int{1, 2, 3}
	destino2 := make([]int, 2)
	copy(destino2, origen2)
	fmt.Println(origen2, destino2) //solo copia los 2 primeros ya que destino tiene 2 huecos ahora mismo.

	// Origen < Destino
	origen3 := []int{1, 2}
	destino3 := []int{3, 4, 5}
	copy(destino3, origen3)
	fmt.Println(origen3, destino3) // Copia los 2 huecos del origen3 y los mete en los dos primeros de destino3. El resto queda intacto.
}
