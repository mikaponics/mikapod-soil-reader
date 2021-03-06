# Mikapod Soil Reader
[![Go Report Card](https://goreportcard.com/badge/github.com/mikaponics/mikapod-soil-reader)](https://goreportcard.com/report/github.com/mikaponics/mikapod-soil-reader)

## Overview

The purpose of this application is to provide a remote procedure call (RPC) interface over an external Arduino device tailered for the `Mikapod Soil` configuration.

Supports collection of instrument time-series data from 6 different sensors via [Mikapod Soil (Arduino) device](https://github.com/mikaponics/mikapod-soil-arduino):

* Temperature
* Humidity
* Pressure
* Altitude
* Illuminance
* Soil Moisture

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.15.6``

## Installation

1. Please visit the [Mikapod Soil (Arduino) device](https://github.com/mikaponics/mikapod-soil-arduino) repository and setup the external device and connect it to your development machine.

2. Please find out what USB port your external device is connected on and export the following environment variable to your console. Note: please replace ``/dev/cu.usbmodem14201`` with the value on your machine, a Raspberry Pi would most likely have the value ``/dev/ttyACM0``.

    ```
    export MIKAPOD_SOIL_READER_DEVICE_PATH=/dev/cu.usbmodem14301
    ```

3. Download the source code, build and install the application.

    ```
    GO111MODULE=on go get -u github.com/mikaponics/mikapod-soil-reader
    ```

4. Run our application.

    ```
    mikapod-soil-reader
    ```

5. If you see a message saying ``Started soil reader service`` then the application has been successfully started.

## License

This application is licensed under the **BSD** license. See [LICENSE](LICENSE) for more information.
