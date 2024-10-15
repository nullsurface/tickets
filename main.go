package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tickets/api/v1"
  "sync"
)


func main() {
	// Setup file server
	pages := http.FileServer(http.Dir("./pages"))
	http.Handle("/", pages)

	// Setup API
  var m sync.Mutex
	router := gin.Default()
	router.GET("/new-ticket", v1.GetNewTicketID(&m))
  router.GET("/verify-ticket/:id", v1.VerifyTicketID(&m))
	go router.Run()

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
