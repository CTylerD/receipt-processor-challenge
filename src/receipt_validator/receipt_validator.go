package receipt_manager

import (
	"fmt"
	receipt "receipt_manager/receipt"
	"regexp"
	"strconv"
)

func StringIsInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func ReceiptMissingFields(receipt receipt.Receipt) bool {
	if receipt.Retailer == "" ||
	   receipt.PurchaseDate == "" ||
	   receipt.PurchaseTime == "" ||
	   len(receipt.Items) == 0 ||
	   receipt.Total == "" {
		return true
	}
	return false
}

func ReceiptFieldsValid(receipt receipt.Receipt) bool {
	return RetailerValid(receipt) &&
		  PurchaseDateValid(receipt) &&
		  PurchaseTimeValid(receipt) &&
		  ItemsValid(receipt) &&
		  TotalValid(receipt)
}

func RetailerValid(receipt receipt.Receipt) bool {
	pattern := "^[\\w\\s\\-]+$"
	retailerIsValid, _ := regexp.MatchString(pattern, receipt.Retailer)
	return retailerIsValid
}

func PurchaseDateValid(receipt receipt.Receipt) bool {
	if len(receipt.PurchaseDate) == 10 &&
	   StringIsInt(receipt.PurchaseDate[0:4]) &&
	   receipt.PurchaseDate[4:5] == "-" &&
	   StringIsInt(receipt.PurchaseDate[5:7]) &&
	   receipt.PurchaseDate[7:8] == "-" &&
	   StringIsInt(receipt.PurchaseDate[8:10]) {
		return true
	}
	return false
}

func PurchaseTimeValid(receipt receipt.Receipt) bool {
	// Unclear if AM timestamps are written as 07:30 or 7:30, so will handle both for now
	if (len(receipt.PurchaseTime) == 5 &&
	    StringIsInt(receipt.PurchaseTime[0:2]) &&
	    receipt.PurchaseTime[2:3] == ":" &&
	    StringIsInt(receipt.PurchaseTime[3:5])) ||
	   (len(receipt.PurchaseTime) == 4 &&
		StringIsInt(receipt.PurchaseTime[0:1]) &&
	    receipt.PurchaseTime[1:2] == ":" &&
	    StringIsInt(receipt.PurchaseTime[2:4])) {
		return true
	}
	return false
}

func ItemsValid(receipt receipt.Receipt) bool {
	for _, item := range receipt.Items {
		if !validateItemDescription(item.ShortDescription) || !validateItemPrice(item.Price) {
			return false
		}
	}
	return true
}

func validateItemDescription(description string) bool {
	pattern := "^[\\w\\s\\-]+$"
	descriptionIsValid, err := regexp.MatchString(pattern, description)
	if err != nil {
		fmt.Println("Error while validating item description:", err)
		return false
	}
	return descriptionIsValid
}

func validateItemPrice(price string) bool {
	pattern := "^\\d+\\.\\d{2}$"
	priceIsValid, err := regexp.MatchString(pattern, price)
	if err != nil {
		fmt.Println("Error while validating item price:", err)
		return false
	}
	return priceIsValid
}

func TotalValid(receipt receipt.Receipt) bool {
	pattern := "^\\d+\\.\\d{2}$"
	totalIsValid, err := regexp.MatchString(pattern, receipt.Total)
	if err != nil {
		fmt.Println("Error while validating total price:", err)
		return false
	}
	return totalIsValid
}
