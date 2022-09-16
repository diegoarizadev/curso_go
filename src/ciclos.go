package main

import "fmt"

func main() {
	fmt.Println("--FOR CONDICIONAL--")
	// For en forma Basica.. condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("--FOR WHILE--")
	counter := 10
	//For While
	for counter < 20 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("--FOR ForEver--")
	counterForever := 100
	for { //Se va a ejecutar por siempre!!!
		fmt.Println(counterForever)
		counterForever++
	}

}
