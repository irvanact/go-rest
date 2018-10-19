package controller

import (
	"encoding/json"
	. "go-rest-sample/model"
	. "go-rest-sample/repository"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var accountRepository = AccountRepository{}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := accountRepository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, accounts)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account, err := accountRepository.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
		return
	}
	respondWithJson(w, http.StatusOK, account)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := accountRepository.Update(account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func AddAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	account.ID = bson.NewObjectId()
	if err := accountRepository.Insert(account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := accountRepository.Delete(account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
