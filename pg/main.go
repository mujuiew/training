package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "postgres"
	dbname    = "payment"
	layoutISO = "2006-01-02"
	layoutUS  = "2006-01-02"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// Rate ...
type Rate struct {
	Rate          string  `json:"rate"`
	InterestRate  float64 `json:"interest_rate"`
	PromotionName string  `json:"promotion_name"`
}

// Promotion ...
type Promotion struct {
	PromotionName string `json:"promotion_name"`
}

var pro Promotion
var rate Rate

// Input ...
type Input struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    float64 `json:"number_of_payment"`
	CalDate            string  `json:"cal_date"`
	PaymentFrequency   float64 `json:"payment_frequency"`
	PaymentUnit        string  `json:"payment_unit"`
	AccountNumber      float64 `json:"account_number"`
}

// Input2 ...
type Input2 struct {
	Req Input `json:"rqbody"`
}

func repoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", psqlInfo)
	rows, err := db.Query(
		`SELECT rate, interest_rate, promotion_name FROM "Rate"`)
	for rows.Next() {

		var rate Rate
		if err := rows.Scan(&rate.Rate, &rate.InterestRate, &rate.PromotionName); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, rate)
	}
	if err != nil {
		panic(err)
	}

}
func personCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in Input
	_ = json.NewDecoder(r.Body).Decode(&in)

	date := in.CalDate
	t, _ := time.Parse(layoutISO, date)
	caldate := t.Format(layoutUS)

	proname := fineProname(caldate)

	inrate := fineinrate(proname)
	pmt := insertAc(in.AccountNumber, inrate, in.DisbursementAmount, in.NumberOfPayment)

	fmt.Fprintf(w, "promotion_name = %v\n", proname)
	fmt.Fprintf(w, "interest_rate  = %.2f\n", inrate)
	fmt.Fprintf(w, "account_number = %.2f\n", in.NumberOfPayment)
	fmt.Fprintf(w, "Payment Amount with Nember of payment %v = %.2f\n", in.NumberOfPayment, pmt)
}

func fineProname(date string) string {
	db, _ := sql.Open("postgres", psqlInfo)
	rows, _ := db.Query(`SELECT promotion_name FROM "Promotion" WHERE '` + date + `'  between start_date and end_date `)
	for rows.Next() {
		if err := rows.Scan(&pro.PromotionName); err != nil {
			log.Fatal(err)
		}
	}
	proname := pro.PromotionName

	return proname
}
func fineinrate(pro string) float64 {
	db, _ := sql.Open("postgres", psqlInfo)

	rows, _ := db.Query(`SELECT interest_rate FROM "Rate" WHERE promotion_name = '` + pro + `'`)
	for rows.Next() {
		if err := rows.Scan(&rate.InterestRate); err != nil {
			log.Fatal(err)
		}
	}
	inrate := rate.InterestRate
	return inrate
}
func insertAc(acc float64, i float64, dis float64, num float64) float64 {
	i = i / 100 / 12
	pmt := dis / ((1 - (1 / math.Pow((1+i), num))) / i)

	db, err := sql.Open("postgres", psqlInfo)
	err = db.Ping()
	queryStr := `INSERT INTO public."Account"(
		account_number, installment_amount)
		VALUES ($1, $2)`
	_, err = db.Exec(queryStr, acc, pmt)
	if err != nil {
		panic(err)
	}
	return pmt
}
func calPmt1(dis float64, i float64, num float64, message chan<- float64) {
	i = i / 100 / 12
	pmt := dis / ((1 - (1 / math.Pow((1+i), num))) / i)
	fmt.Printf("Payment Amount with Nember of payment %v = %.2f\n", num, pmt)
	message <- pmt
}

func main() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	r := mux.NewRouter()
	r.HandleFunc("/dloan-payment/v1/calculate-installment-amount", repoHandler).Methods("GET")
	r.HandleFunc("/person/create", personCreate).Methods("POST")

	http.ListenAndServe(":8080", r)

}
