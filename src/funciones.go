package main

import "fmt"

func normalFunction(message string) {
	fmt.Println("Texto función nomalFunction \n")
}

func treeFunction(message string, a, b int) {
	fmt.Printf("Texto función treeFunction con tres parametros : %s, %d, %d \n", message, a, b)
}

func returnFunction(a, b int) int {
	return a * b
}

func returnDosFunction(a, b int) (c, d int) {
	return a * b, b + a
}

func main() {
	texto := "Desde Madrid España"
	normalFunction(texto)
	treeFunction("ESto es un string", 10, 300343)
	fmt.Printf("Función de retorno : %d \n", returnFunction(10, 44))
	value1, value2 := returnDosFunction(10, 67)
	fmt.Printf("Función con doble retorno : 1 -> %d, 2 -> %d \n", value1, value2)
	value3, _ := returnDosFunction(23, 67) // Al marcar el parametro con _ se descarta
	fmt.Printf("Función con doble retorno y un solo valor : 3 -> %d \n", value3)
}
