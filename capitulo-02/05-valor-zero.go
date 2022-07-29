package main

import "fmt"

var aaa int //Declaração

func main() {

	aaa = 4 //Inicialização e atribuição (primeira atribuição de valor a uma variável)
	aaa = 5 //Apenas atribuição, pois a variável já foi inicializada anteriormente.

	ccc := 10 //Esse operador declara, inicializa e atribui um valor à variável.

	//Cada tipo da linguagem é inicializado com um valor, como "0", se for
	//integer, 0.0 se for float, e etc.

	//Devemos utilizar o ":=" sempre que possível, e, para declararmos variáveis
	//em nível de pacote, devemos utilizar o "var".

	fmt.Println("Teste", aaa, ccc)

}
