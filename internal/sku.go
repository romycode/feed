package internal

import (
	"fmt"
	"math/rand"
	"regexp"
)

// SKU struct to work with SKUs
type SKU struct {
	value string
}

// NewSKU generates a valid sku
func NewSKU() SKU {
	return SKU{
		value: fmt.Sprintf("%s-%s%s%s%s", randomLetter()+randomLetter()+randomLetter()+randomLetter(), randomInt(), randomInt(), randomInt(), randomInt()),
	}
}

// NewSKUFromString validate and return the sku for the given string
func NewSKUFromString(sku string) (SKU, error) {
	if valid := validateSKU(sku); !valid {
		return SKU{}, fmt.Errorf("%s is invalid sku", sku)
	}
	return SKU{
		value: sku,
	}, nil
}

func (s SKU) Value() string {
	return s.value
}

// validateSKU validate the string following the requirements defined in the guides
func validateSKU(input string) bool {
	re, _ := regexp.Compile(`^[a-zA-Z]{4}-[0-9]{4}$`)
	valid := re.MatchString(input)

	return valid
}

// randomLetter return a random letter
func randomLetter() string {
	var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	return string(letters[rand.Intn(len(letters))])
}

// randomInt return a random letter
func randomInt() string {
	var numbers = []byte("0123456789")
	return string(numbers[rand.Intn(len(numbers))])
}
