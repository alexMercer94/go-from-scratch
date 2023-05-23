package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-from-scratch/ejercicios"
)

var fileName string = "./files/txtFiles/table.txt"

func SaveTable() {
	var txt string = ejercicios.CreateTable()
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(file, txt)
	file.Close()
}

func SumTableInFile() {
	var txt string = ejercicios.CreateTable()
	if !Append(txt) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(txt string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(txt)

	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo: ", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		reg := scanner.Text()
		fmt.Println(reg)
	}

	file.Close()
}
