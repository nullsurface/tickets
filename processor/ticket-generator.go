package processor

import (
  "encoding/csv"
  "log"
  "strconv"
)

TICKET_FILE int = "./persistence/tickets.csv"

func GetNextTicketNumber() int {
  f, err =: os.OpenFile(TICKET_FILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)  
  if err != nil {
        log.Fatal(err)
  }
  defer f.Close()

  csvReader := csv.NewReader(f)
  data, err := csvReader.ReadAll()
  if err != nil {
      log.Fatal(err)
  }

  lastIndex, err = strconv.Atoi(data[len(data)-1][0])
  if err != nil {
    log.Fatal(err)
  }
  
  csvWriter := csv.NewWriter(f)
  defer csvWriter.Flush()

  newTicket []string{lastIndex + 1}
  csvWriter.Write(newTicket) 
  return lastIndex + 1
}

