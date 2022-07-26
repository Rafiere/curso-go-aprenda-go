package main

import "fmt"

//O "var" deve ser usado fora de blocos de código, pois ele declara
//variáveis globalmente, enquanto que o operador de declaração ":=" deve
//ser utilizado apenas dentro de blocos de código, pois ele apenas declara
//variáveis em um escopo.
var nome = "testando"

func main() {

	//O operador ":=" serve para declararmos uma variável.
	//Esse operador possui a tipagem automática, assim, ele infere
	//O tipo de acordo com o valor que ele receber.

	x := 10      //Está sendo inferido o tipo numérico.
	y := "teste" //Está sendo inferido o tipo string.

	fmt.Println(x, y)

	//Estamos realizando a atribuição do valor "20" à variável "x" já existente.
	x = 20

	fmt.Println(x)

	//O operador ":=" deve ser usado para declararmos ao menos uma variável. Se
	//tentarmos utilizar o operador ":=" para apenas atribuirmos valores a
	//variáveis que já foram declaradas, um erro de compilação será gerado.

	x, z := 20, 30

	//Se tivermos apenas variáveis que já foram declaradas, devemos utilizar
	//apenas o operador "=".
	x, z = 40, 60 //Atribuindo os valores "40" e "60" às variáveis "x" e "z" que já estão declaradas.

	fmt.Println(x, z)

	//Operadores:
	numero := 10 + 10

	fmt.Println(numero)

	//O operador "==" fará uma comparação lógica entre dois valores.
	resultadoDaExpressao := 10 == 10

	fmt.Println(resultadoDaExpressao)
}
