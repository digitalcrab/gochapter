package request

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

// ReadBody reads body of request and closes reader
func ReadBody(body io.ReadCloser) ([]byte, error) {
	if nil == body {
		return nil, errors.New("body is empty")
	}
	defer body.Close()
	return ioutil.ReadAll(body)
}

// ParseBody reads and parses body
func ParseBody(body io.ReadCloser, v interface{}) error {
	data, err := ReadBody(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
