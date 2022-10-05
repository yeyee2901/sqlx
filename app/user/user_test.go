package user_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEndpoint struct {
	Name         string
	TestFunction func(t *testing.T)
}

func TestGetUser(t *testing.T) {
	client := http.DefaultClient
	path := "http://localhost:8767/user"

	testTable := []TestEndpoint{
		{
			// CASE #1:
			// get 1 user terdaftar
			// assert http status = 200
			Name: "Exist_HttpStatus",
			TestFunction: func(t *testing.T) {
				req, _ := http.NewRequest("GET", path, nil)

                // add query string
                query := req.URL.Query()
                query.Add("id", "4")
                req.URL.RawQuery = query.Encode()

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					t.FailNow()
				}

				assert.Equal(t, http.StatusOK, resp.StatusCode)
			},
		},
		{
			// CASE #2:
			// get 1 user, tidak terdaftar
			// assert http status = 404
			Name: "NotExist_HttpStatus",
			TestFunction: func(t *testing.T) {
				req, _ := http.NewRequest("GET", path, nil)

                // add query string
                query := req.URL.Query()
                query.Add("id", "100000")
                req.URL.RawQuery = query.Encode()

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					t.FailNow()
				}

				assert.Equal(
					t,
					http.StatusNotFound,
					resp.StatusCode,
				)
			},
		},
	}

	// run all get user tests
	for _, test := range testTable {
		t.Run(test.Name, test.TestFunction)
	}
}
