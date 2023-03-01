package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://economia.awesomeapi.com.br/json/%s-USD"

type QuotationAPIResp struct {
	Code      string `json:"code"`
	CodeIn    string `json:"codein"`
	Name      string `json:"name"`
	Ask       string `json:"ask"`
	UpdatedAt string `json:"timestamp"`
}

type QuotationClient struct {
	Url    string
	Client *http.Client
}

func (c *QuotationClient) GetCurrentUSDQuotation(ctx context.Context, code string) (*QuotationAPIResp, error) {
	url := fmt.Sprintf(c.Url, code)
	//TODO add timeout
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.StatusCode == http.StatusNotFound:
		return nil, errors.New("Coin not exists")
	case resp.StatusCode != http.StatusOK:
		text := fmt.Sprintf("QuotationAPI returned an unexpected status code: %s", resp.Status)
		return nil, errors.New(text)
	default:
		var q []QuotationAPIResp
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &q)
		if err != nil {
			return nil, err
		}
		return &q[0], nil
	}
}
