package main

import (
	"UsersService/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type fetchRequest struct {

}

type fetchResponse struct {
	Users	[]*models.User	`json:"users"`
	Err 	string			`json:"err,omitempty"`
}

type createRequest struct {
	User 	*models.User	`json:"user"`
}

type createResponse struct {
	User 	*models.User 	`json:"user"`
	Err 	string			`json:"err"`
}

func makeFetchEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		users, err := svc.Fetch()
		if err != nil{
			return fetchResponse{nil, err.Error()}, err
		}
		return fetchResponse{users, ""}, nil
	}
}

func makeCreateEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println(request)
		req := request.(createRequest)
		fmt.Println(req.User)
		user, err := svc.Create(req.User)
		if err != nil{
			return createResponse{nil, err.Error()}, err
		}
		return createResponse{user, ""}, nil
	}
}

func decodeFetchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request fetchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	return request, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
