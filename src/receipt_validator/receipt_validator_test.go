package receipt_manager

import (
	receipt "receipt_manager/receipt"
	"testing"
)

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
        purchaseDateValidity := PurchaseDateValid(testCase.receipt)
        
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
			expectedValidity: true,
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
        purchaseTimeValidity := PurchaseTimeValid(testCase.receipt)
        
        if purchaseTimeValidity != testCase.expectedValidity {
            test.Errorf("Purchase Time '%s', validity is %t, but expected %t",
                testCase.receipt.PurchaseTime, purchaseTimeValidity, testCase.expectedValidity)
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
        totalValidity := TotalValid(testCase.receipt)
        
        if totalValidity != testCase.expectedValidity {
            test.Errorf("Total '%s', validity is %t, but expected %t",
                testCase.receipt.Total, totalValidity, testCase.expectedValidity)
        }
    }
}
