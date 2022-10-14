package routes

import (
	controller "data_logger/controllers"
	"fmt"
	"log"
	"net/http"
)

func InitServer() {
	http.HandleFunc("/", controller.GetAllData) // Update this line of code
	http.HandleFunc("/filter", controller.GetData)
	http.HandleFunc("/delete", controller.DeleteData)
	http.HandleFunc("/count", controller.CountData)
	http.HandleFunc("/find/count", controller.FilterAndCountData)
	http.HandleFunc("/itens", controller.ListItens)

	fmt.Printf("Starting server at port 8883\n")
	if err := http.ListenAndServe("0.0.0.0:8883", nil); err != nil {
		log.Fatal(err)
	}
}
