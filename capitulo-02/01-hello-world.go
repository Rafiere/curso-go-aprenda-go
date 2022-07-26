package main

import "fmt"

func main() {

	//Todos os tipos em GO implementam uma interface vazia.

	//O "Println" é uma função que pode receber um número
	//indeterminado de argumentos.

	//A função abaixo retorna dois valores, que estamos armazenando nas duas variáveis.
	numeroDeBytes, erros := fmt.Println("Olá, mundo!")

	//Se queremos descartar o valor de uma variável, devemos utilizar
	//o "_".
	numeroDeBytes2, _ := fmt.Println(numeroDeBytes, erros)

	fmt.Println(numeroDeBytes2)

	//Declaração de variáveis:
	x := 16
	y := "string"
	z := true

	fmt.Println(x, y, z)
}
