package main

import (
	"data_logger/mqtt"
	"data_logger/routes"
	"fmt"
)

func main() {
	fmt.Print("Hello Word!\n")

	// start mqtt manager
	mqtt.InitMqttClient()
	routes.InitServer()
}
