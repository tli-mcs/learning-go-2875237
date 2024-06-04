package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tli-mcs/learning-go-2875237/practice/handler" // Import the missing package "handler"
)

func main() {
	hh := handler.NewHello(log.New(os.Stdout, "Hello: ", log.LstdFlags))
	http.ListenAndServe(":8081", nil)
}
