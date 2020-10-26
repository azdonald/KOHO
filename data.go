package main

// Result represents the strcuture of the response we give
// after a transaction
type Result struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

// LoadData represents the structure of the incoming data
type LoadData struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	LoadAmount string `json:"load_amount"`
	Time       string
}
