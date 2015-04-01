package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/rcmgleite/labEngSoft_Estoque/models"
)

func makeRequest(httpMethod string, url string, requestObj []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestObj))
	addHeaders(req, headers)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func parseJSON(decodedJSON io.ReadCloser, object interface{}) {
	decoder := json.NewDecoder(decodedJSON)
	decoder.Decode(object)
}

func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

//defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/html/index.html")
	t.Execute(w, nil)
}

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/html/product.html")
	var products []models.Product

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	response, err := makeRequest("GET", "http://127.0.0.1:8080/product", nil, headers)
	if response != nil && err == nil {
		parseJSON(response.Body, &products)

		t.Execute(w, products)
	}
	t.Execute(w, nil)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var toDelete models.Product
	parseJSON(r.Body, &toDelete)

	bJSON, err := getJSON(toDelete)
	response, err := makeRequest("DELETE", "http://127.0.0.1:8080/product", bJSON, headers)

	if response != nil && err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(body))
		//TODO redirect
	}

}
