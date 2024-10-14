package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tickets/processor"
)

func GetTicketNumber(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, processor.GetNextTicketNumber())
}
