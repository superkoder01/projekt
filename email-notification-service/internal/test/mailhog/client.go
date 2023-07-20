package mailhog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMessages(url string) ([]Messages, error) {
	response, err := http.Get(url + "/api/v1/messages")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("cannot get messages. response.StatusCode: %v", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	m := make([]Messages, 0)
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
