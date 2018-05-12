// Package usecases containces all method for performing allthe use cases of the application
// This package does not dpened on any other package except the domain. This package contains
// methods which can run indpendent of choice of db, web framework etc
package usecases

import (
	"math/rand"
	"time"

	"github.com/anirbanroydas/fraud_police/pkg/domain"
)

// FraudProcessor is an interface with one method isFraud which takes in a pointer to a Transaction object
// and a TransactionHistoryRepository object and processes the Transaction using information from the
// TransactionHistoryRepository ans returns if transaction is fraudulent or not
type FraudProcessor interface {
	isFraud(*domain.Transaction, domain.TransactionHistoryRepository) (bool, error)
}

// dummyFraudProcessor is an implementation of the interface FraudProcessor
// this is dummy implementation which does a very simplistic dumb isFraud logic
type dummyFraudProcessor struct {
	count int
}

// isFraud is a dumb function to return if a tranasction is fraudulent or not, this function
// does not take into account any real transaction or transactionHistoryRepository information
// It performs a rather simple logic of returning true/false randomly
func (f *dummyFraudProcessor) isFraud(t *domain.Transaction, tRepo domain.TransactionHistoryRepository) (bool, error) {
	// Perform a dummy random fraud Check, real app would use
	// transaction and transactionHistoryRepository to figure this out
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// sleep anything between 1-2 secs randomly
	time.Sleep(time.Second * time.Duration(r1.Intn(3)+1))

	// choose a random number and return Flase when the number is divisible by 3
	// This is to return True 2/3rd of the time and return False 1/4th of the time
	// but still keeping the random behaviour
	// s2 := rand.NewSource(time.Now().UnixNano())
	// r2 := rand.New(s2)
	// num := r2.Intn(999999)
	// if ((num + 1) % 5) == 0 {
	// 	return true, nil
	// }
	f.count = f.count + 1
	if f.count == 10 {
		f.count = 0
		return true, nil
	}
	return false, nil
}

// NewDummyFraudProcessor is a constructor/ factor method which returns an empty
// dummyFraudProcessor which is type of FraudProcessor(interface)
func NewDummyFraudProcessor() *dummyFraudProcessor {
	fp := dummyFraudProcessor{
		count: 0,
	}
	return &fp
}
