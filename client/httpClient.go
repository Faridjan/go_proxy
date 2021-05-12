package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	Client http.Client
}

func (c *HttpClient) Init()  {
	c.Client = http.Client{Timeout: time.Duration(1) * time.Second}
}

func (c *HttpClient) Get(URL string, header map[string]interface{}) (string, int) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	addHeaders(req, header)

	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyResponse, _ := ioutil.ReadAll(resp.Body)

	return string(bodyResponse), resp.StatusCode
}

func (c *HttpClient) Post(URL string, header, body map[string]interface{}) (string, int) {

	bodyJson, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(bodyJson))
	if err != nil {
		panic(err)
	}

	addHeaders(req, header)

	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyResponse,_ := ioutil.ReadAll(resp.Body)

	return string(bodyResponse), resp.StatusCode
}

func addHeaders(req *http.Request, header map[string]interface{})  {
	req.Header.Add("Accept", `application/json`)

	for k, v := range header {
		req.Header.Add(k, v.(string))
	}
}