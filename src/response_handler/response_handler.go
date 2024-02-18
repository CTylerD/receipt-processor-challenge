package receipt_manager

import (
	"encoding/json"
	"net/http"
)

type IdResponse struct {
	Id string `json:"id"`
}

func SendIdResponse(id string, response http.ResponseWriter) {
	responseStruct := IdResponse {
		Id: id,
	}
	sendHttpResponse(responseStruct, response)
}

type PointsResponse struct {
	Points int `json:"points"`
}

func SendPointsResponse(points int, response http.ResponseWriter) {
	responseStruct := PointsResponse {
		Points: points,
	}
	sendHttpResponse(responseStruct, response)
}

func sendHttpResponse(responseStruct interface{}, response http.ResponseWriter) {
	responseBody, err := json.Marshal(responseStruct)
		if err != nil {
			HandleInternalServerError(response)
			return
		}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	response.Write(responseBody)
}

func HandleBadRequestError(response http.ResponseWriter, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "The request is invalid"
	}
	handleClientError(response, errorMsg, http.StatusBadRequest)
}

func HandleNotFoundError(response http.ResponseWriter, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "The request resource was not found"
	}
	handleClientError(response, errorMsg, http.StatusNotFound)
}

func HandleMethodNotAllowed(response http.ResponseWriter) {
	handleClientError(response,
		"This method is not allowed on this endpoint",
		http.StatusMethodNotAllowed)
}

func HandleInternalServerError(response http.ResponseWriter) {
	handleClientError(response,
		"Internal Server Error", 
		http.StatusInternalServerError)
}

func handleClientError(response http.ResponseWriter, errorMsg string, statusCode int) {
	errorResponse := map[string]string{"Error": errorMsg}

	jsonResponse, err := json.Marshal(errorResponse)
	if err != nil {
		HandleInternalServerError(response)
		return
	}

	response.WriteHeader(statusCode)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
