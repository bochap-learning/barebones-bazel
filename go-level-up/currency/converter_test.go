package currency

import (
	"math/big"
	"reflect"
	"testing"
)

func getValidConverter() (Converter, error) {
	rates := map[string]map[string]*big.Rat{
		"USD": {"JPY": big.NewRat(100, 1)},
		"JPY": {"USD": big.NewRat(1, 100)},
	}
	return NewConverter(rates)
}

func TestValidNewConverter(t *testing.T) {
	target, err := getValidConverter()
	if err != nil {
		t.Fatalf("expect error: nil, got: %v", err)
	}
	if reflect.DeepEqual(target, Converter{}) {
		t.Fatal("expect converter: non nil, got: nil")
	}
}

func TestInvalidNewConverter(t *testing.T) {
	testCases := []struct {
		description string
		rates       map[string]map[string]*big.Rat
	}{
		{
			description: "nil rates",
			rates:       nil,
		},
		{
			description: "empty rates",
			rates:       map[string]map[string]*big.Rat{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			target, err := NewConverter(testCase.rates)
			if err == nil {
				t.Fatal("expect error: non nil, got: nil")
			}
			if !reflect.DeepEqual(target, Converter{}) {
				t.Fatalf("expect converter: nil, got: %v", target)
			}
		})
	}
}

func TestValidConvert(t *testing.T) {
	target, _ := getValidConverter()
	testCases := []struct {
		description string
		from        string
		to          string
		amount      *big.Rat
		expect      *big.Rat
	}{
		{
			description: "from: USD, to: JPY, amount: 1",
			from:        "USD",
			to:          "JPY",
			amount:      big.NewRat(1, 1),
			expect:      big.NewRat(100, 1),
		},
		{
			description: "from: JPY, to: USD, amount: 100",
			from:        "JPY",
			to:          "USD",
			amount:      big.NewRat(100, 1),
			expect:      big.NewRat(1, 1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got, err := target.Convert(testCase.from, testCase.to, testCase.amount)
			if err != nil {
				t.Fatalf("expect error: nil, got: %v", err)
			}
			if !reflect.DeepEqual(testCase.expect, got) {
				t.Fatalf("expect: %v, got: %v to be equal", testCase.expect, got)
			}
		})
	}
}

func TestInvalidConvert(t *testing.T) {
	target, _ := getValidConverter()
	testCases := []struct {
		description string
		from        string
		to          string
		amount      *big.Rat
	}{
		{
			description: "from: EUR, to: JPY, amount: 1 to fail",
			from:        "EUR",
			to:          "JPY",
			amount:      big.NewRat(1, 1),
		},
		{
			description: "from: JPY, to: EUR, amount: 100 to fail",
			from:        "JPY",
			to:          "EUR",
			amount:      big.NewRat(100, 1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got, err := target.Convert(testCase.from, testCase.to, testCase.amount)
			if err == nil {
				t.Fatal("expect error: non nil, got: nil")
			}
			if got != nil {
				t.Fatalf("expect: nil, got: %v", got)
			}
		})
	}
}
