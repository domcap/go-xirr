//go:build linux || darwin || windows
// +build linux darwin windows

// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "C"

import (
	"log"
	"sync"
	"time"

	. "github.com/AndreyZWorkAccount/XIRR/netPresentValue"
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	"github.com/AndreyZWorkAccount/XIRR/xirr"
)

var size int
var amounts []float64
var dates []string
var mtx sync.Mutex

//export reset_data
func reset_data() {
	amounts = nil
	dates = nil
}

//export add_values
func add_values(aAmount float64, aDateString string) {
	mtx.Lock()
	defer mtx.Unlock()
	amounts = append(amounts, aAmount)
	dates = append(dates, aDateString)
}

//export calc_xirr
func calc_xirr() float64 {
	if len(amounts) == 0 || len(dates) == 0 || len(amounts) != len(dates) {
		log.Fatal("Invalid Data provided - size check")
		return -1
	}

	var layout string
	strLen := len(dates[0])
	if strLen == 10 {
		layout = "2006-01-02"
	}
	if strLen == 25 {
		layout = time.RFC3339
	}

	datasetSize := len(amounts)
	methodParams := Params{MaxIterationsCount: 1000, Epsilon: 0.0000001}

	var payments = make([]IPayment, datasetSize)

	for i := 0; i < datasetSize; i++ {
		time, error := time.Parse(layout, dates[i])
		if error != nil {
			log.Println(strLen)
			log.Fatal("Bad Date format: ", layout, " -> DATE: ", dates[i], " -> ERROR: ", error)
			return -1
		}
		payments[i] = NewPayment(amounts[i], time)
	}

	var xirrCalcMethod xirr.CalcMethod = xirr.NewXIRRMethod(0.00000001, 365, &methodParams)
	var orderedPayments = xirr.OrderPayments(payments)
	res := xirrCalcMethod.Calculate(orderedPayments)

	if !res.IsSolution() {
		log.Println("Successful solution is expected.")
		return -1
	}
	if res.Error() != nil {
		msg := res.Error()
		log.Fatal(msg.Error())
		return -1
	}
	return res.Value()
}

func main() {
	size := 20
	amounts := [...]float64{
		-1470000,
		2162.92,
		7236.075,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		6787.083333,
		21487.08333,
		9062.083333,
		9062.083333,
		9062.083333,
		9062.083333,
		9062.083333,
		1479062.083,
	}

	dates := [...]string{
		"2022-03-21",
		"2022-03-21",
		"2022-05-01",
		"2022-06-01",
		"2022-07-01",
		"2022-08-01",
		"2022-09-01",
		"2022-10-01",
		"2022-11-01",
		"2022-12-01",
		"2023-01-01",
		"2023-02-01",
		"2023-03-01",
		"2023-04-01",
		"2023-05-01",
		"2023-06-01",
		"2023-07-01",
		"2023-08-01",
		"2023-09-01",
		"2023-10-01",
	}
	reset_data()
	for i := 0; i < size; i++ {
		add_values(amounts[i], dates[i])
	}
	res := calc_xirr()
	log.Println("GOT RESULT: ", res)

	reset_data()
	for i := 0; i < size; i++ {
		add_values(amounts[i], dates[i])
	}
	res = calc_xirr()
	log.Println("GOT RESULT: ", res)

}
