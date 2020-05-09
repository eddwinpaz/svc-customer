package web

import (
	"encoding/json"
	"net/http"
	"owner/owner/usecase"
	"strconv"
	"svc-customer/customer/entity"

	"github.com/gorilla/mux"
)

// Handler structure for usecase
type Handler struct {
	GetCustomerUsecase usecase.Usecase
}

/*HealthCheck swagger:route GET /customer/healthcheck/ HealthCheck
  Get customer's api status

responses:
   200: swaggerResponse
   404: swaggerResponse
*/
func (handler *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {

	err := handler.GetCustomerUsecase.HealthCheck()
	// Return Error Response
	if err != nil {
		Response(false, "Database unaccessible", nil, w, http.StatusNotFound)
		return
	}
	// Return Successfull
	Response(true, "UP", nil, w, http.StatusOK)
}

/*GetByUUID swagger:route GET /customer/{uuid} GetByUUID
  Get customer's information by passing a valid customer UUID
responses:
   200: swaggerResponse
   404: swaggerResponseFail
*/
func (handler *Handler) GetByUUID(w http.ResponseWriter, r *http.Request) {
	// get customer uuid
	params := mux.Vars(r)
	customerUUID := params["uuid"]

	okUUID := entity.IsValidUUID(customerUUID)

	// Validate UUID from comming customer request.
	if !okUUID {
		Response(false, "Customer UUID is not valid", nil, w, http.StatusNotFound)
		return
	}

	customr, err := handler.GetCustomerUsecase.GetByUUID(customerUUID)

	// Return Error Response
	if err != nil {
		// Verify if it was a 404
		Response(false, err.Error(), nil, w, http.StatusNotFound)
		return
	}
	// Return Successfull
	Response(true, "Customer Found", entity.CustomerResponse{Customer: customr}, w, http.StatusOK)
}

// Pagination Generate Pagination calculations
func (handler *Handler) Pagination(r *http.Request) (int, int) {
	keys := r.URL.Query()

	if keys.Get("page") == "" {
		return 1, 10
	}

	if keys.Get("limit") == "" {
		return 1, 10
	}

	page, _ := strconv.Atoi(keys.Get("page"))
	pageLimit, _ := strconv.Atoi(keys.Get("limit"))

	if page < 0 {
		page = 1
	}

	if pageLimit <= 0 {
		pageLimit = 10
	}

	return page, pageLimit
}

/*Fetch swagger:route GET /customer Fetch
  Get all registered customer's in a list
responses:
   200: swaggerResponseArray
   404: swaggerResponseFail
*/
func (handler *Handler) Fetch(w http.ResponseWriter, r *http.Request) {

	page, pageLimit := handler.Pagination(r)
	//log.Infof("Current Page: %d, limit %d\n", page, pageLimit)

	customers, totalrows, pages, page, err := handler.GetCustomerUsecase.Fetch(page, pageLimit)
	if err != nil {
		// Verify if it was a 404
		Response(false, err.Error(), nil, w, http.StatusNotFound)
		return
	}

	response := entity.CustomerPaginationResponse{
		Customer: customers,
		Pagination: entity.Pagination{
			Pages:       pages,
			CurrentPage: page,
			TotalRows:   totalrows,
		},
	}
	Response(true, "Customers Found", response, w, http.StatusOK)

}

/*UpdateByUUID swagger:route PUT /customer/{uuid} UpdateByUUID
Update existing Customer's Information by passing a valid customer UUID
responses:
   200: swaggerResponseFail
   404: swaggerResponseFail
*/
func (handler *Handler) UpdateByUUID(w http.ResponseWriter, r *http.Request) {

	// Define customer entity model variable
	var customer entity.Customer

	// Obtain UUID from customer/update/{uuid} handler
	params := mux.Vars(r)
	customerUUID := params["uuid"]
	// //log.Infof("customer_handler: UpdateByUUID() UUID: %s", customerUUID)

	// decode customer new data to be updated
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		Response(false, "Invalid JSON object", nil, w, http.StatusNotFound)
		return
	}

	// Validation Rules
	jsonErrors := validateUpdate(customer)

	// validation failed show json entity error message
	if len(jsonErrors) != 0 {
		Response(false, "Information Could not be sent. Please Check Errors", jsonErrors, w, http.StatusNotFound)
		return
	}
	// Update customer Information
	usecaseError := handler.GetCustomerUsecase.UpdateByUUID(customer, customerUUID)
	// Return error if it fails
	if usecaseError != nil {
		// msg := fmt.Sprintf("Information Could not be created. %s", usecaseError.Error())
		Response(false, usecaseError.Error(), nil, w, http.StatusNotFound)
		return
	}
	// Return Successfull Response
	Response(true, "Customer Information Updated", nil, w, http.StatusOK)

}

