package gofinance

import "testing"

const key = "6c050039"

func TestGetQuotations(t *testing.T) {

	_, err := GetQuotations(key)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetTaxes(t *testing.T) {

	_, err := GetTaxes(key)
	if err != nil {
		t.Error(err)
		return
	}

}
