package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/anirbanroydas/fraud_police/pkg/domain"
	"github.com/anirbanroydas/fraud_police/pkg/gateways"
	"github.com/anirbanroydas/fraud_police/pkg/infra"
	"github.com/anirbanroydas/fraud_police/pkg/lib"
	uc "github.com/anirbanroydas/fraud_police/pkg/usecases"
)

const (
	SERVICE_BASE_URL = "/service/fraudpolice/api/v1"
)

func main() {
	// read the port to start service on from the environment
	SERVICE_PORT := lib.Getenv("SERVICE_PORT", "9999")

	// create new App instance
	app := &infra.App{}
	addDependencies(app)

	// initialize the web server
	router := gin.Default()
	// pass handlers to routes and pass the app which has all the dependencies
	// that will be used by the handler to perform the usecases
	infra.AddRoutes(router, SERVICE_BASE_URL, app)

	// initialize the server isntance
	server := &http.Server{
		Addr:    ":" + SERVICE_PORT,
		Handler: router,
	}

	// start the server in a goroutine and keep liseting on signals to pefroam graceful shutdown
	go func() {
		// service connections
		log.Println("Starting Server ...")
		log.Printf("Server Listening at: %s\n", server.Addr)
		// ListenAndServe returns only on error, so usually it blocks forever.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 3 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s\n", err)
	}
	log.Println("Server stopped!")
}

// addDependencies is like a dependency injector which initiates all different
// infrastructre objects and instances and adds its to the App instance which can
// be passed anywhere down the dependency tree
func addDependencies(a *infra.App) {
	// get a single DB connection/pool
	a.DB = gateways.NewDummyDB()

	// Add transaction fraud processing usecase interactor
	var tRepo domain.TransactionHistoryRepository
	var f uc.FraudProcessor
	var v uc.TransactionValidator

	// get a concrete implementation of TransactionHistoryRepository interface
	tRepo = gateways.NewDummyTransactionHistoryRepo(a.DB)
	// get a concrete implementation of FraudProcessor interface
	f = uc.NewDummyFraudProcessor()
	// get a concrete implementation of TransactionValidator interface
	v = uc.NewDumbTransactionValidator()

	// get a new TrnasctionFraudProcessor instance and add it to the TFP parameter
	a.TFP = uc.NewTransactionFraudProcessor(tRepo, f, v)
}
