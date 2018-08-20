package go_currency_codes

import (
	"encoding/json"
	"errors"
	"crypto/sha256"
	"fmt"
	"strings"
	"io"
)

// All returns all the currencies
func All() (map[string]interface{}, error) {

	currencyMap, err := parse()
	if err != nil {
		return nil, err
	}

	return currencyMap, nil
}

// FromCurrencyName returns the currency details from the currency name
func FromCurrencyName(name string) (map[string]interface{}, error) {

	if name == "" {
		return nil, errors.New("empty currency name")
	}

	sha := sha256.New()
	sha.Write([]byte(strings.TrimSpace(name)))
	shaName := fmt.Sprintf("%x", sha.Sum(nil))

	currencyMap, err := parse()
	if err != nil {
		return nil, err
	}

	detail, ok := currencyMap[shaName]
	if !ok {
		return nil, errors.New("currency name not found")
	}

	return detail.(map[string]interface{}), nil
}

// WriteJSON writes all the currencies to the given writer as a JSON string
func WriteJSON(indent int, writer io.Writer) error  {

	currencyMap, err := parse()
	if err != nil {
		return err
	}

	enc := json.NewEncoder(writer)
	enc.SetIndent("", strings.Repeat(" ", indent))
	return enc.Encode(currencyMap)
}

// parse encodes the world's currencies into a map
func parse() (map[string]interface{}, error) {

	var currencyMap map[string]interface{}

	err := json.Unmarshal([]byte(Data), &currencyMap)
	if err != nil {
		return nil, err
	}

	return currencyMap, nil
}