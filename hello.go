package main

import "fmt"
import "reflect"

func main() {
	var nome string = "João"
	fmt.Println("Olá sr.", nome, "sua idade é")
	
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", reflect.TypeOf(comando))
}
