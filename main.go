package main

//imports
import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TarunNanduri/goMicroServices/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// logger
	l := log.New(os.Stdout, "MicroServices with Go", log.LstdFlags)
	// create the handlers
	ph := handlers.NewProducts(l)
	sm := mux.NewRouter()
	gb := handlers.NewGoodbye(l)
	hr := handlers.NewHello(l)

	// Registring Hello subroute
	helloRouter := sm.Methods(http.MethodGet).Subrouter()
	helloRouter.HandleFunc("/hello", hr.ServeHTTP)

	// Registring GoodBye subroute
	goodByeRouter := sm.Methods(http.MethodGet).Subrouter()
	goodByeRouter.HandleFunc("/goodbye", gb.ServeHTTP)

	// Register get subroute
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	// Register Put Subroute
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	// MiddleWare validation
	putRouter.Use(ph.MiddlewareValidateProduct)

	// Register post subroute
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	//create a new server
	s := http.Server{
		Addr:         ":3600",           // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	l.Println("Starting server on port 3600")
	err := s.ListenAndServe()
	if err != nil {
		l.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
