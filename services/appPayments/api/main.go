package main

import (
	"github.com/payment/postgres"
	"github.com/payment/repository"
	"github.com/payment/services"

	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	//postgres.RunMigrations()
	defer conn.Close()

	router := mux.NewRouter()
	paymentRepository := repository.NewPaymentRepository(conn)
	paymentService := services.NewPaymentService(paymentRepository)

	port := viper.GetString("server.port")
	router.HandleFunc("/api/v1/payments", paymentService.HandlerRequest).Methods("POST")
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}