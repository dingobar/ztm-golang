//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

func AvgBytesArray(bytesArray []Bytes) float32 {
	var sum = 0
	for _, usage := range bytesArray {
		sum += int(usage)
	}
	return float32(sum) / float32(len(bytesArray))
}

func AvgFloatArray(floatArray []Celcius) float32 {
	var sum Celcius = 0.0
	for _, usage := range floatArray {
		sum += usage
	}
	return float32(sum) / float32(len(floatArray))
}

type BandwidthUsage struct {
	amount []Bytes
}

func (bu BandwidthUsage) AvgBandwidthUsage() float32 {
	return AvgBytesArray(bu.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (ct CpuTemp) AvgTemp() float32 {
	return AvgFloatArray(ct.temp)
}

type MemoryUsage struct {
	amount []Bytes
}

func (mu MemoryUsage) AvgMemoryUsage() float32 {
	return AvgBytesArray(mu.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{bandwidth, temp, memory}

	fmt.Println(dashboard.AvgBandwidthUsage())
	fmt.Println(dashboard.AvgTemp())
	fmt.Println(dashboard.AvgMemoryUsage())
}
