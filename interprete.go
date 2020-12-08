package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func interpretar() {
	reader := bufio.NewReader(os.Stdin)
	finalizar := false
	for !finalizar {
		fmt.Print("Ingrese comando:")
		comando, _ := reader.ReadString('\n')
		if strings.ToLower(comando) == "salir\n" {
			finalizar = true
		} else {
			if comando != "" {
				leerComando(comando)
			}
		}
	}
}

func leerComando(comando string) {
	var commandArray []string
	commandArray = strings.Split(comando, " ")
	ejecutarComando(commandArray)
}

func ejecutarComando(commandArray []string) {

	comando := strings.ToLower(commandArray[0])
	switch comando {
	case "exec":
		exec(commandArray)
		break
	case "pause":
	case "pause\n":
		pause()
		break
	case "#":
		break
	default:
		printError("!!Comando incorrecto!!")
	}
}
