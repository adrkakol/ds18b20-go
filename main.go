package DS18B20

import (
	"bufio"
	"github.com/joho/godotenv"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DS18B20 struct {
}

func (ds *DS18B20) GetTemperature() float64 {
	var temperature string = getTemperatureFromFile(getSensorFilePath())
	var comaIndex int = len(temperature) - 3
	var temperatureFixed string = temperature[:comaIndex] + "." + temperature[comaIndex:]
	parsed, err := strconv.ParseFloat(temperatureFixed, 64)
	if err != nil {
		log.Fatal("Cannot parse float")
	}
	return math.Round(parsed*100) / 100
}

func getSensorFilePath() string {
	godotenv.Load(".env")

	var sensorAddress = os.Getenv("SENSOR_ADDRESS")
	return "/sys/bus/w1/devices/" + sensorAddress + "/w1_slave"
}

func getTemperatureFromFile(filePath string) string {
	var measuredTemperature string

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	if len(measuredTemperature) < 1 {
		log.Fatal("Unable to read temperature")
	}

	return measuredTemperature
}
