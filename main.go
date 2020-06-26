package main

import (
	"fmt"
	"covid-19/api/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func main() {
	route := routes.GetRoutes()
	handler := cors.AllowAll().Handler(route)
	http.Handle("/", route)
	port := strconv.Itoa(8081)
	fmt.Println("listening at port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
