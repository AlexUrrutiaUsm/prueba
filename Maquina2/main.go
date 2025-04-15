package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando Maquina 2 (Servidor)")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Maquina 2 recibió una conexión")
		fmt.Fprint(w, "Hola desde maquina 2")
	})

	fmt.Println("Servidor escuchando en http://localhost:8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("Error iniciando servidor: %v\n", err)
	}
}
