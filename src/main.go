package main

import (
	"encoding/json"
	"net/http"
	receipt_processor "receipt_manager/point_calculator"
	receipt "receipt_manager/receipt"
	receipt_validator "receipt_manager/receipt_validator"
	response_handler "receipt_manager/response_handler"

	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receiptMap = make(map[string]receipt.Receipt)

func newReceiptHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response_handler.HandleMethodNotAllowed(response)
		return 
	}

	newReceipt := receipt.Receipt{}

	decoderError := json.NewDecoder(request.Body).Decode(&newReceipt)
	if decoderError != nil {
		response_handler.HandleBadRequestError(response, "Receipt data decoding failed")
		return
	}

	receiptMissingFields := receipt_validator.ReceiptMissingFields(newReceipt)
	if receiptMissingFields {
		response_handler.HandleBadRequestError(response, "Receipt is missing required data fields")
		return
	}

	receiptValid := receipt_validator.ReceiptFieldsValid(newReceipt)
	if !receiptValid {
		response_handler.HandleBadRequestError(response, "Receipt data has invalid field(s)")
		return
	}

	id := uuid.New().String()
	receiptMap[id] = newReceipt
	response_handler.SendIdResponse(id, response)
}

func getPointsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response_handler.HandleMethodNotAllowed(response)
		return 
	}

	id := mux.Vars(request)["id"]
	receipt, receiptExists := receiptMap[id]
	if !receiptExists {
		response_handler.HandleNotFoundError(response, "The requested receipt doesn't exist")
		return
	}

	points, processorError := receipt_processor.ProcessReceipt(receipt)
	if processorError != nil {
		response_handler.HandleInternalServerError(response)
		return
	}

	response_handler.SendPointsResponse(points, response)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/receipts/process", newReceiptHandler)
    router.HandleFunc("/receipts/{id}/points", getPointsHandler)
    
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
