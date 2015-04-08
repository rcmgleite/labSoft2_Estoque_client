package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

	"github.com/rcmgleite/labSoft2_Estoque/models"
	"github.com/rcmgleite/labSoft2_Estoque/requestHelper"
)

//responseJSON
type responseJSON struct {
	ResponseBody interface{}
	Error        string
}

func parseJSON(encodedJSON io.ReadCloser, object interface{}) {
	decoder := json.NewDecoder(encodedJSON)
	decoder.Decode(object)
}

// Return a json object marshaled as []byte
func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

//defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/html/index.html")
	t.Execute(w, nil)
}

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/html/product.html")
	// var products []models.Product
	var rj responseJSON

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	response, err := requestHelper.MakeRequest("GET", "http://127.0.0.1:8080/product", nil, headers)
	if response != nil && err == nil {
		parseJSON(response.Body, &rj)
		t.Execute(w, rj.ResponseBody)
		return
	}
	t.Execute(w, nil)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var toDelete models.Product
	//Parse json from request to object
	var id int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	toDelete.ID = id
	//get JSON Marshaled as []byte
	bJSON, err := getJSON(toDelete)
	if err == nil {
		response, err := requestHelper.MakeRequest("DELETE", "http://127.0.0.1:8080/product", bJSON, headers)

		if response != nil && err == nil {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println("response Body:", string(body))
			http.Redirect(w, r, "/product", http.StatusFound)

		}
	}
}

// POSTProductHandler ...
func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var newProduct models.Product
	newProduct.Name = r.FormValue("Name")
	newProduct.Description = r.FormValue("Description")
	newProduct.Type, _ = strconv.Atoi(r.FormValue("Type"))
	newProduct.CurrQuantity, _ = strconv.Atoi(r.FormValue("CurrQuantity"))
	newProduct.MinQuantity, _ = strconv.Atoi(r.FormValue("MinQuantity"))

	bJSON, err := getJSON(newProduct)
	if err == nil {
		_, err := requestHelper.MakeRequest("POST", "http://127.0.0.1:8080/product", bJSON, headers)
		if err != nil {
			fmt.Println(err)
		}
	}
	http.Redirect(w, r, "/product", http.StatusFound)

}

// GETProductUpdate ...
func GETProductUpdate(w http.ResponseWriter, r *http.Request) {
	var id int
	id, _ = strconv.Atoi(r.URL.Query().Get("id"))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	url := "http://127.0.0.1:8080/product?"

	url += "ID=" + strconv.Itoa(id)
	fmt.Println(url)
	response, _ := requestHelper.MakeRequest("GET", url, nil, headers)

	var rj responseJSON
	parseJSON(response.Body, &rj)

	t, _ := template.ParseFiles("views/html/productUpdate.html")
	t.Execute(w, rj.ResponseBody)
}

// POSTProductUpdate ...
func POSTProductUpdate(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var toUpdate models.Product

	toUpdate.ID, _ = strconv.Atoi(r.FormValue("ID"))
	toUpdate.Name = r.FormValue("Name")
	toUpdate.Description = r.FormValue("Description")
	toUpdate.Type, _ = strconv.Atoi(r.FormValue("Type"))
	toUpdate.CurrQuantity, _ = strconv.Atoi(r.FormValue("CurrQuantity"))
	toUpdate.MinQuantity, _ = strconv.Atoi(r.FormValue("MinQuantity"))

	bJSON, err := getJSON(toUpdate)
	if err == nil {
		_, err := requestHelper.MakeRequest("PUT", "http://127.0.0.1:8080/product", bJSON, headers)
		if err != nil {
			fmt.Println(err)
		}
	}
	http.Redirect(w, r, "/product", http.StatusFound)
}

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/html/order.html")
	// var products []models.Product
	var rj responseJSON

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	response, err := requestHelper.MakeRequest("GET", "http://127.0.0.1:8080/order", nil, headers)
	if response != nil && err == nil {
		parseJSON(response.Body, &rj)
		t.Execute(w, rj.ResponseBody)
		return
	}
	t.Execute(w, nil)
}

// POSTOrderHandler ...
func POSTOrderHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	id, _ = strconv.Atoi(r.FormValue("ID"))
	order := models.Order{ID: id, Approved: true}
	bJSON, err := getJSON(order)
	if err == nil {
		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"
		_, err := requestHelper.MakeRequest("PUT", "http://127.0.0.1:8080/order", bJSON, headers)
		if err != nil {
			fmt.Println(err)
		}
	}

	http.Redirect(w, r, "/order", http.StatusFound)
}
