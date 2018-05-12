// Package infra constains Structs and Methods which represent the actual implementation of all the
// interfaces, it also has controllers and routes specific methods and structs which help in connecting
// the main app(infrastructure) with the rest of the applications' usecases and domain
package infra

import (
	"log"
	"net/http"
	"time"

	"github.com/anirbanroydas/fraud_police/pkg/gateways"
	uc "github.com/anirbanroydas/fraud_police/pkg/usecases"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// App is an application level struct which has methods as the controllers, handlers
// of the web framework.
// This helps in connecting the handlers with other infrastructures like DB, HttpClient,
// MessageBrokers, also it encapsulates all the usecases interactors so that the handlers
// could use those interactors to perform the usecases.
// We could also directly use these encapsulated struct as globals but that would not adhere
// to good design principles. Also this helps in keeping the Dependency Inversion principle solid
type App struct {
	TFP uc.TransactionFraudProcessor
	DB  *gateways.DummyDB
}

// TransactionRequest represents the request object structure expected by the transaction related
// gin's http handlers
type TransactionRequest struct {
	TransactionID uint64      `json:"transactionID" binding:"required"`
	UserID        uint64      `json:"userID" binding:"required"`
	PaymentMethod string      `json:"paymentMethod" binding:"required"`
	PaymentInfo   interface{} `json:"payment" binding:"required"`
	Order         interface{} `json:"order" binding:"required"`
}

///////////////////////////////////////////////////////
// 			Transaction Fraud Checking Handler		//
//////////////////////////////////////////////////////

// CheckTransactionHandler performs the CheckFraudulency usecase using the app's TFP parameter
// It also validate the request first, and returns the appropriate result sent by CheckFraudlency
func (app *App) CheckTransactionHandler(c *gin.Context) {
	var tReq TransactionRequest
	var err error
	resp := false

	// take the raw data received from http request and try to bind it to TransactionRequest
	err = c.ShouldBindWith(&tReq, binding.JSON)
	if err != nil {
		log.Printf("[CheckTransactionHandler] Booking status request is not valid | { error: %v }", err)
		c.JSON(http.StatusBadRequest, getCheckTransactionResponse(400, resp))
		return
	}

	// check if request object is valid or not
	if !isValidTransacionRequest(tReq) {
		log.Printf("[CheckTransactionHandler] Booking status request is not valid | { error: %v }", err)
		c.JSON(http.StatusBadRequest, getCheckTransactionResponse(400, resp))
		return
	}

	// process the transaction request for fraudency
	resp, err = processTransaction(app, tReq)
	if err != nil {
		log.Printf("[CheckTransactionHandler] Booking status request is not valid | { error: %v }", err)
		c.JSON(http.StatusInternalServerError, getCheckTransactionResponse(500, resp))
		return
	}

	// succesfull result
	c.JSON(http.StatusOK, getCheckTransactionResponse(200, resp))
}

// getCheckTransactionResponse is a helper function to return http response
func getCheckTransactionResponse(code uint, isFraud bool) gin.H {
	log.Printf("getCheckTransactionResponse: code: %d, isFraud: %v\n", code, isFraud)
	return gin.H{
		"code": code,
		"message": gin.H{
			"isFraud": isFraud,
		},
	}
}

// isValidTransacionRequest is a helper function which checks if the transactionReqeust object
// receive from the http request is valid or not
func isValidTransacionRequest(tReq TransactionRequest) bool {
	// dummy method, return True for now
	return true
}

// processTransaction is a helper function whcih uses the App.TFP to perform the usecases CheckFraudulency and returns its result
func processTransaction(app *App, tReq TransactionRequest) (bool, error) {
	// create the transactoin Request object used by the transaction fraud processor
	t := uc.NewTransactionRequest(tReq.TransactionID, tReq.UserID, tReq.PaymentMethod, tReq.PaymentInfo, tReq.Order)
	// perform the check use the app.TFP which is the usecases.TranactionFraudProcessor (interactor object for the CheckFraudulency Usecase)
	return app.TFP.CheckFraudulency(t)
}

///////////////////////////////////
// 			Index Route 		//
//////////////////////////////////

// IndexHandler returns basic index page response, this is just to check the connection, not of actual use
func (app *App) IndexHandler(c *gin.Context) {
	// mimic some time to wait
	time.Sleep(1 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "This is the Index Page",
		"success": true,
	})
}
