package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando Maquina 4 (Servidor)")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Maquina 4 recibió una conexión")
		fmt.Fprint(w, "Hola desde maquina 4")
	})

	fmt.Println("Servidor escuchando en http://localhost:8084")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		fmt.Printf("Error iniciando servidor: %v\n", err)
	}
}
