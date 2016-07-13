package paypal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Client struct {
	User        string `url:"USER"`
	Password    string `url:"PWD"`
	Signature   string `url:"SIGNATURE"`
	Version     string `url:"VERSION"`
	endpoint    string
	redirectURL string
}

// Request makes a request to Paypal's nvp endpoint
// data is the struct being used
// It returns (*url.Values, error)
func (c *Client) Request(data interface{}) (*url.Values, error) {
	apiCallValues, _ := query.Values(data)
	resp, err := http.PostForm(c.endpoint, apiCallValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	resValues, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}
	if resValues.Get("ACK") == "Failure" {
		return nil, errors.New(fmt.Sprintf("Payment failure - %s : %s", resValues.Get("L_SHORTMESSAGE0"), resValues.Get("L_LONGMESSAGE0")))
	}
	return &resValues, nil
}

func (c *Client) SetEnvironment(env string) {
	if env == EnvSandbox {
		c.endpoint = SandboxEndpoint
		c.redirectURL = SandboxRedirectURL
		return
	}
	if env == EnvProduction {
		c.endpoint = ProductionEndpoint
		c.redirectURL = ProductionRedirectURL
		return
	}
}
