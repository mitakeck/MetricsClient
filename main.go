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

	mem, err := getMemoryMetics()
	if err != nil {
		panic(err)
	}

	cpu, err := getCPUMetrics()
	if err != nil {
		panic(err)
	}
	metric := append(mem, cpu...)

	disk, err := getDiskMetrics()
	if err != nil {
		panic(err)
	}
	metric = append(metric, disk...)

	load, err := getLoadMetrics()
	if err != nil {
		panic(err)
	}
	metric = append(metric, load...)

	network, err := getNetworkMetrics()
	if err != nil {
		panic(err)
	}
	metric = append(metric, network...)

	uptime, err := getUptime()
	if err != nil {
		panic(err)
	}
	metric = append(metric, uptime...)

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

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
