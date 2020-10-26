package main

// Week is used to represent a calendar week i.e 7 days
type Week struct {
	days []string
}

// addDay adds a unique day to the week
func (w *Week) addDay(time string) {
	if w.canAddNewDay(time) {
		w.days = append(w.days, time)
	}
}

// checks if its a new day
func (w *Week) canAddNewDay(time string) bool {
	for _, v := range w.days {
		if v == time {
			return false
		}
	}

	return true
}

// reset starts a new week
func (w *Week) reset() {
	w.days = nil
}

// shouldStartNewWeek checks if we should start a new week
func (w *Week) shouldStartNewWeek(time string) bool {
	return w.canAddNewDay(time) && len(w.days) == 7
}
