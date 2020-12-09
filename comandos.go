package main

import (
	"fmt"
	"strconv"
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
			var comandos []string
			comandos = strings.Split(contenido, "\n")
			for i := 0; i < len(comandos); i++ {
				if !strings.Contains(comandos[i], "#") {
					leerComando(comandos[i])
				}
			}
		}

	} else {
		printError("No se ingreso correctamente la ruta!!")
	}
}

func mkdisk(commandArray []string) {

	path := ""
	size := 0
	unit := "m"
	fit := "ff"

	var parametros [8]parametro
	parametros = obtenerParametros(commandArray)
	parametroAux := parametro{}
	for i := 0; i < len(parametros); i++ {
		parametroAux = parametros[i]
		if len(parametroAux.tipo) > 0 {
			//fmt.Println("Tipo:", parametroAux.tipo, " Valor:", parametroAux.valor)
			switch parametroAux.tipo {
			case "path":
				path = parametroAux.valor
				break
			case "size":
				i, err := strconv.Atoi(parametroAux.valor)
				if err != nil {
					printError(err.Error())
				}
				size = i
				break
			case "fit":
				fit = parametroAux.valor
				break
			case "unit":
				unit = parametroAux.valor
				break
			}
		}
	}
	//Mkdisk -Size->3000 -unit->K -path->/home/teitan67/MIA_PY1/bin/Disco1.dsk

	fmt.Println(path)
	fmt.Println(size)
	fmt.Println(fit)
	fmt.Println(unit)
}
