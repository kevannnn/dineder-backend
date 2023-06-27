package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevannnn/dineder-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 3000 at http://localhost:3000!")

	log.Fatalln(http.ListenAndServe(":3000", r))
}
