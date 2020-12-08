package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("		╔═════════════════════════════════════╗")
	fmt.Println("		║             Proyecto 1              ║")
	fmt.Println("		║          Oscar R. V. Leon           ║")
	fmt.Println("		║              201709144              ║")
	fmt.Println("		╚═════════════════════════════════════╝")
	fmt.Println("#Ingrese salir para finalizar la aplicacion" + "\n")
	interpretar()
	print("\033[H\033[2J")
}

func printError(mensaje string) {
	colorRed := "\033[31m"
	colorWhite := "\033[37m"
	fmt.Println(string(colorRed), mensaje)
	fmt.Print(string(colorWhite))
}

func leerArchivo(ruta string) string {

	datosComoBytes, err := ioutil.ReadFile(ruta)
	if err != nil {
		printError(err.Error())
	}
	// convertir el arreglo a string
	datosComoString := string(datosComoBytes)
	// imprimir el string
	fmt.Println(datosComoString)
	//Se imprimen los valores guardados en el struct
	return datosComoString
}
