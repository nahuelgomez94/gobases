package main

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	os.Setenv("PATHCSV", "./tickets.csv")

	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println("ERROR al intentar obtener el total de tickets a Brasil - ", err)
	} else {
		fmt.Println("Total de pasajes a Brazil: ", total)
	}

	viajes, errM := tickets.GetMornings(tickets.AFTERNOON)
	if errM != nil {
		fmt.Println("ERROR al intentar obtener el total de tickets en esa franja horaria - ", errM)
	} else {
		fmt.Println("Total de pasajes a en esa franja horaria es de : ", viajes)
	}

	avg, errP := tickets.AverageDestination("Brazil", 1000)
	if errP != nil {
		fmt.Println("ERROR al intentar obtener el promedio de personas que van al destino - ", errM)
	} else {
		fmt.Println("Promedio de personas que van al destino: ", avg)
	}
}
