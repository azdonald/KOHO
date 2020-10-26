package main

import (
	"fmt"
	"testing"
)

func TestHasReachedDailyLimit(t *testing.T) {
	t.Run("5000.12", testDailyLimitFunc(5000.12, true))
	t.Run("500", testDailyLimitFunc(500.00, false))

}

func testDailyLimitFunc(input float64, expected bool) func(t *testing.T) {
	return func(t *testing.T) {
		ut := userTransaction{}
		result := ut.hasReachedDailyLimit(input)
		if result != expected {
			t.Fail()
			t.Logf("expected %v but got %v", expected, result)
		}
	}
}

func TestAddTransaction(t *testing.T) {
	var firstSet = []transaction{{id: "100", amount: 150.0}}
	t.Run("singleTransaction", testAddTransactions(firstSet, 1))

	var secondSet = []transaction{{id: "1100", amount: 150.0}, {id: "1101", amount: 1050.0}}
	t.Run("doubleTransaction", testAddTransactions(secondSet, 2))
}

func testAddTransactions(tranxs []transaction, expected int) func(t *testing.T) {
	return func(t *testing.T) {
		ut := userTransaction{}
		for _, t := range tranxs {
			ut.addTransaction(t)
		}
		result := len(ut.transactions)
		if result != expected {
			t.Fail()
			t.Logf("expected %v but got %v", expected, result)
		}
	}
}

func TestIsDuplicateTransaction(t *testing.T) {
	first := transaction{id: "100", amount: 150.0}
	second := transaction{id: "1000", amount: 150.0}
	ut := userTransaction{}
	ut.addTransaction(first)
	t.Run("first", testDuplicateTransactionFunc(ut, "100", true))
	t.Run("next", testDuplicateTransactionFunc(ut, "1200", false))
	ut.addTransaction(second)
	t.Run("first", testDuplicateTransactionFunc(ut, "1100", false))
}

func testDuplicateTransactionFunc(ut userTransaction, tranxID string, expected bool) func(t *testing.T) {
	return func(t *testing.T) {
		result := ut.isDuplicateTransaction(tranxID)
		if result != expected {
			t.Fail()
			t.Logf("expected %v but got %v", expected, result)
		}
	}
}

func TestTransactionsAdd(t *testing.T) {
	transactions := Transactions{}
	transactions.weekly = make(map[string]map[string]userTransaction)
	transactions.weeklyAmount = make(map[string]float64)

	invalidAmount := LoadData{"111000", "1234", "6500", "2020-09-11"}
	invalidAmountResult := Result{"111000", "1234", false}

	validAmount := LoadData{"11000", "1234", "5000", "2020-09-11"}
	validAmountResult := Result{"11000", "1234", true}

	dailyLimitReached := LoadData{"1000", "1234", "5000", "2020-09-11"}
	dailyLimitReachedResult := Result{"1000", "1234", false}

	newDay := LoadData{"1400", "1234", "5000", "2020-09-12"}
	newDayResult := Result{"1400", "1234", true}

	t.Run("first test", testTransactionAddFunc(transactions, invalidAmount, invalidAmountResult))
	t.Run("second test", testTransactionAddFunc(transactions, validAmount, validAmountResult))
	t.Run("Third test", testTransactionAddFunc(transactions, dailyLimitReached, dailyLimitReachedResult))
	t.Run("Fourth test", testTransactionAddFunc(transactions, newDay, newDayResult))
}

func testTransactionAddFunc(tranxs Transactions, ld LoadData, expected Result) func(t *testing.T) {
	return func(t *testing.T) {
		result := tranxs.add(ld)
		fmt.Println(tranxs)
		if result != expected {
			t.Fail()
			t.Logf("expected %v but got %v", expected, result)
		}
	}
}
