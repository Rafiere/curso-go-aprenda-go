package main

import "fmt"

func main() {

	x := "Olá, \ntudo bem\t?" //Isso é um string interpretado, pois possui caracteres como o "\n".

	y := `Olá, \ntudo bem?\t?` //Isso é um raw string, ou seja, nada do que está dentro dessa string será interpretado, como os caracteres "\n".

	fmt.Println(x, y)

	//Prints:
	fmt.Println("Olá, mundo!") //Uma linha será pulada no final.
	fmt.Print("Olá, mundo!")   //Não pularemos uma linha no final.
	z := fmt.Sprint("teste")   //Esse print não exibirá nada na tela. Ele retornará a string que inserimos dentro dele como retorno da função.

	fmt.Println(z)
}
