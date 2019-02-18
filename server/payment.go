package server

import (
	"log"
	"os"
	
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func paymentHandler(p int64, email string, token string) bool {

	//export SecretKey="sk_test_1pSlxntEQATjsOv5HLI49FaW"
	var sh_key = os.Getenv("SecretKey")
	stripe.Key = sh_key

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(p),
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		ReceiptEmail:stripe.String(email),
	}

	//Add the token here
	params.SetSource("tok_mastercard")
	//params.SetSource(input.Token)

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.ID)
	return ch.Paid
}
