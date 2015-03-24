package main

import (
	"fmt"
	"net/http"

	"github.com/rcmgleite/labEngSoft_Estoque/router"
)

//Run = main for client
func main() {
	r := router.NewRouter()
	r.AddRoute("/", router.GET, defaultHandler)
	http.Handle("/", r)

	fmt.Println("Client running on port: 8081")
	http.ListenAndServe(":8081", nil)
}
