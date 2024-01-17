package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.GenerateRoute()
	fmt.Printf("Watching on port %d \n", config.Port)
	log.Fatal(http.ListenAndServe(":5000", r))
}
