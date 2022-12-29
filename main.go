package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")

	if err != nil {
		fmt.Println("Opa")
	} else {
		fmt.Println("Total: ", total)
	}
}
