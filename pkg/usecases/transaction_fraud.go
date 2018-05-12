// Package usecases containces all method for performing allthe use cases of the application
// This package does not dpened on any other package except the domain. This package contains
// methods which can run indpendent of choice of db, web framework etc
package usecases

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/anirbanroydas/fraud_police/pkg/domain"
)

// TransactionRequest is kind of a DTO which is used by CheckFraudulency method of
// TransactionFraudProcessor to process the transaction
type TransactionRequest struct {
	TransactionID uint64
	UserID        uint64
	PaymentMethod string
	PaymentInfo   interface{}
	Order         interface{}
}

// TransactionFraudProcessor is a usecase interactor which has methods which checks if a
// transactino if fraudulent or not using other usecase interactors or interfaces.
// It uses a FraudProcessor to process the fraud, a TransactionValidator to validate the transaction
// and a TransactionRepo to store and retreive history transactions
type TransactionFraudProcessor struct {
	TransactionRepo domain.TransactionHistoryRepository
	FraudProcessor  FraudProcessor
	Validator       TransactionValidator
}

// CheckFraudulency use_case takes a TransactionRequest object as input and creates a domain level
// Transaction  object and sends it to a FraudProcess which process it from there on, asynchronously.
// It returns an error if there is a problem in any of the above processes.
func (t TransactionFraudProcessor) CheckFraudulency(tReq TransactionRequest) (bool, error) {
	var err error
	var isFraud bool
	// step 1:  validate the Transaction Request -> Order, PaymentMethod, PaymentInfo
	err = t.Validator.Validate(tReq)
	if err != nil {
		return isFraud, errors.Wrap(err, "TransactionValidator's Validator couldn't validate TransactionRequest")
	}
	// step 2: Create new domain.Transaction ojbect with fraud status false and transaction status pending
	transaction := t.createTransaction(tReq)

	// step e: Check Trnasaction for Fraud
	isFraud, err = t.FraudProcessor.isFraud(transaction, t.TransactionRepo)
	if err != nil {
		return isFraud, errors.New(fmt.Sprintf("Transaction is Fraudulent: %v", tReq))
	}
	return isFraud, nil
}

// createTransaction is a helper function which takes TransactionRequest object and returns pointer instance of domain.Transaction
func (t TransactionFraudProcessor) createTransaction(tReq TransactionRequest) *domain.Transaction {
	return domain.NewTransaction(tReq.TransactionID, tReq.UserID, tReq.PaymentMethod, tReq.PaymentInfo, tReq.Order)
}

// NewTransactionFraudProcessor is factory method/constructor which returns an instance of TransactionFraudProcessor
func NewTransactionFraudProcessor(tRepo domain.TransactionHistoryRepository, f FraudProcessor, v TransactionValidator) TransactionFraudProcessor {
	t := TransactionFraudProcessor{
		TransactionRepo: tRepo,
		FraudProcessor:  f,
		Validator:       v,
	}
	return t
}

// NewTransactionRequest is a factory method/ constructor which returns a pointer instance of domain.Transaction
func NewTransactionRequest(transID, userID uint64, payMethod string, payInfo, order interface{}) TransactionRequest {
	t := TransactionRequest{
		TransactionID: transID,
		UserID:        userID,
		PaymentMethod: payMethod,
		PaymentInfo:   payInfo,
		Order:         order,
	}
	return t
}

// TransactionValidator is an interface which has one method called Validate which takes a TransactionRequest object
// and returns error if its not valid
type TransactionValidator interface {
	Validate(TransactionRequest) error
}

// dumbTransactionValidator implements the TransactionValidator interface and has very simple
// dumb logic to validate a TransactionRequest
type dumbTransactionValidator struct{}

func (t dumbTransactionValidator) Validate(tReq TransactionRequest) error {
	if !t.validOrder(tReq) {
		return errors.New(fmt.Sprintf("DumbTransactionValidator: Bad Transaction Order | transactionRequest: %v", tReq))
	}
	if !t.validPaymentMethod(tReq) {
		return errors.New(fmt.Sprintf("DumbTransactionValidator: Bad Transaction PaymentMethod | transactionRequest: %v", tReq))
	}
	if !t.validPaymentInfo(tReq) {
		return errors.New(fmt.Sprintf("DumbTransactionValidator: Bad Transaction PaymentInfo | transactionRequest: %v", tReq))
	}

	return nil
}

func (t dumbTransactionValidator) validOrder(tReq TransactionRequest) bool {
	return true
}

func (t dumbTransactionValidator) validPaymentMethod(tReq TransactionRequest) bool {
	return true
}

func (t dumbTransactionValidator) validPaymentInfo(tReq TransactionRequest) bool {
	return true
}

// NewDumbTransactionValidator is a facctory method/constructor which returns dumbTransactionValidator which a type
// of TransactionValidator(interface)
func NewDumbTransactionValidator() dumbTransactionValidator {
	return dumbTransactionValidator{}
}
