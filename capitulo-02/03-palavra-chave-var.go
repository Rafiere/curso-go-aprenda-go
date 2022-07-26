package main

import "fmt"

var teste = "Teste" //Esse operador declarará uma variável com visibilidade para todos os arquivos ".go" do pacote em que esse arquivo se encontra.

func main() {

	var nome2 = "oi" //Essa variável ficará disponível apenas para esse arquivo, pois ela foi declarada dentro de um bloco de código.

	fmt.Println(nome2)

	x := 10 //O operador ":=" declarará uma variável com visibilidade apenas para esse bloco de código.

	fmt.Println(teste, x)
}
