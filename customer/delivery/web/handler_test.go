package web

import (
	"customer/customer/entity"
	mocks "customer/customer/mocks"
	"customer/customer/usecase"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// GET_CUSTOMER_BY_UUID
func TestHandlerGetByUUIDValid(t *testing.T) {

	path := fmt.Sprintf("/%s", "039d69ee-f9cb-4a3d-87e4-6eb63c302579")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("GetByUUID", "039d69ee-f9cb-4a3d-87e4-6eb63c302579").Return(mocks.MockCustomer, nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.GetByUUID).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"status\":true,\"description\":\"Customer Found\",\"data\":{\"customer\":{\"id\":\"039d69ee-f9cb-4a3d-87e4-6eb63c302579\",\"name\":\"Eddwin\",\"last_name\":\"Paz\",\"dni\":\"264573076\",\"dni_type\":\"DNI\",\"phone\":\"56900000000\",\"country\":\"CL\",\"email\":\"ep@latamig.com\"}}}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestHandlerGetByUUIDInvalid(t *testing.T) {

	req, err := http.NewRequest("GET", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("GetByUUID", "0").Return(mocks.MockCustomer, nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.GetByUUID).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	expected := "{\"status\":false,\"description\":\"Customer UUID is not valid\",\"data\":null}\n"
	assert.Equal(t, rr.Body.String(), expected)

}

func TestHandlerGetByUUIDFailure(t *testing.T) {

	req, err := http.NewRequest("GET", "/039d69ee-f9cb-4a3d-87e4-6eb63c302579", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("GetByUUID", "039d69ee-f9cb-4a3d-87e4-6eb63c302579").Return(mocks.MockCustomer, entity.ErrNotFound)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.GetByUUID).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, rr.Body.String(), mocks.RequestedCustomerNotFound)

}

// DELETE_CUSTOMER_BY_UUID

func TestHandlerDeleteByUUIDSuccess(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/039d69ee-f9cb-4a3d-87e4-6eb63c302579", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("DeleteByUUID", "039d69ee-f9cb-4a3d-87e4-6eb63c302579").Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.DeleteByUUID).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"status\":true,\"description\":\"Customer records deleted\",\"data\":null}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestDeleteByUUIDInvalid(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("delete_customer_by_uuid", "0").Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.DeleteByUUID).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	expected := "{\"status\":false,\"description\":\"Customer UUID is not valid\",\"data\":null}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

// entity.ErrNotFound

func TestDeleteByUUIDInFailure(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/039d69ee-f9cb-4a3d-87e4-6eb63c302579", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("DeleteByUUID", "039d69ee-f9cb-4a3d-87e4-6eb63c302579").Return(entity.ErrNotFound)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.DeleteByUUID).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, rr.Body.String(), mocks.RequestedCustomerNotFound)
}

// STORE_CUSTOMER

func TestStoreValid(t *testing.T) {
	var mock entity.Customer

	mock.Name = mocks.CName
	mock.LastName = mocks.CLastName
	mock.Dni = mocks.CDni
	mock.DniType = mocks.CDniType
	mock.Phone = mocks.CPhone
	mock.Email = mocks.CMail
	mock.Password = mocks.CPassword
	mock.Country = "CL"

	jsonParam := `{"name": "Bad","last_name": "Bunny","dni": "264573076","dni_type": "DNI","email": "baby@gmail.com", "password": "conejomalo","phone": "+130698569","country":"CL"}`

	req, err := http.NewRequest("POST", "/", strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Store", mock).Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Store).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"status\":true,\"description\":\"Customer was created.\",\"data\":null}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestStoreWrongJSON(t *testing.T) {
	var mock entity.Customer

	mock.Name = mocks.CName
	mock.LastName = mocks.CLastName
	mock.Phone = mocks.CPhone
	mock.Email = mocks.CMail
	mock.Password = mocks.CPassword
	mock.Country = "CL"

	jsonParam := `not json object`

	req, err := http.NewRequest("POST", "/", strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Store", mock).Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Store).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, rr.Body.String(), mocks.InvalidJSONObject)
}

func TestStoreMissingField(t *testing.T) {
	var mock entity.Customer
	mock.Name = ""
	mock.LastName = ""
	mock.Phone = mocks.CPhone
	mock.Email = mocks.CMail
	mock.Password = mocks.CPassword
	mock.Country = "CL"

	jsonParam := `{"email": "baby@gmail.com","password": "conejomalo","phone": "+130698569","country":"CL"}`

	req, err := http.NewRequest("POST", "/", strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Store", mock).Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Store).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	expected := "{\"status\":false,\"description\":\"Information Could not be sent. Please Check Errors\",\"data\":[\"name cannot be empty\",\"last_name cannot be empty\"]}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestStoreMissingError(t *testing.T) {

	var mock entity.Customer

	mock.Name = mocks.CName
	mock.LastName = mocks.CLastName
	mock.Phone = mocks.CPhone
	mock.Email = mocks.CMail
	mock.Password = mocks.CPassword
	mock.Country = "CL"

	jsonParam := `{"name": "Bad","last_name": "Bunny","email": "baby@gmail.com",
	"password": "conejomalo","phone": "+130698569","country":"CL"}`

	req, err := http.NewRequest("POST", "/", strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Store", mock).Return(entity.ErrNotFound)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Store).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, rr.Body.String(), mocks.RequestedCustomerNotFound)
}

// UPDATE_CUSTOMER

func TestUpdateByUUIDSuccess(t *testing.T) {

	var mock entity.Customer

	mock.Name = mocks.CName
	mock.LastName = mocks.CLastName
	mock.Phone = mocks.CPhone
	mock.Email = mocks.CMail
	mock.Password = mocks.CPassword
	mock.Country = "CL"

	jsonParam := `{"name": "Bad","last_name": "Bunny","email": "baby@gmail.com","password": "conejomalo","phone": "+130698569","country":"CL"}`

	path := fmt.Sprintf("/%s", mocks.CustomerUUIDValid)
	req, err := http.NewRequest("PUT", path, strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mock, mocks.CustomerUUIDValid).Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.UpdateByUUID).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"status\":true,\"description\":\"Customer Information Updated\",\"data\":null}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestUpdateByUUIDMissingField(t *testing.T) {

	var mock entity.Customer
	mock.Phone = mocks.CPhone
	mock.Email = mocks.CMail
	mock.Password = mocks.CPassword
	mock.Country = "CL"

	jsonParam := `{"email": "baby@gmail.com","password": "conejomalo","phone": "+130698569","country":"CL"}`

	path := fmt.Sprintf("/%s", mocks.CustomerUUIDValid)
	req, err := http.NewRequest("PUT", path, strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mock).Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.UpdateByUUID).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	expected := "{\"status\":false,\"description\":\"Information Could not be sent. Please Check Errors\",\"data\":[\"name cannot be empty\",\"last_name cannot be empty\"]}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestUpdateByUUIDWrongJSON(t *testing.T) {

	jsonParam := `not a json valid`

	path := fmt.Sprintf("/%s", mocks.CustomerUUIDValid)
	req, err := http.NewRequest("PUT", path, strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mocks.MockCustomerPointless, mocks.CustomerUUIDValid).Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.UpdateByUUID).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, rr.Body.String(), mocks.InvalidJSONObject)
}

func TestUpdateByUUIDError(t *testing.T) {

	jsonParam := `not a json valid`

	req, err := http.NewRequest("PUT", "/", strings.NewReader(jsonParam))
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mocks.MockCustomerPointless).Return(entity.ErrNotFound)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", handler.UpdateByUUID).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, rr.Body.String(), mocks.InvalidJSONObject)
}

func TestHandlerFetchCustomerValid(t *testing.T) {

	path := fmt.Sprintf("/%s", "039d69ee-f9cb-4a3d-87e4-6eb63c302579")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("GetByUUID", "039d69ee-f9cb-4a3d-87e4-6eb63c302579").Return(mocks.MockCustomer, nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{uuid}", handler.GetByUUID).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"status\":true,\"description\":\"Customer Found\",\"data\":{\"customer\":{\"id\":\"039d69ee-f9cb-4a3d-87e4-6eb63c302579\",\"name\":\"Eddwin\",\"last_name\":\"Paz\",\"dni\":\"264573076\",\"dni_type\":\"DNI\",\"phone\":\"56900000000\",\"country\":\"CL\",\"email\":\"ep@latamig.com\"}}}\n"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestHandlerHealthCheckValid(t *testing.T) {
	httpURL := "/healthcheck"
	req, err := http.NewRequest("GET", httpURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("HealthCheck").Return(nil)

	handler := &Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: mockCase,
		},
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc(httpURL, handler.HealthCheck).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"status\":true,\"description\":\"UP\",\"data\":null}\n"
	assert.Equal(t, rr.Body.String(), expected)
}
