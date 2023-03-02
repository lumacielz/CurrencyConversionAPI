package external

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lumacielz/challenge-bravo/entities"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type QuotationClient struct {
	Url    string
	Client *http.Client
}

func (c QuotationClient) GetCurrentUSDQuotation(ctx context.Context, code string) (*entities.QuotationData, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	url := fmt.Sprintf(c.Url, code)
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	errC := make(chan error, 1)
	respC := make(chan *http.Response, 1)
	go func() {
		resp, err := c.Client.Do(r)
		if err != nil {
			errC <- err
			return
		}
		respC <- resp
		close(respC)
		close(errC)
	}()

	select {
	case <-ctx.Done():
		return nil, entities.ErrQuotationAPITimeout
	case err := <-errC:
		return nil, err
	case resp := <-respC:
		switch {
		case resp.StatusCode == http.StatusNotFound:
			return nil, entities.ErrCurrencyNotFound

		case resp.StatusCode != http.StatusOK:
			log.Error(resp)
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

			if len(q) <= 0 {
				return nil, entities.ErrCurrencyNotFound
			}

			return &q[0], nil
		}
	}
}
