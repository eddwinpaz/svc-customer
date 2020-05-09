package entity

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

// Customer business model
type Customer struct {
	CustomerUUID string `json:"id,omitempty" sql:"customer_uuid" `
	Name         string `json:"name,omitempty" sql:"name"`           // validate:"max=30"
	LastName     string `json:"last_name,omitempty" sql:"last_name"` // validate:"max=30"
	Dni          string `json:"dni,omitempty" sql:"dni"`
	DniType      string `json:"dni_type,omitempty" sql:"dni_type"`
	Phone        string `json:"phone,omitempty" sql:"phone"`       // validate:"min=3,max=20"
	Country      string `json:"country,omitempty" sql:"country"`   // validate:"min=2,max=2"
	Email        string `json:"email,omitempty" sql:"email"`       // validate:"nonzero,min=6,max=50,regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"
	Password     string `json:"password,omitempty" sql:"password"` // validate:"nonzero,max=15"
	// validate:"nonzero,min=6,max=50,regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$" gorm:"unique"
}

// BeforeCreate executes when customer is about to store Create() GORM method
func (c Customer) BeforeCreate() {
	// Generate UUID for customer_uuid field
	customerUUID, _ := uuid.NewRandom()
	c.CustomerUUID = customerUUID.String()

	// Generate Sha1 encrypted password for password field
	if len(c.Password) != 0 {
		c.Password = EncryptPassword(c.Password)
	}
}

// BeforeUpdate executes when customer is about to update Update() GORM Method
// func (c *Customer) BeforeUpdate(scope *gorm.Scope) error {
func (c *Customer) BeforeUpdate() {
	// Generate Sha1 encrypted password for password field
	if len(c.Password) != 0 {
		c.Password = EncryptPassword(c.Password)
	}
}

// SQLUpdate Update SQL generated on only the fields that not empty.
func SQLUpdate(c Customer) (string, []interface{}) {
	// Local
	fields := reflect.TypeOf(c)
	values := reflect.ValueOf(c)
	num := fields.NumField()

	// Returns
	var sqlColumns = []string{}
	var sqlValues []interface{}

	for i := 0; i < num; i++ {

		field := fields.Field(i)
		value := values.Field(i)

		if value.String() != "" { // len(value.String()) != 0
			// split := strings.Split(field.Tag.Get("sql"), ",")
			column := fmt.Sprintf(" %s = ? ", field.Tag.Get("sql"))

			sqlColumns = append(sqlColumns, column)
			sqlValues = append(sqlValues, value.String())
		}
	}
	return strings.Join(sqlColumns[:], ","), sqlValues
	/*
		cols, values := entity.SQLUpdate(customer)
		columns := strings.Join(cols[:], ",")
		values = append(values, customerUUID)
		Exec(values...)
	*/
}

// Safe void field from showing.
func (c *Customer) Safe() Customer {
	return Customer{
		CustomerUUID: c.CustomerUUID,
		Name:         c.Name,
		LastName:     c.LastName,
		Dni:          c.Dni,
		DniType:      c.DniType,
		Phone:        c.Phone,
		Email:        c.Email,
		Country:      c.Country,
	}
}
