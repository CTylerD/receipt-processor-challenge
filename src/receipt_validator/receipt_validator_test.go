package receipt_manager

import (
	item "receipt_manager/item"
	receipt "receipt_manager/receipt"
	rv "receipt_manager/receipt_validator"
	"testing"
)

func TestRetailerValid(test *testing.T) {
	testCases := []struct {
		receipt           receipt.Receipt
		expectedValidity  bool
	}{
		{
			receipt: receipt.Receipt{Retailer: "The Corner Store"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{Retailer: "Walmart 123"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{Retailer: "Target!"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{Retailer: "Best Buy"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{Retailer: "123 Grocery Store"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{Retailer: ""},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{Retailer: "Test*Store"},
			expectedValidity: false, // False case: contains special character
		},
		{
			receipt: receipt.Receipt{Retailer: "Supermarket!"},
			expectedValidity: false, // False case: contains special character
		},
		}

	for _, testCase := range testCases {
		retailerValidity := rv.RetailerValid(testCase.receipt)

	if retailerValidity != testCase.expectedValidity {
		test.Errorf("Retailer '%s', validity is %t, but expected %t",
		testCase.receipt.Retailer, retailerValidity, testCase.expectedValidity)
		}
	}
}

func TestPurchaseDateValid(test *testing.T) {
	testCases := []struct {
		receipt receipt.Receipt
		expectedValidity bool
	}{
		{
			receipt: receipt.Receipt{PurchaseDate: "2000-01-01"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{PurchaseDate: "2000001-01"},
			expectedValidity: false,
		},
		
		{
			receipt: receipt.Receipt{PurchaseDate: "2000-1-1"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{PurchaseDate: "today's date"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{PurchaseDate: ""},
			expectedValidity: false,
		},
	}

	for _, testCase := range testCases {
		purchaseDateValidity := rv.PurchaseDateValid(testCase.receipt)
		
		if purchaseDateValidity != testCase.expectedValidity {
			test.Errorf("Purchase Date '%s', validity is %t, but expected %t",
				testCase.receipt.PurchaseDate, purchaseDateValidity, testCase.expectedValidity)
		}
	}
}

func TestPurchaseTimeValid(test *testing.T) {
	testCases := []struct {
		receipt receipt.Receipt
		expectedValidity bool
	}{
		{
			receipt: receipt.Receipt{PurchaseTime: "12:00"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{PurchaseTime: "9:30"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{PurchaseTime: "A9:30"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{PurchaseTime: "9"},
			expectedValidity: false,
		},
				{
			receipt: receipt.Receipt{PurchaseTime: ""},
			expectedValidity: false,
		},
	}
		
	for _, testCase := range testCases {
		purchaseTimeValidity := rv.PurchaseTimeValid(testCase.receipt)
		
		if purchaseTimeValidity != testCase.expectedValidity {
			test.Errorf("Purchase Time '%s', validity is %t, but expected %t",
				testCase.receipt.PurchaseTime, purchaseTimeValidity, testCase.expectedValidity)
		}
	}
}

func TestValidateItems(test *testing.T) {
	testCases := []struct {
		receipt          receipt.Receipt
		expectedValidity bool
	}{
		{
			receipt: receipt.Receipt{ 
				Items: []item.Item{
					{ShortDescription: "400 fish sticks", Price: "6.49"},
					{ShortDescription: "Coca Cola 6pack", Price: "5.99"},
				},
			},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{ 
				Items: []item.Item{
					{ShortDescription: "4,000 fish sticks", Price: "10.00"},
					{ShortDescription: "Coca Cola 6-pack", Price: "5.99"},
				},
			},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{
				Items: []item.Item{
					{ShortDescription: "400 fish sticks", Price: "6.49"},
					{ShortDescription: "Coca Cola 6PK", Price: "invalid_price"},
				},
			},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{
				Items: []item.Item{
					{ShortDescription: "", Price: "6.49"},
					{ShortDescription: "Coca Cola 6PK", Price: "5.99"},
				},
			},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{
				Items: []item.Item{
					{ShortDescription: "400 fish sticks", Price: ""},
					{ShortDescription: "Coca Cola 6PK", Price: "5.99"},
				},
			},
			expectedValidity: false,
		},
	}

	for _, testCase := range testCases {
		result := rv.ItemsValid(testCase.receipt)
		if result != testCase.expectedValidity {
			test.Errorf("Validation result mismatch for items %+v, expected %t but got %t", testCase.receipt.Items, testCase.expectedValidity, result)
		}
	}
}

func TestTotalValid(test *testing.T) {
	testCases := []struct {
		receipt receipt.Receipt
		expectedValidity bool
	}{
		{
			receipt: receipt.Receipt{Total: "50.00"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{Total: "0.00"},
			expectedValidity: true,
		},
		{
			receipt: receipt.Receipt{Total: "10.5"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{Total: "50"},
			expectedValidity: false,
		},
		{
			receipt: receipt.Receipt{Total: ""},
			expectedValidity: false,
		},
	}
		
	for _, testCase := range testCases {
		totalValidity := rv.TotalValid(testCase.receipt)
		
		if totalValidity != testCase.expectedValidity {
			test.Errorf("Total '%s', validity is %t, but expected %t",
				testCase.receipt.Total, totalValidity, testCase.expectedValidity)
		}
	}
}
