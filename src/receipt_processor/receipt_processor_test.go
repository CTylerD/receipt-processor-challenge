package receipt_manager

import (
	"errors"
	item "receipt_manager/item"
	receipt "receipt_manager/receipt"
	"testing"
)

func TestRetailerNamePoints(test *testing.T) {
    testCases := []struct {
        receipt receipt.Receipt
        expectedPoints int
		expectedErr error
    }{
        {
            receipt: receipt.Receipt{Retailer: "The Stuff Store"},
            expectedPoints: 13,
			expectedErr: nil,
        },
        {
            receipt: receipt.Receipt{Retailer: "22 Aghast!?!?!"},
            expectedPoints: 8,
			expectedErr: nil,
		},
		{
            receipt: receipt.Receipt{Retailer: ""},
            expectedPoints: 0,
			expectedErr: nil,
   	 	},
		{
            receipt: receipt.Receipt{Retailer: "(@*& #$)(* @)#$(*/)"},
            expectedPoints: 0,
			expectedErr: nil,
   	 	},
	}
		
    for _, testCase := range testCases {
        actualPoints, err := RetailerNamePoints(testCase.receipt)
        
        if err != nil {
            test.Errorf("Test failed with error: %v", err)
        }
        
        if actualPoints != testCase.expectedPoints {
            test.Errorf("Retailer '%s', got %d points, but expected %d",
                testCase.receipt.Retailer, actualPoints, testCase.expectedPoints)
        }
    }
}

func TestRoundDollarAmountPoints(test *testing.T) {
    testCases := []struct {
        receipt receipt.Receipt
        expectedPoints int
        expectedErr error
    }{
        {
            receipt: receipt.Receipt{Total: "100.00"},
            expectedPoints: 50,
            expectedErr: nil,
        },
		{
            receipt: receipt.Receipt{Total: "-100.00"},
            expectedPoints: 50,
            expectedErr: nil,
        },
		{
            receipt: receipt.Receipt{Total: "200.50"},
            expectedPoints: 0,
            expectedErr: nil,
        },
		{
            receipt: receipt.Receipt{Total: "150.75"},
            expectedPoints: 0,
            expectedErr: nil,
        },
		{
            receipt: receipt.Receipt{Total: ""},
            expectedPoints: 0,
            expectedErr: nil,
        },
    }
		
    for _, testCase := range testCases {
        actualPoints, err := RoundDollarAmountPoints(testCase.receipt)
        
        if testCase.expectedErr == nil && err != nil {
            test.Errorf("Test failed with error: %v", err)
        }
        
        if actualPoints != testCase.expectedPoints {
            test.Errorf("Total '%s', got %d points, but expected %d",
                testCase.receipt.Total, actualPoints, testCase.expectedPoints)
        }
    }
}

func TestMultipleOfQuarterPoints(test *testing.T) {
	testCases := []struct {
        receipt receipt.Receipt
        expectedPoints int
        expectedErr error
    }{
        {
            receipt: receipt.Receipt{Total: "1.50"},
            expectedPoints: 25,
            expectedErr: nil,
        },
        {
            receipt: receipt.Receipt{Total: "1.78"},
            expectedPoints: 0,
            expectedErr: nil,
        },
		{
            receipt: receipt.Receipt{Total: "0.00"},
            expectedPoints: 25,
            expectedErr: nil,
        },
        {
            receipt: receipt.Receipt{Total: "-2.00"},
            expectedPoints: 25,
            expectedErr: nil,
        },
        {
            receipt: receipt.Receipt{Total: "abc"},
            expectedPoints: -1,
            expectedErr: errors.New("strconv.ParseFloat() in multipleOfQuarterPoints() failed"),
        },
        {
            receipt: receipt.Receipt{Total: ""},
            expectedPoints: -1,
            expectedErr: errors.New("strconv.ParseFloat() in multipleOfQuarterPoints() failed"),
        },
    }

	for _, testCase := range testCases {
        actualPoints, err := MultipleOfQuarterPoints(testCase.receipt)
        
        if testCase.expectedErr == nil && err != nil {
            test.Errorf("Test failed with error: %v", err)
        }
        
        if actualPoints != testCase.expectedPoints {
            test.Errorf("Total '%s', got %d points, but expected %d",
                testCase.receipt.Total, actualPoints, testCase.expectedPoints)
        }
    }
}

