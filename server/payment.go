package server

import (
	"log"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func paymentHandler(price int64, email string, token string) (bool, string) {

	//export SecretKey="sk_test_1pSlxntEQATjsOv5HLI49FaW"
	var sh_key = os.Getenv("SecretKey")
	stripe.Key = sh_key

	params := &stripe.ChargeParams{
		Amount:       stripe.Int64(price),
		Currency:     stripe.String(string(stripe.CurrencyJPY)),
		ReceiptEmail: stripe.String(email),
	}
	params.SetSource(token)

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.ID)
	return ch.Paid,ch.FailureMessage
}
