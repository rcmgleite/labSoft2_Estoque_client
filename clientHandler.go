package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/soriani/labSoft2_Estoque/models"
)

func makeRequest(httpMethod string, url string, requestObj []byte, headers map[string]string) (*http.Response, error) {
	//creating request obj
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestObj))

	//Adding request headers
	addHeaders(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

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
	parseJSON(r.Body, &toDelete)

	//get JSON Marshaled as []byte
	bJSON, err := getJSON(toDelete)
	if err == nil {
		response, err := makeRequest("DELETE", "http://127.0.0.1:8080/product", bJSON, headers)

		if response != nil && err == nil {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println("response Body:", string(body))
			//TODO redirect
		}
	}
}
