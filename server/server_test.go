package server_test

import (
	"bytes"
	"fmt"
	//"github.com/parkmenow/PMN-api/models"
	"io"

	//"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/parkmenow/PMN-api/server"
)


// performRequest performs a http request and returns the response
func performRequest(r http.Handler, method, path string, body io.Reader, ch chan *httptest.ResponseRecorder  )  {
	req, _ := http.NewRequest(method, path, body)

	// Create a Bearer string by appending string access token, adding a bearer token for checking the authentication
	// We have created a token for user: test3, pass: test3, token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA3Mzg1OTAsImlkIjozLCJvcmlnX2lhdCI6MTU1MDczODU5MH0.AwyBaGE31Yq2dURoP7uIe91zIQwHlTkkYW6a2kLoVa8
	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA3Mzg1OTAsImlkIjozLCJvcmlnX2lhdCI6MTU1MDczODU5MH0.AwyBaGE31Yq2dURoP7uIe91zIQwHlTkkYW6a2kLoVa8"
	req.Header.Add("Authorization", bearer)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	ch <- w

}

var _ = Describe("Server", func() {

	var (
		db       *gorm.DB
		router   *gin.Engine
		response *httptest.ResponseRecorder

	)
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	driver := os.Getenv("DB_DRIVER")
	url := os.Getenv("DATABASE_URL")

	BeforeEach(func() {

		d, err := gorm.Open(driver, url)
		db = d
		if err != nil {
			panic(err)
		}

		router = CreateRouter(db)


	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Test cases for server API end points, / ", func() {
		Describe("The / endpoint or helloworld", func() {
			BeforeEach(func() {
				var body=new(bytes.Buffer)
				ch1 := make(chan *httptest.ResponseRecorder)
				go performRequest(router, "GET", "/",body, ch1)
				response =  <- ch1

			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns the String 'Hello World'", func() {
				Expect(response.Body.String()).To(Equal("Hello World"))
			})
		})
fmt.Println("-----First--------")
		// TODO: Needs to fix this.
		/*
		Describe("API /login", func() {
			BeforeEach(func() {
				    body := new(bytes.Buffer)
					body.Write([]byte(`{"U_Name": "test3", "Password": "test3"}`))
					response = performRequest(router, "POST", "/login", body)
					fmt.Println((response.Body.String()))
				//bytes.NewBuffer([]byte(body))
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

		})
		*/


		Describe("API /dashboard/:id/", func() {
			BeforeEach(func() {
				body := new(bytes.Buffer)
				ch1 := make(chan *httptest.ResponseRecorder)
				go performRequest(router, "GET", "/dashboard/1/", body, ch1)
				response = <- ch1
				fmt.Println((response.Body.String()))
				//bytes.NewBuffer([]byte(body))
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns with Output 'Bharath'", func() {
				Expect(response.Body.String()).To(Equal("\"Bharath\""))
			})

		})
		fmt.Println("-----second--------")
		Describe("API /dashboard/:id/mylistings", func() {
			BeforeEach(func() {
				body := new(bytes.Buffer)
				ch1 := make(chan *httptest.ResponseRecorder)
				go performRequest(router, "GET", "/dashboard/1/mylistings", body, ch1)
				response = <- ch1
				fmt.Println((response.Body.String()))
				//bytes.NewBuffer([]byte(body))
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

		})
		fmt.Println("-----third--------")
		Describe("API /dashboard/:id/parkmenow", func() {
			BeforeEach(func() {
				body := new(bytes.Buffer)
				body.Write([]byte(`{
									"Type" : 1,
    								"Lat": 35.660736,
    								"Long": 139.72955,
    								"StartTime" : "2006-01-02T14:00:00.000Z",
									"EndTime" : "2006-01-02T15:00:00.000Z"
									}`))
				ch1 := make(chan *httptest.ResponseRecorder)
				go performRequest(router, "POST", "/dashboard/1/parkmenow", body, ch1)
				response = <- ch1
				fmt.Println((response.Body.String()))
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

		})
		fmt.Println("-----fourth--------")


		Describe("API /dashboard/:id/regparking", func() {
			BeforeEach(func() {
				body := new(bytes.Buffer)
				body.Write([]byte(`{
       							 "Line1" : "1-5-6, 1108",
       							 "Line2" : "Higashi-ojima",
        							"Pincode": "132-0034",
       							 "lat": 35.68981,
       							 "long": 139.84755,
       							 "OwnerID": 3
    								}
								`))
				var ch1 = make(chan *httptest.ResponseRecorder)
				go performRequest(router, "POST", "/dashboard/1/regparking", body,ch1)
				response = <- ch1
				fmt.Println((response.Body.String()))
			})

			It("Returns with Status 201", func() {
				Expect(response.Code).To(Equal(201))
			})

			It("Returns saying succesful booking", func() {
				Expect(response.Body.String()).To(Equal("\"Listed a new parking Spot Successfully!\""))
			})

		})


	})


})

// sign up user testing
var jsonSignUp = `{
	"FName": "ssssss",
	"LName": "pppppp",
	"UName": "subodh101",
	"Password": "password",
	"Email": "subodh.pushkar@gmail.com",
	"PhoneNo": "9911991199",
	"Line1" : "1-5-5, 1108",
	"Line2" : "Higashi-ojima",
	"Pincode": "132-0034"
}`
