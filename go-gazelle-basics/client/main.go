package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/bochap-learning/barebones-bazel/go-gazelle-basics/currency"
)

func main() {
	from := flag.String("from", "USD", "From currency")
	to := flag.String("to", "JPY", "To currency")
	amount := flag.Float64("amount", 1, "Amount")
	flag.Parse()
	rates := map[string]map[string]*big.Rat{
		"USD": {"JPY": big.NewRat(100, 1)},
		"JPY": {"USD": big.NewRat(1, 100)},
	}
	converter, err := currency.NewConverter(rates)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	ratAmount := new(big.Rat)
	ratAmount.SetFloat64(*amount)
	converted, err := converter.Convert(*from, *to, ratAmount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	result, isPrecised := converted.Float64()
	fmt.Printf("Convert amount: %f, from: %s, to: %s returns %f and %t\n", *amount, *from, *to, result, isPrecised)
}
