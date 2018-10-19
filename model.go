package main

// Values :
type Values map[string]float64

// Payload : pyaload data struct to lambda function
type Payload struct {
	Namespace string   `json:"Namespace"`
	Data      []Metric `json:"Data"`
}

// Metric : data list
type Metric struct {
	MetricName string      `json:"MetricName"`
	Value      float64     `json:"Value"`
	Dimensions []Dimension `json:"Dimensions"`
}

// Dimension : dimention data
type Dimension struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}
