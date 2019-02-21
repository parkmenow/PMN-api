package server_test

import (
	"bytes"
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
func performRequest(r http.Handler, method, path string, reqbody string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer([]byte(reqbody)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
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

	Describe("Version 1 API at /api/v1", func() {
		Describe("The / endpoint or helloworld endpoint", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/", "")
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns the String 'Hello World'", func() {
				Expect(response.Body.String()).To(Equal("Hello World"))
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
