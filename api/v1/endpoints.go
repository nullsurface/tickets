package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "log"
	"tickets/processor"
  "strconv"
)

func GetNewTicketID(c *gin.Context) {
  res := map[string]int{"id": processor.GetNewTicketNumber()}
	c.IndentedJSON(http.StatusOK, res)
}

func VerifyTicketID(c *gin.Context) {
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
