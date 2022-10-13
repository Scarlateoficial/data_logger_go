package controllers

import (
	"data_logger/configs"
	"fmt"
	"net/http"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var body = configs.GetData("data")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", body)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	var topic = r.URL.Query().Get("topic")

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var body = configs.FilterData("data", topic)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", body)
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	var topic = r.URL.Query().Get("topic")

	if r.Method != "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	configs.DeleteData("data", topic)

	fmt.Fprintf(w, "Deleted data with topic: %v", topic)
}
