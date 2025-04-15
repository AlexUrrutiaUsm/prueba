package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando Maquina 3 (Servidor)")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Maquina 3 recibió una conexión")
		fmt.Fprint(w, "Hola desde maquina 3")
	})

	fmt.Println("Servidor escuchando en http://localhost:8083")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Printf("Error iniciando servidor: %v\n", err)
	}
}
