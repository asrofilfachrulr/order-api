package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseJSONToInterface(d io.ReadCloser, m interface{}) (interface{}, error) {
	jsonByte, err := ioutil.ReadAll(d)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonByte, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
