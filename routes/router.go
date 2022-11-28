package routes

import (
	controller "data_logger/controllers"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.GetAllData) // Update this line of code
	mux.HandleFunc("/filter", controller.GetData)
	mux.HandleFunc("/delete", controller.DeleteData)
	mux.HandleFunc("/count", controller.CountData)
	mux.HandleFunc("/find/count", controller.FilterAndCountData)
	mux.HandleFunc("/itens", controller.ListItens)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	fmt.Printf("Starting server at port 8883\n")
	if err := http.ListenAndServe("0.0.0.0:8883", handler); err != nil {
		log.Fatal(err)
	}
}
