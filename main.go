package main

import (
	"customer/customer/repository"
	"customer/customer/usecase"
	"customer/logging"
	"net"
	"net/http"
	"os"
	"svc-customer/customerdelivery/web"

	// _ "net/http/pprof"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	logging.InitializeLogging("vendor.log")

	port := ":9000"
	server, err := net.Listen("tcp", port)

	if err != nil {
		log.Infof("Port %s is in use... ", port)
		return
	}
	_ = server.Close()

	// Open Database Connection
	db := repository.OpenConnection()

	// mysql repository init
	repo := repository.NewMySQLCustomersRepository(db)

	// Migrate the schema
	defer repo.Close()

	// delivery/web interface
	handler := &web.Handler{
		GetCustomerUsecase: &usecase.GetCustomerImpl{
			Repo: repo,
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handler.HealthCheck).Methods("GET")
	r.HandleFunc("/", handler.Store).Methods("POST")
	r.HandleFunc("/", handler.Fetch).Methods("GET")
	r.HandleFunc("/{uuid}", handler.GetByUUID).Methods("GET")
	r.HandleFunc("/{uuid}", handler.UpdateByUUID).Methods("PUT")
	r.HandleFunc("/{uuid}", handler.DeleteByUUID).Methods("DELETE")

	// Swagger Documentation
	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Profiler
	// r.HandleFunc("/debug/pprof/", pprof.Index)
	// r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// support linked index from /debug/pprof manually
	// r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	// r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	// r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	// r.Handle("/debug/pprof/block", pprof.Handler("block"))

	log.Infof("Server running on port %s", port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	err = http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(loggedRouter))
	if err != nil {
		log.Errorf("Error starting server: %s\n", err)
		os.Exit(1)
	}

}
