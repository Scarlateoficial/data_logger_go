package routes

import (
	controller "data_logger/controllers"
	"fmt"
	"log"
	"net/http"
)

func InitServer() {
	http.HandleFunc("/", controller.GetAllData) // Update this line of code
	http.HandleFunc("/data", controller.GetData)
	http.HandleFunc("/delete", controller.DeleteData)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
