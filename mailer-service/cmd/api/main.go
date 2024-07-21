package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = "80"

func main() {
	log.Printf("starting mailer service on port %s\n", port)
	srv := newServer().Router

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("mailer service closed")
	} else if err != nil {
		log.Println("error starting mailer service: ", err)
		os.Exit(1)
	}
}
