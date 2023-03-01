package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lumacielz/challenge-bravo/entities"
	"io/ioutil"
	"net/http"
)

type QuotationClient struct {
	Url    string
	Client *http.Client
}

func (c QuotationClient) GetCurrentUSDQuotation(ctx context.Context, code string) (*entities.QuotationData, error) {
	url := fmt.Sprintf(c.Url, code)
	//TODO add timeout
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.StatusCode == http.StatusNotFound:
		return nil, entities.ErrCurrencyNotFound

	case resp.StatusCode != http.StatusOK:
		return nil, entities.ErrUnexpected(resp.Status)

	default:
		var q []entities.QuotationData
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
