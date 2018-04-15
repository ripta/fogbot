package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultEndpointPath = "/f/api/0/jsonapi"
	defaultHostname     = "%s.fogbugz.com"
)

type Client struct {
	Endpoint *url.URL
	Token    Token

	statusCache []*StatusDescription
	tagCache    []*Tag
}

type Request map[string]interface{}

type Response struct {
	Data   interface{}      `json:"data"`
	Errors []*ResponseError `json:"errors"`
}

type Token string

func New(hostname string) (*Client, error) {
	if hostname == "" {
		return nil, fmt.Errorf("Hostname must not be empty")
	}
	if !strings.Contains(hostname, ".") {
		hostname = fmt.Sprintf(defaultHostname, hostname)
	}
	e := &url.URL{
		Scheme: "https",
		Host:   hostname,
		Path:   defaultEndpointPath,
	}
	c := &Client{
		Endpoint: e,
	}
	return c, nil
}

func (c *Client) Request(cmd string, req Request) ([]byte, error) {
	req["cmd"] = cmd
	if c.Token != "" {
		req["token"] = c.Token
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(payload)
	rsp, err := http.Post(c.Endpoint.String(), "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	// log.Printf("Payload: %s\n", payload)
	// log.Printf("Data: %s\n", body)
	// if err := ioutil.WriteFile("/var/folders/cd/74s1hr4j2tvcjcn1kwhh52xc0000gn/T/tmp.FJ7WrmHS", body, 0644); err != nil {
	// 	return nil, err
	// }

	var out Response
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	if out.Errors != nil && len(out.Errors) > 0 {
		return nil, out.Errors[0]
	}
	return json.Marshal(out.Data)
}

func BuildRequest(v interface{}) (Request, error) {
	p, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	req := Request{}
	if err = json.Unmarshal(p, &req); err != nil {
		return req, err
	}
	return req, nil
}
