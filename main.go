package main

import (
	"fmt"
	"github.com/gorilla/mux"
	. "golang-api/conf"
	. "golang-api/conf/DAO"
	route "golang-api/router"
	"log"
	"net/http"
)

var dao = ArticlesDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}


func main() {
	fmt.Print("Olá, Api")
	router := mux.NewRouter()
	router.HandleFunc("/contato", route.GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
