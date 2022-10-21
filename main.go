package main

import (
	"fmt"
	"net/http"
	"waysfood/database"
	"waysfood/pkg/mysql"
	"waysfood/routes"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000/api/v1/")
	http.ListenAndServe("localhost:5000", r)
}
