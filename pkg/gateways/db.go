// Package gateways constains all the interfaces, gateways etc to communicate between other service, db
// and act as a bridge between the infra and usecases
package gateways

import (
	"time"

	"github.com/anirbanroydas/fraud_police/pkg/domain"
)

// DummyDB represent some DB gatewauy whihc will have a DB Client, here DBClient
// for now has been represented as interface{} for dummy pupose, but in real life,
// it will hold somea ctual db client like mongodb or mysql db client
type DummyDB struct {
	DBClient interface{}
}

// DummyTransactionHistoryRepo satifies the domain.TransactionHistoryRepository interface
// This also has DBClient as a parameter which is of type DummyDB, in rel life it will be
// some other concrete dbclient, which the actual FindByID and Store implementation will make use of.
type DummyTransactionHistoryRepo struct {
	DBClient *DummyDB
}

func (tRepo DummyTransactionHistoryRepo) FindByID(someID uint64) (*domain.Transaction, error) {
	// for now just return an empty dummy Transaction just to satisfy
	// the domain.TransactionHistoryRepository interface method
	return &domain.Transaction{}, nil
}

func (tRepo DummyTransactionHistoryRepo) Store(t *domain.Transaction) (uint64, error) {
	// for now just return some random uint64 value just to satisfy
	// the domain.TransactionHistoryRepository interface method
	return uint64(time.Now().Unix()), nil
}

// NewDummyTransactionHistoryRepo is a factory function which returns a DummyTransactionHistoryRepo
func NewDummyTransactionHistoryRepo(db *DummyDB) DummyTransactionHistoryRepo {
	tRepo := DummyTransactionHistoryRepo{
		DBClient: db,
	}
	return tRepo
}

// NewDummyDB is a factory function which returns a DummyDB
func NewDummyDB() *DummyDB {
	// initialise an empty DummyDb with no concrete client
	db := DummyDB{}
	return &db
}
