package mocks

import "svc-customer/customer/entity"

// MockCustomer is a customer entity model with pointers
var MockCustomer = &entity.Customer{
	CustomerUUID: "039d69ee-f9cb-4a3d-87e4-6eb63c302579",
	Name:         "Eddwin",
	LastName:     "Paz",
	Dni:          "264573076",
	DniType:      "DNI",
	Email:        "ep@latamig.com",
	Phone:        "56900000000",
	Password:     "",
	Country:      "CL",
}

// MockCustomerPointless is a Customer entity model pointer less
var MockCustomerPointless = entity.Customer{
	Name:     "Bad",
	LastName: "Bunny",
	Dni:      "264573076",
	DniType:  "DNI",
	Phone:    "+130698569",
	Email:    "baby@gmail.com",
	Password: "conejomalo",
	Country:  "CL",
}

// CName constant value of customer Name
var CName = "Bad"

// CLastName constant value of customer LastName
var CLastName = "Bunny"

// CPhone constant value of customer Phone
var CPhone = "+130698569"

// CMail constant value of customer Mail
var CMail = "baby@gmail.com"

// CPassword constant value of customer Password
var CPassword = "conejomalo"

// CDni Dni number of customer's document
var CDni = "264573076"

// CDniType type of dni document
var CDniType = "DNI"

// InvalidJSONObject return JSON when user sends wrong JSON object in body
var InvalidJSONObject = "{\"status\":false,\"description\":\"Invalid JSON object\",\"data\":null}\n"

// RequestedCustomerNotFound return JSON when user cannot be found on database
var RequestedCustomerNotFound = "{\"status\":false,\"description\":\"requested is not found\",\"data\":null}\n"

// CustomerUUIDValid this is a valid customer UUID used to mock real customer UUID
var CustomerUUIDValid = "039d69ee-4a3d-87e4-6eb63c302579"
