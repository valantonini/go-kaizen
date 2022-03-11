package web_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ExampleJsonPost() {
	jsonData := []byte("{ foo: bar }")

	req, _ := http.NewRequest("POST", "http://example.com", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	type coffee struct {
		name string
	}
	c := coffee{}
	err := json.Unmarshal(body, &c)

	if err != nil {
		_ = fmt.Errorf("error performing post %v", err)
	}
}
