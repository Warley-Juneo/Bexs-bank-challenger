package main

import (
	"github.com/wjuneo/bexs/postgres"
	"github.com/wjuneo/bexs/repository/partnerrepository"
	"github.com/wjuneo/bexs/services/partnerservices"

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
	postgres.RunMigrations()
	defer conn.Close()

	router := mux.NewRouter()
	partnerRepository := partnerrepository.NewPartnerRepository(conn)
	partnerService := partnerservices.NewPartnerService(partnerRepository)

	port := viper.GetString("server.port")
	router.HandleFunc("/api/v1/partners", partnerService.SavePartners).Methods("POST")
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
