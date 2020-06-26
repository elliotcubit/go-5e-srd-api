package go5e

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
)

// Get the bytes for an API request response
func doRequest(query string) ([]byte, error) {
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

// Return a usable struct
func doRequestAndUnmarshal(query string, v interface{}) error {
	response, err := doRequest(baseURL + query)
	if err != nil {
		return err
	}
	json.Unmarshal(response, v)
	return nil
}

// Do not prepend url
func doRequestRawAndUnmarshal(url string, v interface{}) error {
  response, err := doRequest(url)
  if err != nil {
    return err
  }
  json.Unmarshal(response, v)
  return nil
}
