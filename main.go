package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tickets/api/v1"
)

func main() {
	// Setup file server
	pages := http.FileServer(http.Dir("./pages"))
	http.Handle("/", pages)

	// Setup API
	router := gin.Default()
	router.GET("/new-ticket", v1.GetNewTicketID)
  router.GET("/verify-ticket/:id", v1.VerifyTicketID)
	go router.Run()

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
