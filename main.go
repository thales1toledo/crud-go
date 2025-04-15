package main

import (
	"GoLand/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
