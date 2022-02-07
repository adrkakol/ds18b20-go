# Package DS18B20

```go
import "github.com/adrkakol/ds18b20-go"
```

This package supports 1-wire temperature sensor DS18B20.
Package dedicated to Raspberry Pi devices.

This program reads data from 1-wire termometer.


## Usage


DS18B20 represents the temperature sensor.

Firstly, create an instance of DS18B20 by calling Init method with the address of your sensor.
Next you can use GetTemperature method to read the temperature from the sensor.
 
```go
ds := DS18B20.Init("28-01020304")
ds.GetTemperature() // result: 18.9
```


# Configuration of your Raspberry Pi

## Enable 1-wire bus

First, make sure the 1-wire interface is enabled.

You can do this in 2 ways:
```sh
sudo raspi-config
# in raspi config find interfaces -> 1-wire -> enable
```

or edit the config file:
```sh
sudo nano /boot/config.txt
```
find the line and uncomment it (by removing # sign)
```sh
# dtoverlay=w1-gpio
```

By default the 1-wire pin is set to GPIO-4, but you can use custom GPIO pin to connect the sensor.
You can do it as follows. In the `/boot/config.txt` file add `,gpiopin=x` where x is your custom 1-wire GPIO pin.
```sh
dtoverlay=w1-gpio,gpiopin=21
```

You can check pins numeration on this site: https://pinout.xyz/pinout/1_wire

Next, setup the pullup option for the chosen gpio pin.
```sh
sudo dtoverlay w1-gpio gpiopin=21 pullup=0
sudo modprobe w1-gpio
```

Now, when it is done, you can connect your temperature sensor and check if it is visible in connecteded devices:
```sh
ls /sys/bus/w1/devices/
```

It will list all connected 1-wire devices. You should see a device simillar to `28-01195240e3ff`.
28-01195240e3ff is an address of my temperature sensor. Yours can be named differently.

When you open the file w1_slave of your sensor, it will show the temperature. 
```sh
less /sys/bus/w1/devices/28-01195240e3ff/w1_slave

# 32 01 4b 46 7f ff 0c 10 8f : crc=8f YES
# 32 01 4b 46 7f ff 0c 10 8f t=19125
# t=19125 is the temperature with 3 decimal places
```

If you can't find a w1_slave, it means you haven't set the 1-wire bus properly or your device is not correctly connected.
You can try rebooting or checking if the connection is correct or maybe you haven't set up the pullup option or gave incorrect pin number.

# Building for raspberry

```sh
GOOS=linux GOARCH=arm CC_FOR_TARGET=arm-linux-gnueabi-gcc go build
```


