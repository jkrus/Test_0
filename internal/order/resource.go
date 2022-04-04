package order

import (
	"html/template"
	"log"
	"net/http"

	"wb_L0/pkg/models"
)

func OrderResponseTable(w http.ResponseWriter, order *models.Order) {
	tmpl, _ := template.ParseFiles("/workSpace/go/src/wb_L0/internal/order//html/order.page.tmpl")
	tmpl.Execute(w, order)
}

func Search(w http.ResponseWriter) {
	ts, err := template.ParseFiles("/workSpace/go/src/wb_L0/internal/order/html/search.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func BadSearch(w http.ResponseWriter, notfound string) {
	ts, err := template.ParseFiles("/workSpace/go/src/wb_L0/internal/order/html/badSearch.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, notfound)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
