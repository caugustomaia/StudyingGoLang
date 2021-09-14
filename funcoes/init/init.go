package main

import "fmt"

func init() { // é executada mesmo sem ser chamada e pode ter 1 para cada arquivo
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
