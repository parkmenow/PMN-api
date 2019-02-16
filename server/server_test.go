package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/parkmenow/PMN-api/server"
	"github.com/parkmenow/PMN-api/models"
	"github.com/parkmenow/PMN-api/data"
)

// performRequest performs a http request and returns the response
func performRequest(r http.Handler, method, path string, reqbody string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer([]byte(reqbody)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Server", func() {

	// Setup initial pointer variable to be reused for each new test
	var (
		router   *gin.Engine
		response *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		router = CreateRouter()
	})

	Describe("Version 1 API at /api/v1", func() {
		Describe("The / endpoint or helloworld endpoint", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/", "")
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns the String 'Hello World'", func() {
				Expect(response.Body.String()).To(Equal("Hello World"))
			})
		})
	})

	Describe("POST the /signup endpoint", func() {
		BeforeEach(func() {
			response = performRequest(router, "POST", "/api/v1/signup", jsonSignUp)
		})
		It("Returns with Status 201", func() {
			Expect(response.Code).To(Equal(201))
		})
		It("Added a new User", func() {
			newuser := models.User{}
			json.Unmarshal([]byte(jsonSignUp), &newuser)
			Expect(data.Users().List[len(data.Users().List)-1]).To(Equal(newuser))
		})
	})

})

var jsonSignUp = `{
	"FName": "Subodh",
	"LName": "Pushkar",
	"UName": "subodh101"
	"Email": "subodh.pushkar@gmail.com"
	"PhoneNo": "9911991199"
	"Password": "password"
}`
