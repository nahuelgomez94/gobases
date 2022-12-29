package tickets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	os.Setenv("PATHCSV", "../../tickets.csv")

	tickets, err := GetTotalTickets("Brazil")

	if err != nil {
		t.Fatal("Ni siquiera pudo ejecutarse GetTotalTickets - ERROR" + err.Error())
	}

	assert.Equal(t, 45, tickets)
}

func TestGetMornings(t *testing.T) {
	os.Setenv("PATHCSV", "../../tickets.csv")

	viajes, err := GetMornings(AFTERNOON)

	if err != nil {
		t.Fatal("No se pudo ejecutar el programa - ERROR " + err.Error())
	}

	assert.Equal(t, 289, viajes)
}

func TestAverageDestination(t *testing.T) {
	os.Setenv("PATHCSV", "../../tickets.csv")

	avg, err := AverageDestination("Brazil", 1000)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, 4.5, avg)
}
