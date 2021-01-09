package main

import (
	"log"
	"os"

	lunchmoney "github.com/amaraliou/lunch-money-go"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	client, err := lunchmoney.NewClient(os.Getenv("LUNCH_MONEY_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	opts := &lunchmoney.GetTransactionsOptions{
		StartDate: "2020-06-01",
		EndDate:   "2021-01-01",
	}

	resp, err := client.GetTransactions(opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(resp)
}
