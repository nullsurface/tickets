package v1

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetTicketNumber(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, "YUH")
}
