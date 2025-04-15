package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Iniciando Maquina 1 (Cliente)")
	fmt.Println("Conectando a las otras máquinas...")

	servers := map[string]string{
		"Maquina 2": "http://maquina2:8082",
		"Maquina 3": "http://maquina3:8083",
		"Maquina 4": "http://maquina4:8084",
	}

	time.Sleep(5 * time.Second)

	for name, url := range servers {
		fmt.Printf("Conectando a %s (%s)...\n", name, url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error conectando a %s: %v\n", name, err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error leyendo respuesta de %s: %v\n", name, err)
			continue
		}

		fmt.Printf("Respuesta de %s: %s\n", name, body)
	}
}
