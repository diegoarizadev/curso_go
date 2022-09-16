package main

import "fmt"

func main() {

	const baseCuadrado = 20
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadraro : ", areaCuadrado)

	x := 10
	y := 40

	result := x + y
	fmt.Println("Suma : ", result)

	result = x - y
	fmt.Println("Resta : ", result)

	result = x * y
	fmt.Println("Multiplicacion : ", result)

	result = x / y
	fmt.Println("Division : ", result)

	result = x % y
	fmt.Println("Modulo : ", result)

	x++
	fmt.Println("Incremento : ", x)

	x--
	fmt.Println("Decremento : ", x)
}
