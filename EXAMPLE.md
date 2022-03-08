package main

import (
	"fmt"
	"os"

	dsSensor "github.com/adrkakol/ds18b20-go"
)

func main() {
	sensorAddress := ""
	sensor := dsSensor.Init(sensorAddress)
	temp, err := sensor.GetTemperature()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("temp: %s", temp)
	return
}
