package main

import (
	"fmt"
	"net/http"

	"github.com/soriani/labSoft2_Estoque/router"
)

//Run = main for client
func main() {
	r := router.NewRouter()
	r.AddRoute("/", router.GET, defaultHandler)
	r.AddRoute("/product", router.GET, GETProductHandler)
	r.AddRoute("/product", router.DELETE, DELETEProductHandler)

	http.Handle("/", r)

	fmt.Println("Client running on port: 8081")
	http.ListenAndServe(":8081", nil)
}
