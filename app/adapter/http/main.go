package main

import (
	"context"
	"fmt"
	"log"
	"net/http"


	"githut.com/warley-juneo/bexs-bank-challenger/adapter/postgres"
	"githut.com/warley-juneo/bexs-bank-challenger/di"
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
	defer conn.Close()

	postgres.RunMigrations()
	userService := di.ConfigUserDI(conn)

	router := mux.NewRouter()
	router.Handle("/users", http.HandlerFunc(userService.Create)).Methods("POST")
	router.Handle("/product", http.HandlerFunc(userService.Fetch)).Queries(
		"page", "{page}",
		"ObjectPerPage", "{ObjectsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"serach", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}