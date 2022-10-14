package main

import (
	"data_logger/mqtt"
	"data_logger/routes"
	"fmt"
)

func main() {
	fmt.Print("Hello Word!\n")

	mqtt.InitMqttClient() // start mqtt manager
	routes.InitServer()   // start web server
}
