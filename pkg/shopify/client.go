package shopify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FetchProductData(shopifyBaseUrl string) ([]Product, error) {
	resp, err := http.Get(shopifyBaseUrl + "/products.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var productsResponse ProductsResponse
	err = json.Unmarshal(body, &productsResponse)
	if err != nil {
		return nil, err
	}

	return productsResponse.Products, nil
}