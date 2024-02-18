package receipt_manager

import (
	receipt "receipt_manager/receipt"
	"strconv"
)

func StringIsInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func ReceiptMissingFields(receipt receipt.Receipt) bool {
	// Let's assume that all receipt fields are required
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
	// Since the decoder worked, we can assume the retailer is a string
	return true
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
	// Since the decodor worked, we can assume that items is a list of Item structs
	return true
}

func TotalValid(receipt receipt.Receipt) bool {
	if len(receipt.Total) >= 4 &&
	    StringIsInt(receipt.Total[:len(receipt.Total)-3]) &&
	    receipt.Total[len(receipt.Total)-3:len(receipt.Total)-2] == "." &&
		StringIsInt(receipt.Total[len(receipt.Total)-2:]) {
			return true
	}
	return false
}
