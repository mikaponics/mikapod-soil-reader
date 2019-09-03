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

* ``Go 1.12.7``

## Installation

1. Please visit the [Mikapod Soil (Arduino) device](https://github.com/mikaponics/mikapod-soil-arduino) repository and setup the external device and connect it to your development machine.

2. Please find out what USB port your external device is connected on and export the following environment variable to your console. Note: please replace ``/dev/cu.usbmodem1D132101`` with the value on your machine.

    ```
    export MIKAPOD_SOIL_READER_DEVICE_PATH=/dev/cu.usbmodem1D132101
    ```

3. Get our latest code.

    ```
    go get -u github.com/mikaponics/mikapod-soil-reader
    ```

4. Install the depencies for this project.

    ```
    go get -u google.golang.org/grpc
    go get -u github.com/tarm/serial
    ```

5. Run our application.

    ```
    cd github.com/mikaponics/mikapod-soil-reader
    go run main.go
    ```

6. If you see a message saying ``gRPC server is running`` then the application has been successfully started.


## Production
The following instructions are specific to getting setup for [Raspberry Pi](https://www.raspberrypi.org/).

### Deployment

1. Please visit the [Mikapod Soil (Arduino) device](https://github.com/mikaponics/mikapod-soil-arduino) repository and setup the external device and connect it to your development machine.

2. Please find out what USB port your external device is connected on and export the following environment variable to your console. Note: please replace ``/dev/ttyACM0`` with the value on your machine.

    ```
    export MIKAPOD_SOIL_READER_DEVICE_PATH=/dev/ttyACM0
    ```

3. (Optional) If already installed old golang with apt-get and you want to upgrade to the latest version. Run the following:

    ```
    sudo apt remove golang
    sudo apt-get autoremove
    source .profile
    ```

4. Install [Golang 1.11.8]():

    ```
    wget https://storage.googleapis.com/golang/go1.11.8.linux-armv6l.tar.gz
    sudo tar -C /usr/local -xzf go1.11.8.linux-armv6l.tar.gz
    export PATH=$PATH:/usr/local/go/bin # put into ~/.profile
    ```

5. Confirm we are using the correct version:

    ```
    go version
    ```

4. Install ``git``:

    ```
    sudo apt install git
    ```

5. Get our latest code.

    ```
    go get -u github.com/mikaponics/mikapod-soil-reader
    ```

6. Install the depencies for this project.

    ```
    go get -u google.golang.org/grpc
    go get -u github.com/tarm/serial
    ```

7. Go to our application directory.

    ```
    cd ~/go/src/github.com/mikaponics/mikapod-soil-reader
    ```

8. (Optional) Confirm our application builds on the raspberry pi device. You now should see a message saying ``gRPC server is running`` then the application is running.

    ```
    go run main.go
    ```

9. Build for the ARM device and install it in our ``~/go/bin`` folder:

    ```
    go install
    ```

### Operation

1. While being logged in as ``pi`` run the following:

    ```
    sudo vi /etc/systemd/system/mikapod-soil-reader.service
    ```

2. Copy and paste the following contents.

    ```
    [Unit]
    Description=Mikapod Soil Reader Daemon
    After=multi-user.target

    [Service]
    Environment=MIKAPOD_SOIL_READER_DEVICE_PATH=/dev/ttyACM0
    Type=idle
    ExecStart=/home/pi/go/src/github.com/mikaponics/mikapod-soil-reader/cmd/mikapod-soil-reader/mikapod-soil-reader
    Restart=on-failure
    KillSignal=SIGTERM

    [Install]
    WantedBy=multi-user.target
    ```

3. We can now start the Gunicorn service we created and enable it so that it starts at boot:

    ```
    sudo systemctl start mikapod-soil-reader
    sudo systemctl enable mikapod-soil-reader
    ```

4. Confirm our service is running.

    ```
    sudo systemctl status mikapod-soil-reader.service
    ```

5. If the service is working correctly you should see something like this at the bottom:

    ```
    raspberrypi systemd[1]: Started Mikapod Soil Reader Daemon.
    ```

6. Congradulations, you have setup instrumentation micro-service! All other micro-services can now poll the latest data from the soil-reader we have attached.

7. If you see any problems, run the following service to see what is wrong. More information can be found in [this article](https://unix.stackexchange.com/a/225407).

    ```
    sudo journalctl -u mikapod-soil-reader
    ```

8. To reload the latest modifications to systemctl file.

    ```
    sudo systemctl daemon-reload
    ```

## License

This application is licensed under the **BSD** license. See [LICENSE](LICENSE) for more information.
