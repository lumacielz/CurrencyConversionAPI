package external

import (
	"context"
	"encoding/json"
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
	Client *http.Client
}

func (c *QuotationClient) GetCurrentUSDQuotation(ctx context.Context, code string) (*QuotationAPIResp, error) {
	url := fmt.Sprintf(baseUrl, code)
	//TODO add timeout
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

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
