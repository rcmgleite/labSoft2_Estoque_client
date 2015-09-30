package main

import (
	"fmt"
	"net/http"

	"github.com/rcmgleite/router"
)

//Run = main for client
func main() {
	r := router.NewRouter()
	r.AddRoute("/", router.GET, defaultHandler)
	r.AddRoute("/product", router.GET, GETProductHandler)
	r.AddRoute("/productDelete", router.GET, DELETEProductHandler)
	r.AddRoute("/productAdd", router.POST, POSTProductHandler)
	r.AddRoute("/productUpdate", router.GET, GETProductUpdate)
	r.AddRoute("/productUpdate", router.POST, POSTProductUpdate)
	r.AddRoute("/order", router.GET, GETOrderHandler)
	r.AddRoute("/order", router.POST, POSTOrderHandler)

	fmt.Println("Client running on port: 8081")
	http.ListenAndServe(":8081", r)
}
