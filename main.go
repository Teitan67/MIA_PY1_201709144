package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	//fmt.Println(datosComoString)
	//Se imprimen los valores guardados en el struct
	return datosComoString
}

type parametro struct {
	tipo  string
	valor string
}

func obtenerParametros(comando []string) [8]parametro {
	var parametros [8]parametro
	contadorP := 0
	if len(comando) > 0 {
		for i := 1; i < len(comando); i++ {
			cadenaTemporal := strings.TrimSpace(comando[i])
			parametrosAux := strings.Split(cadenaTemporal, "->")

			parametros[contadorP] = parametro{strings.ToLower(strings.Split(parametrosAux[0], "-")[1]), parametrosAux[1]}
			contadorP++
		}
		//fmt.Println(parametros)
	}
	return parametros
}

//Mkdisk -Size->3000 -unit->K -path->/home/teitan67/MIA_PY1/bin/Disco1.dsk
