# MetricsClient

[![Circle CI](https://circleci.com/gh/mitakeck/MetricsClient/tree/master.svg?style=shield)](https://circleci.com/gh/mitakeck/MetricsClient/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/mitakeck/MetricsClient)](https://goreportcard.com/report/github.com/mitakeck/MetricsClient) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mitakeck/MetricsClient/blob/master/LICENSE)

MetricsClient

## Install

### Install from source

```bash
#!/bin/bash
go get github.com/mitakeck/MetricsClient
dep ensure
go build
```

### Install from binary

Download from here https://github.com/mitakeck/MetricsClient/releases/tag/v0.0.6

```bash
#!/bin/bash
mv MetricsClient_* /usr/local/bin/MetricsClient
chmod +x /usr/local/bin/MetricsClient
```

## Usage

```bash
#!/bin/bash
METRICS_API="htts://xxxxxxxx" METRICS_TOKEN="xxxxxxx" METRICS_NAMESPACE="xxxxxxxxxx" MetricsClient
```

```bash
#!/bin/bash
export METRICS_API="htts://xxxxxxxx"
export METRICS_TOKEN="xxxxxxx"
export METRICS_NAMESPACE="xxxxxxxxxx"

MetricsClient
```

## Dependencies

- github.com/mackerelio/go-osstat
- github.com/shirou/gopsutil

## Task

- [x] add CircleCI conf
- [x] add connectivity param
- [x] add swap data param
- [x] add cpu percentage param
