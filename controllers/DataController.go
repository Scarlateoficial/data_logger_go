package controllers

import (
	"data_logger/configs"
	"fmt"
	"net/http"
	"strings"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var indent = r.URL.Query().Get("indent")

	var body = configs.GetData("data", indent == "true" || indent == "1")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", body)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var value = r.URL.Query().Get("value")
	var arg = r.URL.Query().Get("arg")
	var indent = r.URL.Query().Get("indent")

	if value == "" {
		http.Error(w, "Value is required.", http.StatusPaymentRequired)
		return
	}

	if arg == "" {
		http.Error(w, "Argument is required.", http.StatusPaymentRequired)
		return
	}

	var body = configs.FilterData("data", arg, value, indent == "true" || indent == "1")

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

func CountData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var body = configs.CountData("data")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{'count':%v}", body)
}

func FilterAndCountData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var value = r.URL.Query().Get("value")
	var arg = r.URL.Query().Get("arg")

	if value == "" {
		http.Error(w, "Value is required.", http.StatusPaymentRequired)
		return
	}

	if arg == "" {
		http.Error(w, "Argument is required.", http.StatusPaymentRequired)
		return
	}

	var body = configs.CountFillterData("data", arg, value)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{'count':%v}", body)
}

func ListItens(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var item = r.URL.Query().Get("item")

	if item == "" {
		http.Error(w, "Item is required.", http.StatusPaymentRequired)
		return
	}

	var itens = configs.GetItens("data", item)

	var body string

	body = fmt.Sprintf("{'%v': [%v]}", item, strings.Join(itens, ","))

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", body)
}
