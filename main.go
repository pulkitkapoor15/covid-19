package main

import (
	"fmt"
	"covid-19/api/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /idea-box
func main() {
	route := routes.GetRoutes()
	handler := cors.AllowAll().Handler(route)
	http.Handle("/", route)
	port := strconv.Itoa(8081)
	fmt.Println("listening at port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
