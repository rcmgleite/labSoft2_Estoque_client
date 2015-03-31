package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/rcmgleite/labEngSoft_Estoque/models"
)

//defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	p.ID = 3
	bJSON, err := json.Marshal(p)
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8080/product", bytes.NewBuffer(bJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	t, _ := template.ParseFiles("views/html/index.html")
	t.Execute(w, nil)
}
