package github.com/ddelellis-pkg/simplehttp

import (
	"errors"
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

func main() {
	data, err := BytesFromURL("https://api.sunrise-sunset.org/json?lat=36.7201600&lng=-122.007&date=tomorrow&formatted=0")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}

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

func GetURLAndUnmarshalToStruct(url string, parsed *any) (err error) {
	var data []byte
	if data, err = BytesFromURL(url); err != nil {
		return
	}
	err = json.Unmarshal(data, parsed)

	return
}

func GetURLAndUnmarshalToMap(url string) (dataMap map[string]any, err error) {
	var data []byte
	if data, err = BytesFromURL(url); err != nil {
		return
	}
	dataMap = make(map[string]any)
	err = json.Unmarshal(data, &dataMap)

	return
}
