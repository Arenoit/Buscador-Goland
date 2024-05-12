package main

import "fmt"

func main() {
	//Array
	colores := [...]string{
		0: "Rojo",
		1: "Verde",
		2: "Negro",
		3: "Azul",
	}
	fmt.Println("Array: ", colores, len(colores))
	//SubArray
	subMoneda := colores[:3]
	fmt.Println("subArray", subMoneda)

	//Slicen
	numeros := []int{1, 2, 3}
	fmt.Println("Slicen: ", numeros, len(numeros))
	//Agregar Datos
	numeros = append(numeros, 5)
	fmt.Println("append Slicen: ", numeros, len(numeros))
	//Sub Slicen
	subSlicen := numeros[:2]
	numeros[0] = 100
	fmt.Printf("subSlicen: %v\nCapacidad del Slicen:  %v \nReferencia: %p\n", subSlicen, cap(numeros), numeros)
	//Slicen size statement
	numero := make([]int, 3, 3) //Volver a definir un slicen, length and capacity
	numero[0] = 100
	numero[1] = 200
	numero[2] = 300
	numero = append(numero, 400)
	fmt.Println("Pre-emptied Slicen: ", numero)

	//Mapas
	dias := make(map[int]string)
	fmt.Println("map: ", dias)
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	fmt.Println("Map: ", dias)
	dias[7] = "Extended"
	fmt.Println("Extended Map: ", dias)
	delete(dias, 7)
	fmt.Println("Del map: ", dias, len(dias))

	estudiantes := make(map[string][]int)
	estudiantes["Alex"] = []int{13, 14, 15}
	estudiantes["David"] = []int{12, 18, 12}
	fmt.Println("Json Map: ", estudiantes)
	fmt.Println("explore in Map: ", estudiantes["Alex"][1])
}
