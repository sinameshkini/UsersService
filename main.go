package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main()  {
	svc := userService{}

	fetchHandler := httptransport.NewServer(
		makeFetchEndpoint(svc),
		decodeFetchRequest,
		encodeResponse,
	)

	createHandler := httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)

	http.Handle("/", fetchHandler)
	http.Handle("/create",createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

