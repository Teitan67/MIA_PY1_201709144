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
	printComando("\t" + comando)
	ejecutarComando(commandArray)
}

func ejecutarComando(commandArray []string) {

	comando := strings.ToLower(commandArray[0])
	//fmt.Println("comando:", comando)
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
	case "rmdisk":
		rmdisk(commandArray)
		break
	case "fdisk":
		fdisk(commandArray)
		break
	case "rep":
		rep(commandArray)
		break
	case "mount":
		mount(commandArray)
		break
	case "":
		break
	default:
		printError("!!Comando incorrecto ," + comando + "!! ")
	}

}
