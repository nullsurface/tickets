package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "log"
	"tickets/processor"
  "sync"
  "strconv"
)

func GetNewTicketID(m *sync.Mutex) gin.HandlerFunc {
  return func(c *gin.Context) {
    m.Lock()
    defer m.Unlock()

    res := map[string]int{"id": processor.GetNewTicketNumber()}
    c.IndentedJSON(http.StatusOK, res)
  }
}

func VerifyTicketID(m *sync.Mutex) gin.HandlerFunc {
  return func(c *gin.Context) {
    m.Lock()
    defer m.Unlock()

    ticketID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
      log.Print(err)
      c.String(http.StatusNotFound, "Invalid Input")
      return 
    }
    
    if processor.ValidTicketID(ticketID) {
      c.String(http.StatusOK, "Valid Ticket ID")   
    } else {
      c.String(http.StatusNotFound, "Ticket ID Not Found")
    }
  }
}
