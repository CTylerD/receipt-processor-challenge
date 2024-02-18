package receipt_manager

import (
	"errors"
	"math"
	receipt "receipt_manager/receipt"
	"strconv"
	"strings"
	"unicode"
)

func RetailerNamePoints(receipt receipt.Receipt) (int, error) {
	// One point for every alphanumeric character in the retailer name
	addedPoints := 0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			addedPoints++
		}
	}
	return addedPoints, nil
}

func RoundDollarAmountPoints(receipt receipt.Receipt) (int, error) {
	// 50 points if the total is a round dollar amount with no cents
	if len(receipt.Total) >= 4 && receipt.Total[len(receipt.Total)-3:] == ".00" {
		return 50, nil
	}
	return 0, nil
}

func MultipleOfQuarterPoints(receipt receipt.Receipt) (int, error) {
	// 25 points if the total is a multiple of `0.25`
	totalInt := receipt.Total
	totalFloat, err := strconv.ParseFloat(totalInt, 64)

	if err != nil {
		return -1, errors.New("strconv.ParseFloat() in multipleOfQuarterPoints() failed")
	}

	if math.Mod(totalFloat / .25, 1.0) == 0 {
		return 25, nil
	}
	return 0, nil
}

func EveryTwoItemsPoints(receipt receipt.Receipt) (int, error) {
	// 5 points for every two items on the receipt
	return (len(receipt.Items) / 2) * 5, nil
}

func DescriptionLengthPoints(receipt receipt.Receipt) (int, error) {
	/*
	If the trimmed length of the item description is a multiple of 3,
	multiply the price by `0.2` and round up to the nearest integer
	The result is the number of points earned
	*/
	addedPoints := 0
	for _, item := range receipt.Items {
		priceFloat, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return -1, errors.New("strconv.ParseFloat() in descriptionLengthPoints() failed")
		}

		if len(strings.TrimSpace(item.ShortDescription)) % 3 == 0 {
			addedPoints += int(math.Ceil(priceFloat * .2))
		}
	}
	return addedPoints, nil
}

func OddPurchaseDatePoints(receipt receipt.Receipt) (int, error) {
	// 6 points if the day in the purchase date is odd
	var dateString string
	var dateInt int
	var err error
	if len(receipt.PurchaseDate) == 10 {
		dateString = receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
		dateInt, err = strconv.Atoi(dateString)
	}

	if len(receipt.PurchaseDate) != 10 || err != nil {
		return -1, errors.New("strconv.Atoi() in oddPurchaseDatePoints() failed")
	}

	if int(dateInt) % 2 == 1 {
		return 6, nil
	}
	return 0, nil
}

func PurchaseTimePoints(receipt receipt.Receipt) (int, error) {
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	if receipt.PurchaseTime > "14:00" && receipt.PurchaseTime < "16:00" {
		return 10, nil
	}
	return 0, nil
}

type pointCalculators func(receipt receipt.Receipt) (int, error)

func pointRuleFunctions() []pointCalculators {
	return []pointCalculators {
		RetailerNamePoints,
		RoundDollarAmountPoints,
		MultipleOfQuarterPoints, 
		EveryTwoItemsPoints,
		DescriptionLengthPoints, 
		OddPurchaseDatePoints, 
		PurchaseTimePoints}
}

func ProcessReceipt(receipt receipt.Receipt) (int, error) {
	  totalPoints := 0
	  for _, function := range pointRuleFunctions() {
		points, err := function(receipt)
		if err != nil {
			return -1, err
		}
		totalPoints += points
	}
	return totalPoints, nil
}
