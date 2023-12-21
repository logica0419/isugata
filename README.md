# isugata

[![CI Pipeline](https://github.com/logica0419/isugata/actions/workflows/ci.yml/badge.svg)](https://github.com/logica0419/isugata/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/logica0419/isugata.svg)](https://pkg.go.dev/github.com/logica0419/isugata) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/logica0419/isugata/blob/main/LICENSE)

HTTP response validator for ISUCON benchmarker with Functional Option Pattern

## Usage

An executable example is in [example/main.go](./example/main.go).

- Status Code

```go
err := isugata.Validate(res,
  isugata.WithStatusCode(http.StatusOK),
)
```

- Content Type

```go
err := isugata.Validate(res,
  isugata.WithContentType("application/json"),
)
```

- JSON Body Validation

```go
type user struct {
 ID   int    `json:"id"`
 Name string `json:"name"`
}

err := isugata.Validate(res,
  isugata.WithJSONValidation[user](
    isugata.JSONEquals[user](
      user{
        ID:   1,
        Name: "test",
      },
    ),
  ),
)
```

- JSON Array Body Validation

```go
type user struct {
 ID   int    `json:"id"`
 Name string `json:"name"`
}

err := isugata.Validate(res,
  isugata.WithJSONArrayValidation[user](
    isugata.JSONArrayLengthEquals[user](2),
    isugata.JSONArrayValidateOrder[user, int](
      func(u user) int { return u.ID },
      isugata.Asc,
    ),
    isugata.JSONArrayValidateEach[user](
      func(body user) error {
        if body.Name != fmt.Sprintf("test%d", body.ID) {
          return fmt.Errorf("expected: test%d, actual: %s", body.ID, body.Name)
        }
        return nil
      },
    ),
  ),
)
```
