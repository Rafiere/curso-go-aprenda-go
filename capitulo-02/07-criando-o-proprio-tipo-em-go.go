package main

import "fmt"

func main() {

	type hotdog int //Estamos criando o tipo "hotdog", que possui, por trás, o tipo "int". O tipo "int" é o tipo subjacente do tipo "hotdog".

	var x hotdog //Quando criamos o nosso próprio tipo, podemos fazer muitas coisas com eles com tipos que não estão na linguagem.
	x = 10
	y := 10 //Esse valor é do tipo "int". O valor da variável "x" é do tipo "hotdog". Mesmo que o tipo subjacente da variável "hotdog" seja "int", não podemos atribuir um "int" a essa variável.

	fmt.Println("Olá, mundo!", x, y)

}
