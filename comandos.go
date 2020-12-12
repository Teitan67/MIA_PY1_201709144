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
	"unsafe"
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

	fecha := time.Now().Format(time.RFC822)
	fmt.Println("Creando mbr, fecha obtenida ", fecha, "...")
	mbrDisco := mbr{}
	//fmt.Println("CUENTA: ", int(unsafe.Sizeof(mbrDisco)))
	mbrDisco.Fit = strToBts(fit)
	mbrDisco.Fecha = strToBts(fecha)
	mbrDisco.Tamano = int64(size)
	mbrDisco.ID = int64(rand.Intn(4000) + 1000)
	mbrDisco.MbrSize = int64(unsafe.Sizeof(mbrDisco))
	mbrDisco.Memoria = -1
	//fmt.Println(mbrDisco)
	escribirDisco(path, mbrDisco)
	//fmt.Println("CUENTA: ", int(unsafe.Sizeof(mbrDisco)))
	fmt.Println("Disco creado exitosamente...", fit)

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

func mount(commandArray []string) {
	name := ""
	path := ""
	id := ""
	fmt.Println("La montura del disco...")
	var parametros [8]parametro
	parametros = obtenerParametros(commandArray)
	parametroAux := parametro{}
	for i := 0; i < len(parametros); i++ {
		parametroAux = parametros[i]
		if len(parametroAux.tipo) > 0 {
			switch parametroAux.tipo {
			case "path":
				path = parametroAux.valor
				break
			case "name":
				name = parametroAux.valor
				break
			}
		}
	}
	id = generarID(path, name)
	nuevoDiscoMontado := disco{}
	nuevoDiscoMontado.ID = strToBts(id)
	nuevoDiscoMontado.DiscoMontado = strToBts(name)
	nuevoDiscoMontado.Path = strToBts(path)
	nuevoDiscoMontado.EstadoBorrado = false

}

//====================================================

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

//=============================================================
func fdisk(commandArray []string) {
	//Atributos
	//solo 4 particiones, 1 extendida por disco
	size := 0
	unit := "K"
	path := ""
	tipo := "p"
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
				i, err := strconv.Atoi(parametroAux.valor)
				if err != nil {
					printError(err.Error())
				}
				size = i
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
	fmt.Println(size, unit, path, tipo, fit, delete, name, add)
	mbrDisco := leerDisco(path)
	mbrVacio := mbr{}
	if mbrDisco != mbrVacio {
		if delete != "" && add != "" {
			printError("No puedes usar delete y add juntos")
		} else if delete != "" {
			//borra particion
		} else if add != "" {
			//agrega espacio a la particion
		} else {
			// TRBAJANDO
			//Creando particion
			if tipo != "l" {
				fmt.Println(noExtendidas(path))
				if noExtendidas(path) > 0 && tipo == "e" {
					printError("No puedes insertar mas de una extendida")
				} else {
					if verificarPariciones(path) {
						nuevaParticion := particion{}
						nuevaParticion.Tipo = strToBts(tipo)
						nuevaParticion = ff(mbrDisco.Memoria, nuevaParticion, path)
						if nuevaParticion.Start > 0 {
							nuevaParticion.Name = strToBts(name)

							nuevaParticion.Size = realBytes(int64(size), unit)

							nuevaParticion.Fit = strToBts(fit)
							nuevaParticion.Status = false

							fmt.Println("Nueva particion creada...")

							escribirParticion(path, nuevaParticion)

						} else {
							printError("Se intento insertar al inicio")
						}
					}
				}

			} else {
				printError("No se puede crear una particion logica fuera de una extendida")
			}

		}
	} else {
		printError("No existe el disco de esta ruta: " + path)
	}
}

/*FUNCIONES PARA CREACION DE PARTICION*/
func ff(Start int64, nuevaParticion particion, path string) particion {
	if Start == 0 {
		mbrAux := leerDisco(path)
		if mbrAux.Memoria == -1 {
			fmt.Println("\t\tCaso 1:")
			nuevaParticion.Start = mbrAux.MbrSize
			nuevaParticion.Siguiente = mbrAux.Memoria
			mbrAux.Memoria = nuevaParticion.Start
			escribirDisco(path, mbrAux)
			return nuevaParticion
		} else if mbrAux.Memoria-mbrAux.MbrSize > 0 {
			log.Fatal("Valio que no se tenia que activar")
		}
		fmt.Println("\t\tCaso 2")
		return ff(mbrAux.Memoria, nuevaParticion, path)

	}
	particionAux := leerParicion(Start, path)
	if particionAux.Siguiente == -1 {
		fmt.Println("\t\tCaso 1 B")
		nuevaParticion.Start = particionAux.Start + particionAux.Size
		nuevaParticion.Siguiente = particionAux.Siguiente
		particionAux.Siguiente = nuevaParticion.Start
		escribirParticion(path, particionAux)
		return nuevaParticion
	} else if particionAux.Siguiente-particionAux.Start-particionAux.Size > 0 {
		log.Fatal("Valio que no se tenia que activar 2")
	}
	fmt.Println("\t\tCaso 2 B")
	return ff(particionAux.Siguiente, nuevaParticion, path)

}

func reporteDisco() {

}

/*
func ff(Start int64, nuevaParticion particion, path string) particion {
	if Start == 0 {
		mbrAux := leerDisco(path)
		if mbrAux.Memoria == -1 {

			nuevaParticion.Start = mbrAux.MbrSize
			nuevaParticion.Siguiente = -1
			mbrAux.Memoria = nuevaParticion.Start
			fmt.Println("\ncaso1\n")
			fmt.Println(nuevaParticion)
			fmt.Println("\n\n")
			escribirDisco(path, mbrAux)
			return nuevaParticion
		} else if mbrAux.Memoria-Start-mbrAux.MbrSize >= nuevaParticion.Size {
			fmt.Println("\ncaso3\n")
			fmt.Println(nuevaParticion)
			fmt.Println("\n\n")
			nuevaParticion.Siguiente = mbrAux.Memoria
			mbrAux.Memoria = mbrAux.MbrSize
			nuevaParticion.Start = mbrAux.MbrSize
			escribirDisco(path, mbrAux)
			return nuevaParticion
		}
		fmt.Println("\ncaso5\n")
		fmt.Println(nuevaParticion)
		fmt.Println("\n\n")
		return ff(mbrAux.Memoria, nuevaParticion, path)
	}
	particionAux := leerParicion(Start, path)
	if particionAux.Siguiente == -1 {
		fmt.Println("\ncaso 2\n")
		fmt.Println(nuevaParticion)
		fmt.Println("\n\n")
		nuevaParticion.Siguiente = -1
		nuevaParticion.Start = particionAux.Start + particionAux.Size
		particionAux.Siguiente = nuevaParticion.Start
		escribirParticion(path, particionAux)
		return nuevaParticion
	} else if particionAux.Siguiente-particionAux.Start-particionAux.Size >= nuevaParticion.Size {

		fmt.Println("\ncaso6\n")
		fmt.Println(nuevaParticion)
		fmt.Println("\n\n")
		nuevaParticion.Siguiente = particionAux.Siguiente
		nuevaParticion.Start = particionAux.Start + particionAux.Size
		particionAux.Siguiente = nuevaParticion.Start
		escribirParticion(path, nuevaParticion)
		return nuevaParticion

	} else {
		fmt.Println("\ncaso4\n")
		fmt.Println(nuevaParticion)
		fmt.Println("\n\n")
		return ff(particionAux.Siguiente, nuevaParticion, path)
	}
}
*/
