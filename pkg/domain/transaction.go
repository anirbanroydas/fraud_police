// Package domain consists of all the core domain objects of the business/application.
//
// Domain:)
// 2. Transaction - This places an order for a customer (which makes a customer receive the order),
// this contain transcation infomration like order, payment details etc
//
// So by using the **Domain Experts** knowledge and seeing the **Ubiquitous Language**,
// we can derive at some domain objects
//
// This package basically creates all the domain object and expose some interfaces
// to use them but doesnt depend on anything outside the domain layer
package domain

// Transaction is an aggregrate which specifies what is the Order on which a transaction is taking place
// alongs with information like payment details
type Transaction struct {
	TransactionID uint64
	UserID        uint64
	PaymentMethod string
	PaymentInfo   interface{}
	Order         interface{}
}

// TransactionHistoryRepository exposes the interface to store and find transaction from a repository.
type TransactionHistoryRepository interface {
	FindByID(uint64) (*Transaction, error)
	Store(*Transaction) (uint64, error)
}

// Logger is an interface which exposes two methods LogError and LogInfo both of which takes
// a string input and logs the to anykind of writer.
type Logger interface {
	LogError(string)
	LogInfo(string)
}

// NewTransaction is a factory function which returns a pointer to an instance of Transaction
func NewTransaction(transID, userID uint64, payMethod string, payInfo, order interface{}) *Transaction {
	t := Transaction{
		TransactionID: transID,
		UserID:        userID,
		PaymentMethod: payMethod,
		PaymentInfo:   payInfo,
		Order:         order,
	}
	return &t
}
