package entity

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")

	// ErrRecordNotFound returned
	ErrRecordNotFound = errors.New("record not found")

	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("requested is not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("already exist")

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")

	// ErrorNotFoundOnDB will throw if the given answer is not present on database
	ErrorNotFoundOnDB = errors.New("Element not found on DB")

	// ErrorInvalidForm will throw if the given user input is invalid for some
	// kind when present on delivery/web validation
	ErrorInvalidForm = errors.New("Sent Form is not valid")

	// ErrorAlreadyExists will throw if the new customer email is in our database
	ErrorAlreadyExists = errors.New("Email or Phone already exists")

	// ErrEmailExists email already exists error
	ErrEmailExists = errors.New("Email already exists")

	// ErrPhoneExists phone already exists error
	ErrPhoneExists = errors.New("Phone already exists")

	// ErrDatabaseError this happens when SQL fails
	ErrDatabaseError = errors.New("Database Error, Try again")

	// ErrSQLError Internal Database Error
	ErrSQLError = errors.New("Internal Database Error, Try Later")

	//
)

/* SQL-QUERY-FETCH-CUSTOMER(FASTER)
This is the fastest way but only works ascending.
var SQLQuery = `
SELECT name, last_name, dni, dni_type, email, phone, customer_uuid, country
FROM customers
WHERE deleted_at IS NULL
AND id > ?
ORDER BY id ASC
LIMIT ? `
*/
