package main

import (
	"fmt"
	"primeiroProjeto/calc"
)

func main() {

	var num1 int = 15
	//passando zero 0 no segundo valor ele exibira era uma pequena tratativa no pacote
	//calc.go
	num2 := 0

	resultado, err := calc.Dividir(num1, num2)

	if err != nil {
		fmt.Println("Erro: ", err)
		//utilizar return do que o if else ele vai parar a execução da função assim
		//que chegar no return somente
		return
	}

	//fmt utilizado para printar resultado/informação
	fmt.Println(resultado)

}

/*
func dividir(a, b int) int {
	return a / b
} */
