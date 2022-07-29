package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {

	x := 10 //O tipo da variável "x" é "int", enquanto que o tipo
	//da variável "b" é "hotdog", por mais que ela tenha
	//um valor do tipo "int" atribuído.

	x = int(b) //Estamos transformando o valor da variável "b", que é
	//do tipo "hotdog", em um tipo "int".

	//O GO trabalha apenas com conversão, e não com casting ou coerção.

	fmt.Println(x)

}
