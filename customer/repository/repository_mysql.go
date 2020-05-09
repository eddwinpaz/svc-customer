package repository

import (
	"customer/customer/entity"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for repository
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// MySQLCustomersRepository structure for database connection
type MySQLCustomersRepository struct {
	db *sql.DB
}

// DATABASEDRIVER enables you to select your DBAS
const DATABASEDRIVER = "mysql"

// OpenConnection between server and MySQL(driver) Server
func OpenConnection() *sql.DB {

	// get env variables
	mysqlHost, errHost := os.LookupEnv("MYSQL_HOST")
	if !errHost {
		log.Warnf("MYSQL_HOST: got %s expected value", mysqlHost)
	}

	// mysqlPort, errPort := os.LookupEnv("MYSQL_PORT")
	// if !errPort {
	// 	log.Warnf("MYSQL_PORT: got %s expected value", mysqlPort)
	// }

	mysqlUser, errUser := os.LookupEnv("MYSQL_USER")
	if !errUser {
		log.Warnf("MYSQL_USER: got %s expected value", mysqlUser)
	}

	mysqlPass, errPass := os.LookupEnv("MYSQL_PASS")
	if !errPass {
		log.Warnf("MYSQL_PASS: got %s expected value", mysqlPass)
	}

	mysqlDbName, errName := os.LookupEnv("MYSQL_DB")
	if !errName {
		log.Warnf("MYSQL_DB: got %s expected value", mysqlDbName)
	}

	// Create the database handle, confirm driver is present
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		mysqlUser, mysqlPass, mysqlHost, "3306", mysqlDbName,
	)

	db, err := sql.Open(DATABASEDRIVER, connection)
	log.Info("MySQL Connection: ")
	log.Info(connection)

	if err != nil {
		//log.Errorf("OpenConnection: %s", err.Error())
		panic(err)
	}
	return db
}

// NewMySQLCustomersRepository Database Connection
func NewMySQLCustomersRepository(db *sql.DB) *MySQLCustomersRepository {
	return &MySQLCustomersRepository{db}
}

// Close database connection
func (repo *MySQLCustomersRepository) Close() {
	//log.Error("Close: executed normally")
	_ = repo.db.Close()
}

// HealthCheck Repository check Database Health
func (repo *MySQLCustomersRepository) HealthCheck() error {
	//log.Errorf("HealthCheck: executed normally.")
	if err := repo.db.Ping(); err != nil {
		defer repo.db.Close() //nolint
		return err
	}
	return nil
}

// Store make a SQL Query to Insert customer's information
func (repo *MySQLCustomersRepository) Store(customer entity.Customer) error {
	//log.Info("Store: executed normally.")
	// Email
	if repo.rowExists("SELECT id FROM customers WHERE email = ?", customer.Email) {
		return entity.ErrEmailExists
	}
	// Phone
	if repo.rowExists("SELECT id FROM customers WHERE phone = ?", customer.Phone) {
		return entity.ErrPhoneExists
	}

	// Generate UUID for customer_uuid field
	customerUUID, _ := uuid.NewRandom()

	// Generate Sha1 encrypted password for password field
	customer.Password = entity.EncryptPassword(customer.Password)

	query := fmt.Sprintf(`INSERT INTO customers (id, customer_uuid, name, last_name, dni, dni_type, email, phone, country, password, created_at, updated_at, deleted_at) 
                          VALUES (NULL, "%s", ?, ?, ?, ?, ?, ?, ?, "%s", CURRENT_TIMESTAMP, NULL, NULL)`, customerUUID.String(), customer.Password)

	_, err := repo.db.Exec(
		query,
		customer.Name,
		customer.LastName,
		customer.Dni,
		customer.DniType,
		customer.Email,
		customer.Phone,
		customer.Country,
	)

	if err != nil {
		// return err
		log.Errorf("Store: %s", err.Error())
		return entity.ErrSQLError
	}
	return nil
}

// GetByUUID make a SQL Query to collect customer's information by ID key
func (repo *MySQLCustomersRepository) GetByUUID(customerUUID string) (*entity.Customer, error) {
	//log.Info("GetByUUID: executed normally")

	okUUID := entity.IsValidUUID(customerUUID)

	// Validate UUID from comming customer request.
	if !okUUID {
		return nil, entity.ErrBadParamInput
	}

	customer := &entity.Customer{}

	query := `SELECT name, last_name, dni, dni_type, email, phone, customer_uuid, country 
			  FROM customers 
			  WHERE customer_uuid = ? 
			  AND deleted_at IS NULL LIMIT 1`

	row := repo.db.QueryRow(query, customerUUID)
	err := row.Scan(
		&customer.Name,
		&customer.LastName,
		&customer.Dni,
		&customer.DniType,
		&customer.Email,
		&customer.Phone,
		&customer.CustomerUUID,
		&customer.Country)

	// SQL Error or Other More Critial Error Ocurred.
	if err != nil {
		//log.Errorf("GetByUUID: %s", err.Error())
		// Customer Not found
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		// SQL Error
		return nil, entity.ErrSQLError
	}
	return customer, nil
}

