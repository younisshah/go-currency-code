package go_currency_codes

import (
	"testing"
	"os"
)

func TestAll(t *testing.T) {
	currencyMap, err := All()

	if err != nil {
		t.Fatal(err)
	}

	actualLength := len(currencyMap)

	if actualLength != 118 {
		t.Error("expected 118, got", actualLength)
	}
}

func TestFromCurrencyName(t *testing.T) {
	detail, err := FromCurrencyName("US Dollar")

	if err != nil {
		t.Fatal(err)
	}

	actual := detail["symbol"]
	if  actual != "$" {
		t.Error("expected $, got", actual)
	}
}

func TestWriteJSON(t *testing.T) {

	err := WriteJSON(2, os.Stdout)

	if err != nil {
		t.Fatal(err)
	}
}