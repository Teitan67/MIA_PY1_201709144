package main

import (
	"fmt"
	"strings"
)

func pause() {
	fmt.Print("Sistema un pausa, presione cualquier tecla para continuar...")
	fmt.Scanf("\n")
}

func exec(commandArray []string) {
	var ruta []string
	ruta = strings.Split(commandArray[1], "->")
	fmt.Println("Ejecutando exec...")
	if len(ruta) > 1 {
		fmt.Println("Ruta obtenida: ", ruta[1])
		contenido := leerArchivo(strings.TrimSpace(ruta[1]))
		if contenido != "" {
			fmt.Println("Archivo abierto...")
			fmt.Println(contenido)
		}

	} else {
		printError("No se ingreso correctamente la ruta!!")
	}
}
