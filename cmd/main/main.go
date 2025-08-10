package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iusmanof/go-project/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
