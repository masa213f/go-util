package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masa213f/go-util"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	log := util.NewLogger(util.LogEncodingJSON, true)
	server := util.NewHTTPServer(log, ":8080", http.HandlerFunc(handler))
	err := server.Serve(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
