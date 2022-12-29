package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type Ticket struct {
	// 1,Tait Mc Caughan,tmc0@scribd.com,Finland,17:11,785
	id       int
	pasajero string
	email    string
	pais     string
	hora     string
	precio   float64
}

// ejemplo 1
/*
Una función que calcule cuántas personas viajan a un país determinado:
func GetTotalTickets(destination string) (int, error) {}
(ejemplo 1)
Tip: VS Code nos permite buscar una palabra en un archivo con Ctrl + F o ⌘ + F.
*/
func GetTotalTickets(destination string) (registros int, err error) {
	file, err := os.Open("./tickets.csv")
	if err != nil {
		err = errors.New("ERROR: No se encontró la BBDD")
		panic("ERROR: BBDD No encontrada")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'

	fmt.Println("holi")
	for {
		record, e := reader.Read()

		if e != nil {
			break
		}

		if record[3] == destination {
			registros++
		}
	}

	return
}

// ejemplo 2
func GetMornings(time string) (registros int, err error) {
	file, err := os.Open("./tickets.csv")
	if err != nil {
		err = errors.New("ERROR: No se encontró la BBDD")
		panic("ERROR: BBDD No encontrada")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'

	fmt.Println("holi")
	for {
		record, e := reader.Read()

		if e != nil {
			break
		}

		if record[3] == destination {
			registros++
		}
	}

	return
}

// ejemplo 3
func AverageDestination(destination string, total int) (registros int, err error) {
	file, err := os.Open("./tickets.csv")
	if err != nil {
		err = errors.New("ERROR: No se encontró la BBDD")
		panic("ERROR: BBDD No encontrada")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'

	fmt.Println("holi")
	for {
		record, e := reader.Read()

		if e != nil {
			break
		}

		if record[3] == destination {
			registros++
		}
	}

	return
}
