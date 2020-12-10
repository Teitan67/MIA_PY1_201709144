package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

	fmt.Println("Creando disco...")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var temporal int8 = 0
	s := &temporal
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	if unit == "m" || unit == "M" {
		//fmt.Println("Se creo en megas")
		size = size * 1024 * 1024
	} else {
		size = size * 1024
		//fmt.Println("Se creo en kibtytes")
	}

	for i := 0; i < size; i++ {
		escribirBytes(file, binario.Bytes())
	}
	particionGenerica := particion{false, "", "", 0, 0, ""}
	fecha := time.Now().Format(time.RFC822)
	fmt.Println("Creando mbr, fecha obtenida ", fecha, "...")
	mbrDisco := mbr{size, fecha, rand.Intn(1000), fit, particionGenerica, particionGenerica, particionGenerica, particionGenerica}
	//fmt.Println(mbrDisco)
	escribirDisco(path, mbrDisco)

	fmt.Println("Disco creado exitosamente...")

}

/*
type particion struct {
	status bool
	tipo   string
	fit    string
	start  int
	size   int
	name   string
}

*/
func escribirDisco(path string, mbrDisco mbr) {
	fmt.Println("Escribiendo en el disco...")
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, 0)
	var bufferproducto bytes.Buffer
	//fmt.Println(productoTemporal, bufferproducto)
	binary.Write(&bufferproducto, binary.BigEndian, &mbrDisco)
	escribirBytes(file, bufferproducto.Bytes())

	defer file.Close()
}

func rmdisk(commandArray []string) {
	path := ""
	fmt.Println("Eliminando disco...")
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
			}
		}
	}
	err := os.Remove(path)
	if err != nil {
		printError("Error al eliminar " + path + " ...")
	} else {
		fmt.Println("Archivo " + path + " eliminado exitosamente!")
	}

}

func fdisk(commandArray []string) {
	//Atributos
	//solo 4 particiones, 1 extendida por disco
	size := ""
	unit := "K"
	path := ""
	tipo := "P"
	fit := "WF"
	delete := ""
	name := ""
	add := ""

	fmt.Println("Creando particion...")
	var parametros [8]parametro
	//fmt.Println(commandArray)
	parametros = obtenerParametros(commandArray)
	parametroAux := parametro{}
	for i := 0; i < len(parametros); i++ {
		parametroAux = parametros[i]
		if len(parametroAux.tipo) > 0 {
			//fmt.Println("Tipo:", parametroAux.tipo, " Valor:", parametroAux.valor)
			switch parametroAux.tipo {
			case "size":
				size = parametroAux.valor
				break
			case "unit":
				unit = parametroAux.valor
				break
			case "path":
				path = parametroAux.valor
				break
			case "type":
				tipo = parametroAux.valor
				break
			case "fit":
				fit = parametroAux.valor
				break
			case "delete":
				delete = parametroAux.valor
				break
			case "name":
				name = parametroAux.valor
				break
			case "add":
				add = parametroAux.valor
				break
			}
		}
	}
	//fmt.Println(parametros)
	fmt.Println(size, unit, path, tipo, fit, delete, name, add)

}

func rep(commandArray []string) {

	name := ""
	path := ""
	id := ""

	fmt.Println("Iniciando creacion de repore...")
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
			case "name":
				name = parametroAux.valor
				break
			case "id":
				id = parametroAux.valor
				break
			}
		}
	}
	if name != "" && path != "" && id != "" {
		if name == "mbr" {
			//Crear reporte
		} else if name == "disk" {
			//Crear report
		} else {
			printError("name no tiene el parametro correcto :" + name)
		}

	} else {
		printError("Los parametros no estan completos...")
	}
	fmt.Println(name, path, id)
}

/*
rep -id->vda1 -Path->/home/user/reports/reporte1.jpg -name->mbr
rep -id->vda2 -Path->/home/user/reports/reporte2.pdf -name->disk
*/
