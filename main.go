package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var week *Week
	week = &Week{}

	var transactions *Transactions
	transactions = &Transactions{}
	transactions.weekly = make(map[string]map[string]userTransaction)
	transactions.weeklyAmount = make(map[string]float64)
	
	for scanner.Scan() {
		var transData LoadData
		err := json.Unmarshal([]byte(scanner.Text()), &transData)
		if err != nil {
			log.Fatal(err)
		}

		transData.LoadAmount = transData.LoadAmount[1:]
		transData.Time = transData.Time[0:10]
		if week.shouldStartNewWeek(transData.Time) {
			week.reset()
			transactions.reset()
		}

		week.addDay(transData.Time)
		var result = transactions.add(transData)
		saveResultToFile(result)
	}
}

func saveResultToFile(r Result) {
	file, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(r)

}
