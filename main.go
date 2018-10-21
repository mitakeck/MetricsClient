package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	url       = os.Getenv("METRICS_API")
	token     = os.Getenv("METRICS_TOKEN")
	namespace = os.Getenv("METRICS_NAMESPACE")
)

func main() {
	if url == "" || token == "" || namespace == "" {
		log.Fatal("YOU NEED TO SET `METRICS_API` AND `METRICS_TOKEN` AND `METRICS_NAMESPACE`")
	}

	payload, err := generatePayload()
	if err != nil {
		panic(err)
	}
	err = postMetric(payload)
	if err != nil {
		panic(err)
	}
}

func generatePayload() (string, error) {
	metric := []Metric{}

	con, err := getConnectivityMetrics()
	if err == nil {
		metric = append(metric, con...)
	} else {
		fmt.Printf("Eror: Connectivity\n")
	}

	mem, err := getMemoryMetics()
	if err == nil {
		metric = append(metric, mem...)
	} else {
		fmt.Printf("Error: Memory\n")
	}

	swap, err := getSwapMetrics()
	if err == nil {
		metric = append(metric, swap...)
	} else {
		fmt.Printf("Error: Swap\n")
	}

	cpu, err := getCPUMetrics()
	if err == nil {
		metric = append(metric, cpu...)
	} else {
		fmt.Printf("Error: CPU\n")
	}

	cpus, err := getCPUMetricsSummary()
	if err == nil {
		metric = append(metric, cpus...)
	} else {
		fmt.Printf("Error: CPU Summary\n")
	}

	disk, err := getDiskMetrics()
	if err == nil {
		metric = append(metric, disk...)
	} else {
		fmt.Printf("Error: Disk\n")
	}

	load, err := getLoadMetrics()
	if err == nil {
		metric = append(metric, load...)
	} else {
		fmt.Printf("Error: Load\n")
	}

	network, err := getNetworkMetrics()
	if err == nil {
		metric = append(metric, network...)
	} else {
		fmt.Printf("Error: Network\n")
	}

	uptime, err := getUptime()
	if err == nil {
		metric = append(metric, uptime...)
	} else {
		fmt.Printf("Error: Uptime\n")
	}

	data := &Payload{
		Namespace: namespace,
		Data:      metric,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

func postMetric(payload string) error {
	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	req.Header.Add("x-api-key", token)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
