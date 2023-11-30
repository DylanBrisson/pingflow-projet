package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

func main() {
	// Initialiser le routeur
	r := mux.NewRouter()

