package main

import (
	"fmt"
	"log"
	"net/http"
)

func portfolioSever(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the new world code 4Chainz")
}

func main() {
	http.HandleFunc("/", portfolioSever)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