/*DeleteByUUID swagger:route DELETE /customer/{uuid} DeleteByUUID
  Delete all Customer's data

responses:
   200: swaggerResponseFail
   404: swaggerResponseFail
*/
func (handler *Handler) DeleteByUUID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	customerUUID := params["uuid"]

	okUUID := entity.IsValidUUID(customerUUID)

	// Validate UUID from comming customer request.
	if !okUUID {
		Response(false, "Customer UUID is not valid", nil, w, http.StatusNotFound)
		return
	}

	// Delete Customer information using UUID
	err := handler.GetCustomerUsecase.DeleteByUUID(customerUUID)

	// Return Error Response
	if err != nil {
		Response(false, err.Error(), nil, w, http.StatusNotFound)
		return
	}
	// Return Success Response
	Response(true, "Customer records deleted", nil, w, http.StatusOK)
}

/*Store swagger:route POST /customer/ Store
  Store new Customer's information

responses:
   200: swaggerResponseFail
   404: swaggerResponseFail
*/
func (handler *Handler) Store(w http.ResponseWriter, r *http.Request) {

	var customer entity.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)

	// Return Error Response
	if err != nil {
		Response(false, "Invalid JSON object", nil, w, http.StatusNotFound)
		return
	}

	// Validation Rules
	jsonErrors := validateStore(customer)

	// validation failed show json entity error message
	if len(jsonErrors) != 0 {
		Response(false, "Information Could not be sent. Please Check Errors", jsonErrors, w, http.StatusNotFound)
		return
	}

	usecaseError := handler.GetCustomerUsecase.Store(customer)

	// Return Error Response
	if usecaseError != nil {
		Response(false, usecaseError.Error(), nil, w, http.StatusNotFound)
		return
	}
	// Return Success Response
	Response(true, "Customer was created.", nil, w, http.StatusOK)
}

// Response HTTP returns
func Response(status bool, desc string, data interface{}, w http.ResponseWriter, httpStatus int) {

	// Return Successfull Response
	var response = entity.Response{
		Status:      status,
		Description: desc,
		Data:        data,
	}

	// Return Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(response)
}

func validateUpdate(customer entity.Customer) []string {
	var errs []string
	// var fields []string

	if customer.Name == "" || len(customer.Name) == 0 {
		errs = append(errs, "name cannot be empty")
	}
	if len(customer.Name) > 30 {
		errs = append(errs, "name should be less than 30 chars long")
	}

	if len(customer.Email) > 50 {
		errs = append(errs, "email should be less than 50 chars long")
	}

	if len(customer.Phone) > 20 {
		errs = append(errs, "phone should be less than 20 chars long")
	}

	if len(customer.Country) > 2 {
		errs = append(errs, "country should be less than 2 chars long")
	}

	if customer.LastName == "" || len(customer.LastName) == 0 {
		errs = append(errs, "last_name cannot be empty")
	}
	if len(customer.LastName) > 30 {
		errs = append(errs, "last_name should be less than 30 chars long")
	}
	return errs
}

func validateStore(customer entity.Customer) []string {
	var errs []string
	// var fields []string

	if customer.Name == "" || len(customer.Name) == 0 {
		errs = append(errs, "name cannot be empty")
	}
	if len(customer.Name) > 30 {
		errs = append(errs, "name should be less than 30 chars long")
	}

	if customer.LastName == "" || len(customer.LastName) == 0 {
		errs = append(errs, "last_name cannot be empty")
	}
	if len(customer.LastName) > 30 {
		errs = append(errs, "last_name should be less than 30 chars long")
	}

	if len(customer.Email) == 0 || customer.Email == "" {
		errs = append(errs, "email cannot be empty")
	}

	if len(customer.Email) > 50 {
		errs = append(errs, "email should be less than 50 chars long")
	}

	if len(customer.Phone) > 20 {
		errs = append(errs, "phone should be less than 20 chars long")
	}

	if len(customer.Country) != 2 {
		errs = append(errs, "country should be equal to 2 chars long")
	}

	return errs
}
