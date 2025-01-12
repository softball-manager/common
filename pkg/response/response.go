package response

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type SuccessfulCreateResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type ErrorResponse struct {
	DeveloperText string `json:"developerText,omitempty"`
	Status        string `json:"status"`
}

func formatResponse(resp interface{}, statusCode int) events.APIGatewayProxyResponse {
	respJson, err := json.Marshal(resp)
	if err != nil {
		panic("unable to create response")
	}
	respStr := string(respJson)

	return events.APIGatewayProxyResponse{
		Body:       respStr,
		StatusCode: statusCode,
	}
}

func CreateSuccessfulCreatePlayerResponse(id string) events.APIGatewayProxyResponse {
	resp := &SuccessfulCreateResponse{
		ID:     id,
		Status: "Success",
	}

	return formatResponse(resp, http.StatusOK)
}

func CreateSuccessfulGetResponse(item interface{}) events.APIGatewayProxyResponse {
	return formatResponse(item, http.StatusOK)
}

func CreateSuccesfulUpdateResponse() events.APIGatewayProxyResponse {
	return formatResponse("Success", http.StatusOK)
}

func CreateBadRequestResponse() events.APIGatewayProxyResponse {
	resp := &ErrorResponse{
		Status: "Bad Request",
	}

	return formatResponse(resp, http.StatusBadRequest)
}

func CreateResourceNotFoundResponse() events.APIGatewayProxyResponse {
	resp := &ErrorResponse{
		Status: "Resource Not Found",
	}
	return formatResponse(resp, http.StatusNotFound)
}

func CreateInternalServerErrorResponse() events.APIGatewayProxyResponse {
	resp := &ErrorResponse{
		Status: "Internal Server Error",
	}

	return formatResponse(resp, http.StatusInternalServerError)
}
