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
	fmt.Println("Rodando API na porta", config.Port)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3001", r))
}
