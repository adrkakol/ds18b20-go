// Author: Adrian Kąkol

// This package supports 1-wire temperature sensor DS18B20.
// Package dedicated to Raspberry Pi devices.
//
// This program reads data from 1-wire thermometer.
package DS18B20

import (
	"bufio"
	"errors"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// DS18B20 struct represents the temperature sensor.
//
// Firstly, create an instance of DS18B20.
// Secondly, call Init method with the address of your sensor.
// After these steps you can use GetTemperature method to read the temperature from the sensor.
//
// ds := DS18B20.Init("28-01020304")
//
// temperature, err := ds.GetTemperature() // result: (18.9, nil) orr (0, error)
//
type DS18B20 struct {
	address  string
	filePath string
}

// initialize
func Init(address string) *DS18B20 {
	ds := new(DS18B20)
	ds.address = address
	ds.setSensorFilePath()

	return ds
}

func (ds *DS18B20) setSensorFilePath() {
	ds.filePath = "/sys/bus/w1/devices/" + ds.address + "/w1_slave"
}

func (ds *DS18B20) GetTemperature() (float64, error) {
	temperature, err := ds.getTemperatureFromFile()

	if err != nil {
		return 0, err
	}

	var comaIndex int = len(temperature) - 3
	var temperatureFixed string = temperature[:comaIndex] + "." + temperature[comaIndex:]

	parsed, err := strconv.ParseFloat(temperatureFixed, 64)

	if err != nil {
		return 0, errors.New("Unable to parse temperature.")
	}

	return math.Round(parsed*100) / 100, nil
}

func (ds *DS18B20) getTemperatureFromFile() (string, error) {
	var measuredTemperature string

	file, err := os.Open(ds.filePath)

	if err != nil {
		return "", errors.New("Failed to read the temperature. Cannot open the file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line string = scanner.Text()

		re := regexp.MustCompile("t=.*")
		var temperature string = re.FindString(line)

		if len(temperature) > 0 {
			measuredTemperature = strings.Split(temperature, "=")[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("Failed to read the temperature. Cannot read the file.")
	}

	if len(measuredTemperature) < 1 {
		return "", errors.New("Failed to read the temperature. Cannot parse the file.")
	}

	return measuredTemperature, nil
}