// Pagination calculator
func (repo *MySQLCustomersRepository) Pagination(page int, limit int) (int, int, int, error) {
	// Find out how many items are in the table
	total, err := repo.getRowCount()

	if err != nil {
		return 0, 0, 0, err //entity.ErrSQLError
	}

	// Calculate starting number
	begin := (limit * page) - limit
	if begin >= 1 {
		begin = limit
	}
	// Calculate number of pages
	pages := (total / limit)

	if (total % limit) != 0 {
		pages++
	}
	if page > pages {
		return begin, pages, total, nil
	}
	return begin, pages, total, nil
}

// Fetch make a SQL Query to collect customer's information by ID key
func (repo *MySQLCustomersRepository) Fetch(page int, limit int) ([]*entity.Customer, int, int, int, error) {
	//log.Info("Fetch: executed normally")

	begin, pages, total, err := repo.Pagination(page, limit)

	if err != nil {
		return nil, begin, pages, 0, err
	}

	//log.Infof("Current Page: %d, pages: %d begin: %d total: %d limit: %d", page, pages, begin, total, limit)

	customers := []*entity.Customer{}

	var SQLQuery = `SELECT name, last_name, dni, dni_type, email, phone, customer_uuid, country FROM customers WHERE deleted_at IS NULL ORDER BY id DESC LIMIT ? OFFSET ?`

	row, err := repo.db.Query(SQLQuery, limit, begin)

	if err != nil {
		//log.Errorf("Fetch: %s", err.Error())
		return nil, total, pages, page, err
	}

	defer row.Close() //nolint

	for row.Next() {

		customer := &entity.Customer{}

		err := row.Scan(
			&customer.Name,
			&customer.LastName,
			&customer.Dni,
			&customer.DniType,
			&customer.Email,
			&customer.Phone,
			&customer.CustomerUUID,
			&customer.Country,
		)

		customers = append(customers, customer)

		if err != nil {
			//log.Errorf("Fetch: %s", err.Error())
			return nil, total, pages, page, err
		}
	}

	if err = row.Err(); err != nil {
		//log.Errorf("Fetch: %s", err.Error())
		return nil, total, pages, page, err
	}

	return customers, total, pages, page, nil
}

// UpdateByUUID Update customer information by UUID key
func (repo *MySQLCustomersRepository) UpdateByUUID(customer entity.Customer, customerUUID string) error {
	//log.Info("UpdateByUUID: executed normally")

	// Email
	if repo.rowExists("SELECT id FROM customers WHERE email = ? AND customer_uuid != ?", customer.Email, customerUUID) {
		return entity.ErrEmailExists
	}
	// Phone
	if repo.rowExists("SELECT id FROM customers WHERE phone = ? AND customer_uuid != ?", customer.Phone, customerUUID) {
		return entity.ErrPhoneExists
	}

	customer.BeforeUpdate()

	cols, vals := entity.SQLUpdate(customer)
	vals = append(vals, customerUUID)
	// Update customer's account
	query := fmt.Sprintf(`UPDATE customers SET %s ,updated_at = CURRENT_TIMESTAMP WHERE customer_uuid = ? AND deleted_at IS NULL LIMIT 1`, cols)
	result, err := repo.db.Exec(query, vals...)

	// SQL Error or Other More Critial Error Ocurred.
	if err != nil {
		// SQL Error
		//log.Errorf("UpdateByUUID: %s", err.Error())
		return entity.ErrSQLError
	}
	// Check if UPDATE Query actually affected customer's information by customerUUID
	if rows, _ := result.RowsAffected(); rows == 0 {
		//log.Errorf("UpdateByUUID: %s", err.Error())
		// Customer Not found
		return entity.ErrNotFound
	}
	return nil
}

// DeleteByUUID make a SQL Query to Delete customer's information by ID key
func (repo *MySQLCustomersRepository) DeleteByUUID(customerUUID string) error {
	//log.Info("DeleteByUUID: executed normally")

	query := `UPDATE customers SET deleted_at = CURRENT_TIMESTAMP WHERE customer_uuid = ? AND deleted_at IS NULL LIMIT 1`
	result, err := repo.db.Exec(query, customerUUID)

	// SQL Error or Other More Critial Error Ocurred.
	if err != nil {
		//log.Errorf("DeleteByUUID: %s", err.Error())
		// SQL Error
		return entity.ErrSQLError
	}
	// Check if UPDATE Query actually affected customer's information by customerUUID
	if rows, _ := result.RowsAffected(); rows == 0 {
		//log.Errorf("DeleteByUUID: rows %d", rows)
		// Customer Not found
		return entity.ErrNotFound
	}
	return nil
}

func (repo *MySQLCustomersRepository) rowExists(query string, args ...interface{}) bool {
	//log.Info("rowExists: executed normally")

	var exists bool

	query = fmt.Sprintf("SELECT exists (%s)", query)

	err := repo.db.QueryRow(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		log.Errorf("rowExists: %s", err.Error())
		return false
	}
	return exists
}

// getRowCount counts how many rows are present in customer's table
func (repo *MySQLCustomersRepository) getRowCount() (int, error) {
	//log.Info("getRowCount: executed normally")

	var total int

	query := "SELECT COUNT(id) AS total FROM customers WHERE deleted_at IS NULL"
	err := repo.db.QueryRow(query).Scan(&total)
	//log.Infof("getRowCount: total %d", total)

	if err != nil && err != sql.ErrNoRows {
		log.Errorf("getRowCount: %s", err.Error())
		return 0, err
	}
	return total, nil

}
