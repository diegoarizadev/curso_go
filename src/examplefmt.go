package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println("Hola Mundo desde Madrid")

	nombre := "Diego Ariza" //%s para string
	edad := 19              //%d int
	//%v si no se puede identificar el tipo de dato

	fmt.Printf("%s tiene más de %d años de edad\n", nombre, edad)
	fmt.Printf("%v, tiene más de %v motos\n", nombre, edad)

	message := fmt.Sprintf("%s tiene más de %d años de edad\n", nombre, edad) //fmt como una variable
	fmt.Println(message)

	fmt.Printf("El tipo de dato de la variable nombre es %T \n", nombre)
	fmt.Printf("El tipo de dato de la variable edad es %T \n", edad)

}
