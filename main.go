package simplehttp

import (
	"errors"
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

var BadStatusCode = errors.New("Bad status code in response")
func BytesFromURL(url string) (body []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		err = fmt.Errorf("Error doing http.GET on %s: %w", url, err)
		return
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("Error decoding body: %w", err)
		return
	}
	resp.Body.Close()

	if resp.StatusCode > 299 {
		err = fmt.Errorf("%w; GET '%s' gave response code (%d)", BadStatusCode, url, resp.StatusCode)
		return
	}

	return
}

func StringFromURL(url string) (s string, err error) {
	b, err := BytesFromURL(url)
	s = string(b)

	return
}

func UnmarshalToStructFromURL(url string, parsed any) (err error) {
	var data []byte
	if data, err = BytesFromURL(url); err != nil {
		return
	}
	err = json.Unmarshal(data, parsed)

	return
}

func UnmarshalToMapFromURL(url string) (dataMap map[string]any, err error) {
	var data []byte
	if data, err = BytesFromURL(url); err != nil {
		return
	}
	dataMap = make(map[string]any)
	err = json.Unmarshal(data, &dataMap)

	return
}
