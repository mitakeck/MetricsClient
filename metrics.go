package main

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func generateMetric(name string, value float64, dimension_name string, dimension_value string) Metric {
	return Metric{
		MetricName: name,
		Value:      value,
		Dimensions: []Dimension{
			Dimension{
				Name:  dimension_name,
				Value: dimension_value,
			},
		},
	}
}

func getLoadMetrics() ([]Metric, error) {
	loads, err := load.Avg()
	if err != nil {
		return nil, err
	}

	ret := make([]Metric, 3)
	ret[0] = generateMetric("load", float64(loads.Load1), "load", "load.1")
	ret[1] = generateMetric("load", float64(loads.Load5), "load", "load.5")
	ret[2] = generateMetric("load", float64(loads.Load15), "load", "load.15")

	return ret, nil
}

// func getUptime() (Values, error) {
// 	uptime, err := host.Uptime()
//   if err != nil {
//     return nil, err
//   }
//
//   ret := map[string]float64 {
//     "uptime": float64(uptime)
//   }
//
//   return ret
// }

func getCPUMetrics() ([]Metric, error) {
	cpus, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	ret := []Metric{}
	for _, c := range cpus {
		ret = append(ret, generateMetric("cpu", c.User, "cpu"+c.CPU, "cpu."+c.CPU+".user"))
		ret = append(ret, generateMetric("cpu", c.System, "cpu"+c.CPU, "cpu."+c.CPU+".system"))
		ret = append(ret, generateMetric("cpu", c.Idle, "cpu"+c.CPU, "cpu."+c.CPU+".idle"))
		ret = append(ret, generateMetric("cpu", c.Nice, "cpu"+c.CPU, "cpu."+c.CPU+".nice"))
		ret = append(ret, generateMetric("cpu", c.Iowait, "cpu"+c.CPU, "cpu."+c.CPU+".iowait"))
		ret = append(ret, generateMetric("cpu", c.Irq, "cpu"+c.CPU, "cpu."+c.CPU+".irq"))
		ret = append(ret, generateMetric("cpu", c.Softirq, "cpu"+c.CPU, "cpu."+c.CPU+".softirq"))
		ret = append(ret, generateMetric("cpu", c.Steal, "cpu"+c.CPU, "cpu."+c.CPU+".steal"))
		ret = append(ret, generateMetric("cpu", c.Guest, "cpu"+c.CPU, "cpu."+c.CPU+".guest"))
		ret = append(ret, generateMetric("cpu", c.GuestNice, "cpu"+c.CPU, "cpu."+c.CPU+".guestnice"))
		ret = append(ret, generateMetric("cpu", c.Stolen, "cpu"+c.CPU, "cpu."+c.CPU+".stolen"))
	}

	return ret, nil
}

func getMemoryMetics() ([]Metric, error) {
	metric, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	ret := make([]Metric, 10)
	ret[0] = generateMetric("memory", float64(metric.Total), "memory", "memory.total")
	ret[1] = generateMetric("memory", float64(metric.Available), "memory", "memory.available")
	ret[2] = generateMetric("memory", float64(metric.Used), "memory", "memory.used")
	ret[3] = generateMetric("memory", float64(metric.UsedPercent), "memory", "memory.percent")
	ret[4] = generateMetric("memory", float64(metric.Free), "memory", "memory.free")
	ret[5] = generateMetric("memory", float64(metric.Free), "memory", "memory.active")
	ret[6] = generateMetric("memory", float64(metric.Inactive), "memory", "memory.inactive")
	ret[7] = generateMetric("memory", float64(metric.Wired), "memory", "memory.wired")
	ret[8] = generateMetric("memory", float64(metric.Buffers), "memory", "memory.buffers")
	ret[9] = generateMetric("memory", float64(metric.Cached), "memory", "memory.cached")

	return ret, nil
}

func getDiskMetrics() ([]Metric, error) {
	disks, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	// ret := make(map[string]float64, 4*len(disks))
	ret := []Metric{}

	for _, d := range disks {
		metric, _ := disk.Usage(d.Mountpoint)
		ret = append(ret, generateMetric("disk", float64(metric.Total), "disk"+d.Device, "disk."+d.Device+".total"))
		ret = append(ret, generateMetric("disk", float64(metric.Free), "disk"+d.Device, "disk."+d.Device+".free"))
		ret = append(ret, generateMetric("disk", float64(metric.Used), "disk"+d.Device, "disk."+d.Device+".used"))
		ret = append(ret, generateMetric("disk", float64(metric.UsedPercent), "disk"+d.Device, "disk."+d.Device+".percent"))
	}

	return ret, nil
}
