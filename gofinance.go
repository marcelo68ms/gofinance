package gofinance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiHgFinancePrefix = "https://api.hgbrasil.com/finance"

func GetQuotations(key string) (*DataQuotation, error) {

	url := fmt.Sprintf("%s/quotations?key=%s", apiHgFinancePrefix, key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *DataQuotation

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetTaxes(key string) (*DataTaxe, error) {

	url := fmt.Sprintf("%s/taxes?key=%s", apiHgFinancePrefix, key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *DataTaxe

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
