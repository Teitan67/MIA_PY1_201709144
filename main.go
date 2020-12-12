package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"
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
func printComando(mensaje string) {
	colorBlue := "\033[34m"
	colorWhite := "\033[37m"
	fmt.Println(string(colorBlue), mensaje)
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

func leerDisco(path string) mbr {
	fmt.Println("Leyendo disco...")

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		printError(err.Error())
	}

	//Creamos una variable temporal que nos ayudará a leer los productos
	mbrAux := mbr{}
	//obtenemor el size del producto para saber cuantos bytes leer
	var size int = int(unsafe.Sizeof(mbrAux))

	file.Seek(0, 0)
	mbrAux = obtenerMbr(file, size, mbrAux)

	defer file.Close()
	return mbrAux
}

func obtenerMbr(file *os.File, size int, mbrAux mbr) mbr {
	data := leerBytes(file, size)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &mbrAux)
	if err != nil {
		printError("Lectura de binario fallida \n" + err.Error())
	}
	return mbrAux
}
func obtenerParticion(file *os.File, size int, paricionAux particion) particion {
	data := leerBytes(file, size)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &paricionAux)
	if err != nil {
		printError("Lectura de binario fallida \n" + err.Error())
	}
	return paricionAux
}
func leerParicion(start int64, path string) particion {
	fmt.Println("Leyendo disco...")

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		printError(err.Error())
	}

	//Creamos una variable temporal que nos ayudará a leer los productos
	particionAux := particion{}
	//obtenemor el size del producto para saber cuantos bytes leer
	var size int = int(unsafe.Sizeof(particionAux))

	file.Seek(start, 0)
	particionAux = obtenerParticion(file, size, particionAux)

	defer file.Close()
	return particionAux
}
func leerBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func strToBts(str string) [25]byte {
	var bts [25]byte
	copy(bts[:], str)
	return bts
}
func btsToStr(bts [25]byte) string {
	var str string
	str = string(bts[:])
	return str
}

func realBytes(size int64, unidad string) int64 {
	switch unidad {
	case "b":
		size = size * 1
		break
	case "k":
		size = size * 1 * 1024
		break
	case "m":
		size = size * 1 * 1024 * 1024
		break
	}
	return size
}

func escribirDisco(path string, mbrDisco mbr) {
	fmt.Println("Actualizando mbr...")
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, 0)
	var bufferproducto bytes.Buffer
	binary.Write(&bufferproducto, binary.BigEndian, &mbrDisco)
	escribirBytes(file, bufferproducto.Bytes())
	defer file.Close()
}

func escribirParticion(path string, nuevaParticion particion) {
	fmt.Println("Actualizando particion...")
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(nuevaParticion.Start, 0)
	var bufferproducto bytes.Buffer
	binary.Write(&bufferproducto, binary.BigEndian, &nuevaParticion)
	escribirBytes(file, bufferproducto.Bytes())

	defer file.Close()
}

func verificarPariciones(path string) bool {
	mbrAux := leerDisco(path)
	noParticiones := 0
	var siguiente int64 = mbrAux.Memoria

	for siguiente > 0 {
		//fmt.Println("==========================================>", siguiente)
		particionAux := leerParicion(siguiente, path)
		siguiente = particionAux.Siguiente
		fmt.Println("==========================================>", btsToStr(particionAux.Name))
		noParticiones++
		if noParticiones > 3 {
			printError("Haz llegado al numero maximo de particion!!!")
			return false
		}
	}
	return true
}
func noExtendidas(path string) int {
	mbrAux := leerDisco(path)
	noExtendidas := 0
	var siguiente int64 = mbrAux.Memoria
	for siguiente > 0 {
		particionAux := leerParicion(siguiente, path)
		siguiente = particionAux.Siguiente
		tipo := particionAux.Tipo
		if tipo[0] == 101 {
			noExtendidas++
		}
	}
	return noExtendidas
}

var arrayMontados []disco

func agregarMontada(nuevoDisco disco) {
	fmt.Println("Particion montada en el sistema")
	arrayMontados = append(arrayMontados, nuevoDisco)
}
func eliminarMontada(id string) {
	for i := 0; i < len(arrayMontados); i++ {
		discoAux := arrayMontados[i]
		if discoAux.ID == strToBts(id) {
			discoAux.EstadoBorrado = true
			arrayMontados[i] = discoAux
			fmt.Println("Particion desmontada del sistema")
			break
		}
	}
}

func generarID(path string, name string) string {
	var noMount int = 1
	var lMount int = 65
	for i := 0; i < len(arrayMontados); i++ {
		path1 := arrayMontados[i].Path
		path2 := strToBts(path)
		if compararBytes(path1, path2) {
			lMount++
		}
		noMount++
	}
	return "VD" + string(lMount) + strconv.Itoa(noMount)

}

func compararBytes(b1 [25]byte, b2 [25]byte) bool {
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
