package main

import (
	"net/http"
	"fmt"
)

func main(w http.ResponseWriter) {
	fmt.Fprintf("Ola, Mundo!")
}