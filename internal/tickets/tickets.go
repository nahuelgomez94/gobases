package tickets

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id       int
	pasajero string
	email    string
	pais     string
	hora     string
	precio   float64
}

/*
Requerimiento 1: Una función que calcule cuántas personas viajan a un país determinado:
func GetTotalTickets(destination string) (int, error) {}
(ejemplo 1)
Tip: VS Code nos permite buscar una palabra en un archivo con Ctrl + F o ⌘ + F.
*/
func GetTotalTickets(destination string) (registros int, err error) {
	var pathCsv, _ = os.LookupEnv("PATHCSV")

	file, err := os.Open(pathCsv)
	if err != nil {
		err = errors.New("ERROR: No se encontró la BBDD")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'

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

var (
	EARLY_MORNING = "MADRUGADA"
	MORNING       = "MAÑANA"
	AFTERNOON     = "TARDE"
	EVENING       = "NOCHE"
)

/*
Requerimiento 2: Una o varias funciones que calcule cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19) y noche (20 → 23):
func GetCountByPeriod(time string) (int, error) {}
(ejemplo 2)
Tip: En Go para manipular caracteres tenemos el paquete strings.
*/
func GetMornings(time string) (registros int, err error) {
	var pathCsv, _ = os.LookupEnv("PATHCSV")
	file, errOpen := os.Open(pathCsv)
	if errOpen != nil {
		err = errors.New("ERROR: No se encontró la BBDD")
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'
	for {
		record, e := reader.Read()

		if e != nil {
			break
		}

		hora := record[4]
		hora_minuto := strings.Split(hora, ":")
		//		hora_parse, errParse := strconv.ParseInt(hora_minuto[0], 10, 64)

		hora_parse, errParse := strconv.Atoi(hora_minuto[0])
		if errParse != nil {
			err = errors.New(errParse.Error())
			break
		}

		switch {
		case time == EARLY_MORNING && hora_parse >= 0 && hora_parse <= 6:
			registros++
		case time == MORNING && hora_parse >= 7 && hora_parse <= 12:
			registros++
		case time == AFTERNOON && hora_parse >= 13 && hora_parse <= 19:
			registros++
		case time == EVENING && hora_parse >= 20 && hora_parse <= 23:
			registros++
		}
	}

	return

}

/*
Requerimiento 3:
Calcular el promedio de personas que viajan a un país determinado en un dia:
func AverageDestination(destination string, total int) (float64, error) {}
(ejemplo 3)
Tip: El promedio de x se calcula como: x̄ =  xn
*/
func AverageDestination(destination string, total int) (avg float64, err error) {
	var pathCsv, _ = os.LookupEnv("PATHCSV")
	file, err := os.Open(pathCsv)
	if err != nil {
		err = errors.New("ERROR: No se encontró la BBDD")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'

	registros, total := 0, 0
	for {
		record, e := reader.Read()

		if e != nil {
			break
		}

		total++

		if record[3] == destination {
			registros++
		}
	}

	avg = 100 / float64(total) * float64(registros)

	return
}
