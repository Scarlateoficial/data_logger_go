package mqtt

import (
	config "data_logger/configs"
	data "data_logger/models"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	datamodel := data.NewDataModel(msg.Topic(), string(msg.Payload()), "123")

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
	var broker = "127.0.0.1"
	var port = 5222
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
