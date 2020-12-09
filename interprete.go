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
	fmt.Println(comando)
	ejecutarComando(commandArray)
}

func ejecutarComando(commandArray []string) {

	comando := strings.ToLower(commandArray[0])
	//fmt.Println(comando)
	switch comando {
	case "exec":
		exec(commandArray)
		break
	case "pause":
		fmt.Println("se pauso")
		pause()
		break
	case "pause\n":
		fmt.Println("se pauso")
		pause()
		break
	case "mkdisk":
		mkdisk(commandArray)
		break
	default:
		printError("!!Comando incorrecto!!")
	}

}
