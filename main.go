package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
			//	fmt.Println(parametrosAux)
			//	fmt.Println(len(parametrosAux))
			if len(parametrosAux) == 2 {
				if len(strings.Split(parametrosAux[0], "-")) == 2 {
					parametros[contadorP] = parametro{strings.ToLower(strings.Split(parametrosAux[0], "-")[1]), parametrosAux[1]}
					contadorP++
				} else {
					printError("Parece que falta un - en cerca de " + parametrosAux[0])
				}

			} else {
				printError("Hubo un parametro las escrito cerca de " + parametrosAux[0])
			}
			//fdisk -size->10240 -unit->m -path->/home/teitab/12.disk -type->L -fit->ff -delete->full -name->parte2 -add->13

		}
		//fmt.Println(parametros)
	} else {
		printError("Error con el numero de parametros")
	}
	return parametros
}

//Mkdisk -Size->3000 -unit->K -path->/home/teitan67/MIA_PY1/Disco1.dsk

func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}
