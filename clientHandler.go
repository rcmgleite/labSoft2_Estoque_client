package main

import (
	"net/http"
	"text/template"
)

//defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/html/index.html")
	t.Execute(w, nil)
}
