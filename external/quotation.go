package external

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseUrl = "https://economia.awesomeapi.com.br/json/last/%s-USD"

type QuotationAPIResp struct {
	Code      string    `json:"code"`
	CodeIn    string    `json:"codein"`
	Ask       float64   `json:"ask"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetCurrentUSDQuotation(ctx context.Context, code string) (*QuotationAPIResp, error) {
	url := fmt.Sprintf(baseUrl, code)
	//TODO add timeout
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	var q QuotationAPIResp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &q)
	if err != nil {
		return nil, err
	}

	return &q, nil
}
