package main

import (
	"fmt"
)
	var x int
	var y int
	var operacao string
	var nome string
func main (){
		
		fmt.Println("Digite seu nome: ")
		fmt.Scan(&nome)

		fmt.Println("--- SEJA BEM VINDO", nome, "---")
		fmt.Println("Aqui está a sua calculadora :)")

		fmt.Println("Digite o primeiro número inteiro: ")
		fmt.Scan(&x)
		fmt.Println("Digite o segundo número inteiro: ")
		fmt.Scan(&y)
		fmt.Println("Escolha a operação desejada (+ - * /): ")
		fmt.Scan(&operacao)

		switch operacao {
		case "+":
			fmt.Println("Aqui está o resultado da sua operação", nome, ":\n", x + y)	
		case "-":
			fmt.Println("Aqui está o resultado da sua operação", nome, ":\n", x - y)	
		case "*":
			fmt.Println("Aqui está o resultado da sua operação", nome, ":\n", x * y)	
		case "/":
			if y == 0 || x == 0 {
				fmt.Println("Impossivel dividir por zero!")
			} else {
				fmt.Println("Aqui está o resultado da sua operação", nome, ":\n", x/y)
			}
		}
	}

			
		

