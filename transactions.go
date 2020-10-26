package main

import (
	"log"
	"strconv"
)

type transaction struct {
	id     string
	amount float64
}

type userTransaction struct {
	transactions []transaction
	totalAmount  float64
}

func (ut *userTransaction) hasReachedDailyLimit(amount float64) bool {
	return len(ut.transactions) == 3 || (ut.totalAmount+amount) > 5000
}

func (ut *userTransaction) isDuplicateTransaction(transID string) bool {
	for _, trans := range ut.transactions {
		if trans.id == transID {
			return true
		}
	}
	return false
}

func (ut *userTransaction) addTransaction(tx transaction) {
	ut.transactions = append(ut.transactions, tx)
	ut.totalAmount = ut.totalAmount + tx.amount
}

// Transactions is our data store that holds weekly transaction data
type Transactions struct {
	weekly       map[string]map[string]userTransaction
	weeklyAmount map[string]float64
}

func (t *Transactions) hasReachedWeeklyLimit(customerID string, amount float64) bool {
	weeklyTotal, ok := t.weeklyAmount[customerID]
	if ok {
		return (weeklyTotal + amount) > 20000.00
	}

	return false
}

func (t *Transactions) reset() {
	t.weekly = make(map[string]map[string]userTransaction)
	t.weeklyAmount = make(map[string]float64)
}

func (t *Transactions) updateUserTransaction(ut userTransaction, ld LoadData) {
	t.weekly[ld.Time][ld.CustomerID] = ut
	t.weeklyAmount[ld.CustomerID] = t.weeklyAmount[ld.CustomerID] + ut.totalAmount
}

func (t *Transactions) add(ld LoadData) Result {
	// convert amount from string to float64
	amount, err := strconv.ParseFloat(ld.LoadAmount, 64)
	if err != nil {
		log.Fatal(err)
	}

	if amount > 5000.00 {
		return transactionResult(ld, false)
	}

	if t.hasReachedWeeklyLimit(ld.CustomerID, amount) {
		return transactionResult(ld, false)
	}

	customerID, ok := t.weekly[ld.Time]
	ut := userTransaction{}
	tranx := transaction{id: ld.ID, amount: amount}

	if !ok {
		ut.addTransaction(tranx)
		t.weekly[ld.Time] = make(map[string]userTransaction)
		t.updateUserTransaction(ut, ld)
		return transactionResult(ld, true)
	}

	userTranx, ok := customerID[ld.CustomerID]
	if !ok {
		ut.addTransaction(tranx)
		t.updateUserTransaction(ut, ld)
		return transactionResult(ld, true)
	}
	if userTranx.hasReachedDailyLimit(amount) || userTranx.isDuplicateTransaction(ld.ID) {
		return transactionResult(ld, false)
	}

	userTranx.addTransaction(tranx)
	t.updateUserTransaction(userTranx, ld)
	return transactionResult(ld, true)
}

func transactionResult(ld LoadData, status bool) Result {
	result := Result{ld.ID, ld.CustomerID, status}
	return result
}
