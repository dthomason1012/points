package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
)

type Transaction struct {
	Payer 		string	`json:"Payer"`
	Points 		int		`json:"Points"`
	Timestamp 	string	`json:"Timestamp"`
}

type Payer struct {
	Payer string `json:"Payer"`
	Points int `json:"Points"`
}

var Transactions []Transaction
var Payers []Payer

func main() {
	fmt.Printf("Starting server at port 8080\n")

	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the bare minimum.")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/transactions", returnAllTransactions)
	http.HandleFunc("/add", addTransaction)
	http.HandleFunc("/balances", returnBalances)
	http.HandleFunc("/spend", spend)
	// http.HandleFunc("/balances", balanceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnAllTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Transactions)
}

func returnBalances(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Payers)
}

func addTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	var payer Payer
	
	transaction.Payer = r.URL.Query().Get("payer")
	transaction.Points, _ = strconv.Atoi(r.URL.Query().Get("points"))

	for i := range Payers {
		if Payers[i].Payer == transaction.Payer {
			Payers[i].Points += transaction.Points
			http.Redirect(w, r, "/balances", http.StatusFound)
			return
		}
	}

	payer = Payer{transaction.Payer, transaction.Points}
	Payers = append(Payers, payer)
	
	http.Redirect(w, r, "/balances", http.StatusFound)
}

func spend(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/balances", http.StatusFound)
}
