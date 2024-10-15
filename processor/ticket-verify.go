package processor

import (
	"encoding/csv"
	"log"
	"os"
)

func ValidTicketID(ticketID int) bool {
	f, err := os.OpenFile(TICKET_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for id := range data {
		if id == ticketID {
			return true
		}
	}

	return false
}
