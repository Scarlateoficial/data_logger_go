package models

import (
	"encoding/json"
	"time"
)

type DataModel struct {
	Topic    string    `json:"topic"`
	Payload  string    `json:"payload"`
	Uuid     string    `json:"uuid"`
	Datetime time.Time `json:"datetime"`
}

func DataToJson(data DataModel) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func JsonToData(jsonData string) DataModel {
	var data DataModel
	json.Unmarshal([]byte(jsonData), &data)
	return data
}

func NewDataModel(topic string, payload string, uuid string, datetime time.Time) *DataModel {
	return &DataModel{Topic: topic, Payload: payload, Uuid: uuid, Datetime: datetime}
}
