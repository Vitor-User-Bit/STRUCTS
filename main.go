package main

import "fmt"
	
type End struct {
	Cidade string
	Estado string
}
type Pessoa struct {
	Nome string
	Idade int
	Profissao string
	Endereco End
}

func main (){
	pessoa1 := Pessoa{
		Nome: "Vitor",
		Idade:  30,
		Profissao: "Programador",
		Endereco: End{
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}
			fmt.Println("Struct completa: ") // PRINTANDO A STRUCT COMPLETA
			fmt.Println(pessoa1,"\n")


			//PEDINDO ESPECIFICAMENTE

			fmt.Println("Nome: ", pessoa1.Nome)
			fmt.Println("Nome: ", pessoa1.Endereco.Cidade)

	pessoa1.Profissao = "Professor"

			fmt.Println("Nova profissão: ", pessoa1.Profissao)

	var pessoa2 Pessoa
			fmt.Println(pessoa2)

	pessoa2.Nome = "Maria"
	pessoa2.Idade = 25
	pessoa2.Profissao = "Designer"
	pessoa2.Endereco.Cidade = "Brasilia"
	pessoa2.Endereco.Estado = "DF"

			fmt.Println("\nApós preencher os campos: ")
			fmt.Println(pessoa2)

	pessoas := []Pessoa{
		pessoa1,
		pessoa2,
			{
				Nome: "carlos",
			Idade:  40,
			Profissao: "Engenheiro",
			Endereco: End{
				Cidade: "São Paulo",
				Estado: "SP",
			},
		},
	}

	for i, pessoa := range pessoas{

		fmt.Println("Pessoa", i+1)

		fmt.Println("Nome: ", pessoa.Nome)
		fmt.Println("Idade: ", pessoa.Idade)
		fmt.Println("Profissão: ", pessoa.Profissao)
		fmt.Println("Cidade: ", pessoa.Endereco.Cidade)
		fmt.Println("Estado: ", pessoa.Endereco.Estado)
	}


}

			




