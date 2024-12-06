package currency

import (
	"fmt"
	"math/big"
)

type Converter struct {
	conversionRates map[string]map[string]*big.Rat
}

func NewConverter(conversionRates map[string]map[string]*big.Rat) (Converter, error) {
	if len(conversionRates) == 0 {
		return Converter{}, fmt.Errorf("rates cannot be nil or empty")
	}
	return Converter{
		conversionRates: conversionRates,
	}, nil
}

func (c Converter) Convert(from string, to string, amount *big.Rat) (
	*big.Rat,
	error,
) {
	toRate, ok := c.conversionRates[from]
	if !ok {
		return nil, fmt.Errorf("from: %s currency not found", from)
	}
	rate, ok := toRate[to]
	if !ok {
		return nil, fmt.Errorf("to: %s currency not found", to)
	}
	result := new(big.Rat).Mul(amount, rate)
	return result, nil
}
