package mqtt

import (
	config "data_logger/configs"
	data "data_logger/models"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	var parts = strings.Split(string(msg.Topic()), "/")

	var uuid = ""

	if len(parts) > 2 {
		uuid = parts[len(parts)-3]
	} else {
		uuid = "unknown"
	}

	var datamodel = data.NewDataModel(msg.Topic(), string(msg.Payload()), uuid, time.Now())

	config.InsertData(datamodel, "data")
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")

	client.Subscribe("#", 0, nil)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func InitMqttClient() {
	var broker = "mosquitto"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("admin")
	opts.SetPassword("admin")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
