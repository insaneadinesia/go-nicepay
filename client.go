package nicepay

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	APIEnvType  Env
	MID         string
	MKey        string
	NotifURL    string
	CallbackURL string
}

func NewCient() Client {
	return Client{
		APIEnvType: Sandbox,
	}
}

var defHTTPTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

func (c *Client) CallRequest(method, path string, headers []Header, body io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		fmt.Println("Request creation failed : ", err)
		return err
	}

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	start := time.Now()

	fmt.Println("Request ", method, ": ", path)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Cannot send request : ", err)
		return err
	}
	fmt.Println("Completed in ", time.Since(start))
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Cannot read response body: ", err)
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
