package account

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	CreateUserRequest struct {
		Email 		string `json:"email"`
		Password 	string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}

func decodeUserRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var createRequest CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&createRequest)
	if err != nil {
		return nil, err
	}
	return createRequest, nil
}


func decodeEmailRequest(ctx context.Context, request *http.Request) (req interface{}, err error) {
	var requestEmail GetUserRequest
	getParams := mux.Vars(request)

	requestEmail = GetUserRequest{
		Id: getParams["id"],
	}

	return requestEmail, nil
}