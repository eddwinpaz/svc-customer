// Code generated by go-swagger; DO NOT EDIT.

// Package restapi Muble Customer's API
//
//  This is a simple todo list API
//  illustrating go-swagger codegen
//  capabilities.
//
//  Terms Of Service:
//    There are actually none.
//    This is for demonstration purpose only.
//
//  Schemes:
//    http
//    https
//  Host: api.muble.app
//  BasePath: /
//  Version: 0.1.0
//  License: Apache 2.0 https://www.apache.org/licenses/LICENSE-2.0
//  Contact: muble maintainers<api@muble.app>
//
//  Consumes:
//    - application/json
//
//  Produces:
//    - application/json
//
// swagger:meta
package restapi

// swagger:response swaggerResponse
type swaggerResponse struct {
	// in: body
	Body struct {
		Status      bool    `json:"status"`
		Description string  `json:"description"`
		Data        Swagger `json:"data"`
	}
}

// swagger:parameters UpdateByUUID
type Swagger struct {
	// name
	// Min Length: 1
	// Max Length: 30

	Name string `json:"name,omitempty"`
	// last_name
	// Min Length: 1
	// Max Length: 30

	LastName string `json:"last_name,omitempty"`
	// Dni
	// Min Length: 1
	// Max Length: 20

	Dni string `json:"dni,omitempty"`
	// DniType
	// Min Length: 1
	// Max Length: 10

	DniType string `json:"dni_type,omitempty"`
	// Phone
	// Min Length: 1
	// Max Length: 20

	Phone string `json:"phone,omitempty"`
	// Country
	// Min Length: 1

	Country string `json:"country,omitempty"`
	// email
	// Min Length: 2
	// Max Length: 2

	Email string `json:"email,omitempty"`
	// password
	// Min Length: 1
	// Max Length: 15
}

// swagger:response swaggerResponseFail
type swaggerResponseFail struct {
	// in: body
	Body struct {
		Status      bool        `json:"status"`
		Description string      `json:"description"`
		Data        interface{} `json:"data"`
	}
}

// swagger:parameters Store
type SwaggerStore struct {
	// name
	// Required: true
	// Min Length: 1
	// Max Length: 30

	Name string `json:"name,omitempty"`
	// last_name
	// Required: true
	// Min Length: 1
	// Max Length: 30

	LastName string `json:"last_name,omitempty"`
	// Dni
	// Required: true
	// Min Length: 1
	// Max Length: 20

	Dni string `json:"dni,omitempty"`
	// DniType
	// Required: true
	// Min Length: 1
	// Max Length: 10

	DniType string `json:"dni_type,omitempty"`
	// Phone
	// Required: true
	// Min Length: 1
	// Max Length: 20

	Phone string `json:"phone,omitempty"`
	// Country
	// Required: true
	// Min Length: 1

	Country string `json:"country,omitempty"`
	// email
	// Required: true
	// Min Length: 2
	// Max Length: 2

	Email string `json:"email,omitempty"`
	// password
	// Required: true
	// Min Length: 1
	// Max Length: 15

	Password string `json:"password,omitempty"`
}

// swagger:response swaggerResponseArray
type swaggerResponseArray struct {
	// in: body
	Body struct {
		Status      bool      `json:"status"`
		Description string    `json:"description"`
		Data        []Swagger `json:"data"`
	}
}
