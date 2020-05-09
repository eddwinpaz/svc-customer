package mocks

import "svc-customer/customer/entity"

// MockQuerySelectByUUID Test Repository SQL Mocks to select valid customer
var MockQuerySelectByUUID = "SELECT name, last_name, dni, dni_type, email, phone, customer_uuid, country FROM customers WHERE customer_uuid = \\? AND deleted_at IS NULL LIMIT 1"

// MockColumnsSelectByUUID Test Repository Column Mocks for select customer
var MockColumnsSelectByUUID = []string{"name", "last_name", "dni", "dni_type", "email", "phone", "customer_uuid", "country"}

// MockQueryStore is a Customer entity model pointer less
var MockQueryStore = entity.Customer{
	Name:     "Jhon",
	LastName: "Doe",
	Dni:      "264573076",
	DniType:  "DNI",
	Phone:    "+56933375029",
	Email:    "jhond@gmail.com",
	Password: "goodpassword",
	Country:  "US",
}

// MockColumnsInsertStore columns
var MockColumnsInsertStore = []string{"id", "customer_uuid", "name", "last_name", "dni", "dni_type", "email", "phone", "customer_uuid", "country", "password", "created_at", "updated_at", "deleted_at"}

// MockQueryInsertStore Test Repository SQL Mock to store valid customer
var MockQueryInsertStore = "INSERT INTO customers"

// MockQueryUpdateCustomer update customer changes
var MockQueryUpdateCustomer = "UPDATE customers"

// MockSQLQueryDelete Delete Test Repository SQL Mock to Delete valid customer
var MockSQLQueryDelete = "UPDATE customers SET deleted_at = CURRENT_TIMESTAMP WHERE customer_uuid = \\? AND deleted_at IS NULL LIMIT 1"

// MockUpdateCustomerSQL Update Test Repository SQL mock to Update valid customer data
var MockUpdateCustomerSQL = "UPDATE customers SET name = \\? , last_name = \\? ,updated_at = CURRENT_TIMESTAMP WHERE customer_uuid = \\? AND deleted_at IS NULL LIMIT 1"

// MockFetchCustomerSQL Fetch All Test Repository SQL mock to Fetch all valid customers
var MockFetchCustomerSQL = "SELECT name, last_name, dni, dni_type, email, phone, customer_uuid, country FROM customers WHERE deleted_at IS NULL ORDER BY id DESC LIMIT ? OFFSET ?"
