//+build cgo

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	metric := []Metric{}

	mem, err := getMemoryMetics()
	if err == nil {
		metric = append(metric, mem...)
	} else {
		fmt.Printf("Errror: %v", err)
	}

	cpu, err := getCPUMetrics()
	if err == nil {
		metric = append(metric, cpu...)
	} else {
		fmt.Printf("Errror: %v", err)
	}

	disk, err := getDiskMetrics()
	if err == nil {
		metric = append(metric, disk...)
	} else {
		fmt.Printf("Errror: %v", err)
	}

	load, err := getLoadMetrics()
	if err == nil {
		metric = append(metric, load...)
	} else {
		fmt.Printf("Errror: %v", err)
	}

	network, err := getNetworkMetrics()
	if err == nil {
		metric = append(metric, network...)
	} else {
		fmt.Printf("Errror: %v", err)
	}

	uptime, err := getUptime()
	if err == nil {
		metric = append(metric, uptime...)
	} else {
		fmt.Printf("Errror: %v", err)
	}

	data := &Payload{
		Namespace: namespace,
		Data:      metric,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payload)))

	req.Header.Add("x-api-key", token)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
