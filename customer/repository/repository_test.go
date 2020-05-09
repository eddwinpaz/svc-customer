package repository_test

import (
	"svc-customer/customer/entity"
	"svc-customer/customer/mocks"
	"svc-customer/customer/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// Test sending proper data.
func TestRepositoryGetByUUIDValid(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	rows := sqlmock.NewRows(mocks.MockColumnsSelectByUUID).
		AddRow("Jhon", "Doe", "264573076", "DNI", "jdoe@gmail.com", "5600000000", "f48ac180-e8ad-4837-a3c3-66b0e96f19bf", "CL")

	mock.ExpectQuery(mocks.MockQuerySelectByUUID).WillReturnRows(rows)
	h := repository.NewMySQLCustomersRepository(db)

	customerUUID := "f48ac180-e8ad-4837-a3c3-66b0e96f19bf"
	customer, err := h.GetByUUID(customerUUID)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

// Test missing UUID
func TestRepositoryGetByUUIDInvalid(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close() //nolint

	rows := sqlmock.NewRows(mocks.MockColumnsSelectByUUID).
		AddRow("Mike", "Tyson", "264573076", "DNI", "miketyson@gmail.com", "5600000000", "f48ac180-e8ad-4837-a3c3-66b0e96f19bf", "CL")

	mock.ExpectQuery(mocks.MockQuerySelectByUUID).WillReturnRows(rows)
	h := repository.NewMySQLCustomersRepository(db)

	customerUUID := "" //"f48ac180-e8ad-4837-a3c3-66b0e96f19bf"
	customer, err := h.GetByUUID(customerUUID)
	assert.EqualError(t, entity.ErrBadParamInput, err.Error())
	assert.Nil(t, customer)
}

// Test Store Valid
func TestRepositoryStoreValid(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	mockCustomer := entity.Customer{
		Name:     "Jhon",
		LastName: "Doe",
		Dni:      "264573076",
		DniType:  "DNI",
		Phone:    "+56933375029",
		Email:    "jhond@gmail.com",
		Password: "goodPassword",
		Country:  "CL",
	}

	emailRow := sqlmock.NewRows([]string{"exists"})
	mock.ExpectQuery("SELECT exists \\(SELECT id FROM customers WHERE email = \\?\\)").WillReturnRows(emailRow)

	phoneRow := sqlmock.NewRows([]string{"exists"})
	mock.ExpectQuery("SELECT exists \\(SELECT id FROM customers WHERE phone = \\?\\)").WillReturnRows(phoneRow)

	SQLQuery := "INSERT INTO customers"

	mock.
		ExpectExec(SQLQuery).
		WithArgs(
			mockCustomer.Name,
			mockCustomer.LastName,
			mockCustomer.Dni,
			mockCustomer.DniType,
			mockCustomer.Email,
			mockCustomer.Phone,
			mockCustomer.Country,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h := repository.NewMySQLCustomersRepository(db)

	err = h.Store(mockCustomer)

	assert.Equal(t, nil, err)
}

func TestRepositoryDeleteByUUIDValid(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	mock.
		ExpectExec(mocks.MockSQLQueryDelete).
		WithArgs(mocks.CustomerUUIDValid).
		WillReturnResult(sqlmock.NewResult(0, 1))

	h := repository.NewMySQLCustomersRepository(db)

	err = h.DeleteByUUID(mocks.CustomerUUIDValid)

	assert.Equal(t, nil, err)
}

func TestRepositoryDeleteByUUIDSQLError(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	mock.
		ExpectExec(mocks.MockSQLQueryDelete).
		WithArgs("").
		WillReturnResult(sqlmock.NewResult(0, 0))

	h := repository.NewMySQLCustomersRepository(db)

	err = h.DeleteByUUID(mocks.CustomerUUIDValid)

	assert.Equal(t, entity.ErrSQLError, err)
}

func TestRepositoryDeleteByUUIDNotFound(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	// customerUUID := "f48ac180-e8ad-4837-a3c3-66b0e96f19bf"

	mock.
		ExpectExec(mocks.MockSQLQueryDelete).
		WithArgs("").
		WillReturnResult(sqlmock.NewResult(0, 0))

	h := repository.NewMySQLCustomersRepository(db)

	err = h.DeleteByUUID("")

	assert.Equal(t, entity.ErrNotFound, err)
}

func TestRepositoryUpdateByUUIDValid(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	mockCustomer := entity.Customer{
		Name:     mocks.CName,
		LastName: mocks.CLastName,
	}

	emailRow := sqlmock.NewRows([]string{"exists"})
	mock.ExpectQuery("SELECT exists \\(SELECT id FROM customers WHERE email = \\? AND customer_uuid != \\?\\)").WillReturnRows(emailRow)

	phoneRow := sqlmock.NewRows([]string{"exists"})
	mock.ExpectQuery("SELECT exists \\(SELECT id FROM customers WHERE phone = \\? AND customer_uuid != \\?\\)").WillReturnRows(phoneRow)

	mock.
		ExpectExec(mocks.MockUpdateCustomerSQL).
		WithArgs(mocks.CName, mocks.CLastName, mocks.CustomerUUIDValid).
		WillReturnResult(sqlmock.NewResult(0, 1))

	h := repository.NewMySQLCustomersRepository(db)

	err = h.UpdateByUUID(mockCustomer, mocks.CustomerUUIDValid)

	assert.Equal(t, nil, err)
}

func TestRepositoryUpdateByUUIDNoUUID(t *testing.T) {
	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	mockCustomer := entity.Customer{
		Name:     mocks.CName,
		LastName: mocks.CLastName,
	}

	emailRow := sqlmock.NewRows([]string{"exists"})
	mock.ExpectQuery("SELECT exists \\(SELECT id FROM customers WHERE email = \\?\\)").WillReturnRows(emailRow)

	phoneRow := sqlmock.NewRows([]string{"exists"})
	mock.ExpectQuery("SELECT exists \\(SELECT id FROM customers WHERE phone = \\?\\)").WillReturnRows(phoneRow)

	mock.
		ExpectExec(mocks.MockUpdateCustomerSQL).
		WithArgs(mocks.CName, mocks.CLastName, mocks.CustomerUUIDValid).
		WillReturnResult(sqlmock.NewResult(0, 1))

	h := repository.NewMySQLCustomersRepository(db)

	err = h.UpdateByUUID(mockCustomer, "")

	assert.Equal(t, entity.ErrSQLError, err)
}

func TestOpenConnection(t *testing.T) {
	conn := repository.OpenConnection()
	_, err := conn.Driver().Open("foo")

	assert.Error(t, err)
}

func TestRepositoryFetchValid(t *testing.T) {

	page, limit := 1, 10

	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	row := sqlmock.NewRows([]string{"total"})
	mock.ExpectQuery("SELECT COUNT\\(id\\) AS total FROM customers").WillReturnRows(row)

	rows := sqlmock.NewRows(
		[]string{"name", "last_name", "dni", "dni_type", "email", "phone", "customer_uuid", "country"}).
		AddRow("Jhon", "Doe", "264573076", "DNI", "jdoe@gmail.com", "5600000000", "f48ac180-e8ad-4837-a3c3-66b0e96f19bf", "CL")

	mock.ExpectQuery("SELECT name, last_name, dni, dni_type, email, phone, customer_uuid, country FROM customers WHERE deleted_at IS NULL ORDER BY id DESC LIMIT \\? OFFSET \\?").
		WillReturnRows(rows).
		WithArgs(10, 0)

	h := repository.NewMySQLCustomersRepository(db)
	// // begin, pages, total, err := repository.Pagination(page, limit)
	customer, _, _, _, err := h.Fetch(page, limit)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

func TestRepositoryFetchInvalid(t *testing.T) {

	// Open Database Connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	mock.ExpectQuery(mocks.MockFetchCustomerSQL).WillReturnError(entity.ErrSQLError)
	h := repository.NewMySQLCustomersRepository(db)

	_, _, _, _, err = h.Fetch(1, 10)
	assert.Error(t, entity.ErrSQLError, err.Error())
}

func TestRepositoryHealthCheckValid(t *testing.T) {
	db, _, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	defer db.Close() //nolint

	h := repository.NewMySQLCustomersRepository(db)
	check := h.HealthCheck()

	assert.Nil(t, check)

}

func TestRepositoryHealthCheckError(t *testing.T) {
	db, _, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Migrate the schema
	db.Close() //nolint

	h := repository.NewMySQLCustomersRepository(db)
	check := h.HealthCheck()

	assert.Equal(t, "sql: database is closed", check.Error())
}
