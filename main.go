package main

import "fmt"

func main() {
	n1 := 10
	n2 := n1
	n3 := &n1
	fmt.Println("\nPrimeiro número: ", n1)
	fmt.Println("Copia do n1: ", n2)
	fmt.Println("Ponteiro do primeiro número: ", n3)
	fmt.Println("Ponteiro do primeiro número desreferenciado: ", *n3)

	n1 = 12
	fmt.Println("\nPrimeiro número: ", n1)
	fmt.Println("Copia do n1: ", n2)
	fmt.Println("Ponteiro do primeiro número: ", n3)
	fmt.Println("Ponteiro do primeiro número desreferenciado: ", *n3)

	*n3 = 21 //Alterando valor de n1 através do ponteiro
	fmt.Println("\nPrimeiro número: ", n1)
	fmt.Println("Copia do n1: ", n2)
	fmt.Println("Ponteiro do primeiro número: ", n3)
	fmt.Println("Ponteiro do primeiro número desreferenciado: ", *n3)
}

	

