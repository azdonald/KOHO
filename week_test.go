package main

import (
	"testing"
)

func TestCanAddNewDay(t *testing.T) {
	input, expected := "2020-01-01", true
	week := Week{}
	result := week.canAddNewDay(input)

	if result != expected {
		t.Fail()
		t.Logf("expected %v but got %v", expected, result)
	}
}

func TestReset(t *testing.T) {
	input := "2020-01-01"
	beforeReset, afterReset := 1, 0
	week := Week{}
	week.addDay(input)
	beforeResetResult := len(week.days)
	week.reset()
	afterResetResult := len(week.days)

	if beforeReset != beforeResetResult {
		t.Fail()
		t.Logf("expected %v but got %v", beforeReset, beforeResetResult)
	}

	if afterReset != afterResetResult {
		t.Fail()
		t.Logf("expected %v but got %v", afterReset, afterResetResult)
	}

}

func TestAddDay(t *testing.T) {
	week := &Week{}
	t.Run("first test", testAddDayFunc(week, "2020-01-01", 1))
	t.Run("second test", testAddDayFunc(week, "2020-01-01", 1))
	t.Run("Third test", testAddDayFunc(week, "2020-01-03", 2))

}

func testAddDayFunc(w *Week, day string, expected int) func(t *testing.T) {
	return func(t *testing.T) {
		w.addDay(day)
		result := len(w.days)
		if result != expected {
			t.Fail()
			t.Logf("expected %v but got %v", expected, result)
		}
	}
}

func TestShouldStartNewWeek(t *testing.T) {
	input, expected := "2020-01-01", false
	week := Week{}
	week.addDay(input)
	result := week.shouldStartNewWeek(input)

	if result != expected {
		t.Fail()
		t.Logf("expected %v but got %v", expected, result)
	}
}
