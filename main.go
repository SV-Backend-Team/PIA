package main

import (
	"fmt"
	"net/http"
	config "northwindApi/config"
	employees "northwindApi/tables/employee"

	"github.com/gorilla/mux"
)

func main() {
	config.DB_export = config.DB{
		Server:   "localhost",
		Port:     1433,
		User:     "",
		Password: "",
		Database: "Northwind",
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/employee/findall", employees.FindAll).Methods("GET")
	router.HandleFunc("/api/employee/search/{id}", employees.FindByID).Methods("GET")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}