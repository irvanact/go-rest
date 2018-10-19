package main

import (
	"fmt"
	. "go-rest-sample/config"
	"go-rest-sample/controller"
	. "go-rest-sample/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var accountRepository = AccountRepository{}
var config = Config{}

func init() {

	config.Read()
	accountRepository.Server = config.Server
	accountRepository.Database = config.Database
	accountRepository.Connect()
	fmt.Println("init " + accountRepository.Database)
}

func main() {
	fmt.Println("test")
	r := mux.NewRouter()
	r.HandleFunc("/accounts", controller.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", controller.AddAccount).Methods("POST")
	r.HandleFunc("/accounts", controller.UpdateAccount).Methods("PUT")
	r.HandleFunc("/accounts", controller.DeleteAccount).Methods("DELETE")
	r.HandleFunc("/accounts/{id}", controller.GetAccount).Methods("GET")
	if err := http.ListenAndServe(":12000", r); err != nil {
		log.Fatal(err)
	}
}
