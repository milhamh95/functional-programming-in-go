package f_pure_hotdog_shop

import (
	"errors"
	"testing"
)

var (
	testChangeStruct = []struct {
		inputCard  CreditCard
		amount     int
		outputCard CreditCard
		err        CreditError
	}{
		{
			inputCard:  CreditCard{credit: 1000},
			amount:     500,
			outputCard: CreditCard{credit: 500},
			err:        nil,
		},
	}
)

func TestCharge(t *testing.T) {
	for _, test := range testChangeStruct {
		t.Run("", func(t *testing.T) {
			output, err := Charge(test.inputCard, test.amount)
			if output != test.outputCard || !errors.Is(err, test.err) {
				t.Errorf("expected %v but got %v \n, error expected %v but got %v",
					test.outputCard, output, test.err, err)
			}
		})
	}
}

func TestOrderHotdog(t *testing.T) {
	testCC := CreditCard{credit: 1000}
	calledInnerFunction := false
	mockPayment := func(c CreditCard, input int) (CreditCard, CreditError) {
		calledInnerFunction = true
		testCC.credit -= input
		return testCC, nil
	}

	hotdog, resultF := OrderHotdog(testCC, mockPayment)
	if hotdog != NewHotdog() {
		t.Errorf("expected %v but got %v\n", NewHotdog(), hotdog)
	}

	_, err := resultF()
	if err != nil {
		t.Errorf("encountered %v but expected no error\n", err)
	}

	if calledInnerFunction == false {
		t.Errorf("innter function did not get called\n")
	}
}
