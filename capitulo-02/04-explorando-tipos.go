package main

import "fmt"

var teste2 = 5

//teste2 = 6
//teste2 = 6 Isso não vai funcionar, pois, em Go, se utilizarmos o "var", apenas podemos redefinir o valor de uma variável dentro de um bloco de código.

func main() {

	//Os tipos em Go são estáticos. Se declararmos o tipo de
	//uma variável, ele não pode ser alterado durante a execução
	//do código.

	fmt.Println("Olá, mundo!", teste2)
}
