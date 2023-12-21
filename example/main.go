//nolint:all
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/logica0419/isugata"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	body := `
		[
			{
				"id": 1,
				"name": "test1"
			},
			{
				"id": 2,
				"name": "test2"
			}
		]
	`

	res := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}

	err := isugata.Validate(res,
		// status code validation
		isugata.WithStatusCode(http.StatusOK),
		// content type validation
		isugata.WithContentType("application/json"),
		// JSON array body validation
		isugata.WithJSONArrayValidation(
			// length validation
			isugata.JSONArrayLengthEquals[user](2),
			// order validation
			isugata.JSONArrayValidateOrder(
				func(u user) int { return u.ID },
				isugata.Asc,
			),
			// element validation
			isugata.JSONArrayValidateEach(
				// validation with original ValidateOpt
				func(body user) error {
					if body.Name != fmt.Sprintf("test%d", body.ID) {
						return fmt.Errorf("expected: test%d, actual: %s", body.ID, body.Name)
					}
					return nil
				},
			),
		),
	)
	if err != nil {
		panic(err)
	}
}
