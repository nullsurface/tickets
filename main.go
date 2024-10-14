package main 

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "tickets/api/v1"
)

func main() {
  // Setup file server
  pages := http.FileServer(http.Dir("./pages"))
	http.Handle("/", pages)

  // Setup API
  router := gin.Default()
  router.GET("/ticket-number", v1.GetTicketNumber)
  go router.Run()

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