func TestEveryTwoItemsPoints(test *testing.T) {
	testCases := []struct {
        receipt receipt.Receipt
        expectedPoints int
        expectedErr error
    }{
        {
			receipt: receipt.Receipt{
				Items: []item.Item{
					{ShortDescription: "Item1", Price: "10.00"},
					{ShortDescription: "Item2", Price: "15.00"},
					{ShortDescription: "Item3", Price: "20.00"},
					{ShortDescription: "Item4", Price: "25.00"},
				},
			},
			expectedPoints: 10,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{
				Items: []item.Item{
					{ShortDescription: "Item1", Price: "10.00"},
					{ShortDescription: "Item2", Price: "15.00"},
					{ShortDescription: "Item3", Price: "20.00"},
				},
			},
			expectedPoints: 5,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{
				Items: []item.Item{
					{ShortDescription: "Item1", Price: "10.00"},
				},
			},
			expectedPoints: 0,
			expectedErr: nil,
		},
			{
			receipt: receipt.Receipt{Items: []item.Item{}},
			expectedPoints: 0,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
        actualPoints, err := EveryTwoItemsPoints(testCase.receipt)
        
        if testCase.expectedErr == nil && err != nil {
            test.Errorf("Test failed with error: %v", err)
        }
        
        if actualPoints != testCase.expectedPoints {
            test.Errorf("Items '%s', got %d points, but expected %d",
                testCase.receipt.Items, actualPoints, testCase.expectedPoints)
        }
    }
}

func TestDescriptionLengthPoints(test *testing.T) {
	testCases := []struct {
		receipt receipt.Receipt
		expectedPoints int
		expectedErr error
	}{
		{
			receipt: receipt.Receipt{Items: []item.Item{
				{Price: "10.00", ShortDescription: "Apple"},
				{Price: "15.00", ShortDescription: "Banana"},
				{Price: "20.00", ShortDescription: "Grapefruit"},
			}},
			expectedPoints: 3,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{Items: []item.Item{
				{Price: "10.00", ShortDescription: "Apple"},
				{Price: "15.00", ShortDescription: "Pear"},
				{Price: "20.00", ShortDescription: "Grapefruit"},
			}},
			expectedPoints: 0,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{Items: []item.Item{
				{Price: "1.00", ShortDescription: "Hat"},
			}},
			expectedPoints: 1,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{Items: []item.Item{}},
			expectedPoints: 0,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{Items: []item.Item{
				{Price: "-10.00", ShortDescription: "Apple"},
				{Price: "-15.00", ShortDescription: "Banana"},
			}},
			expectedPoints: -3,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		actualPoints, err := DescriptionLengthPoints(testCase.receipt)
        
        if testCase.expectedErr == nil && err != nil {
            test.Errorf("Test failed with error: %v", err)
        }
        
        if actualPoints != testCase.expectedPoints {
            test.Errorf("Items '%s', got %d points, but expected %d",
                testCase.receipt.Items, actualPoints, testCase.expectedPoints)
        }
	}
}

func TestOddPurchaseDatePoints(test *testing.T) {
	testCases := []struct {
		receipt receipt.Receipt
		expectedPoints int
		expectedErr error
	}{
		{
			receipt: receipt.Receipt{PurchaseDate: "2024-02-15"},
			expectedPoints: 6,
			expectedErr: nil,
		},
		{
			receipt: receipt.Receipt{PurchaseDate: "asdf-02-16"},
			expectedPoints: 0,
			expectedErr: errors.New("strconv.Atoi() in oddPurchaseDatePoints() failed"),
		},
		{
			receipt: receipt.Receipt{PurchaseDate: ""},
			expectedPoints: -1,
			expectedErr: errors.New("strconv.Atoi() in oddPurchaseDatePoints() failed"),
		},
		{
			receipt: receipt.Receipt{PurchaseDate: "2024-02-"},
			expectedPoints: -1,
			expectedErr: errors.New("strconv.Atoi() in oddPurchaseDatePoints() failed"),
		},
		{
			receipt: receipt.Receipt{PurchaseDate: "2024-02-5"},
			expectedPoints: -1,
			expectedErr: errors.New("strconv.Atoi() in oddPurchaseDatePoints() failed"),
		},
	}

	for _, testCase := range testCases {
		actualPoints, err := OddPurchaseDatePoints(testCase.receipt)
        
        if testCase.expectedErr == nil && err != nil {
            test.Errorf("Test failed with error: %v", err)
        }
        
        if actualPoints != testCase.expectedPoints {
            test.Errorf("Purchase date '%s', got %d points, but expected %d",
                testCase.receipt.PurchaseDate, actualPoints, testCase.expectedPoints)
        }
	}
}