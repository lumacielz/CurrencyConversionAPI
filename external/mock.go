package external

import (
	"context"
	"fmt"
	"github.com/lumacielz/challenge-bravo/entities"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
)

type MockedHandler struct{}

var MockedServer = httptest.NewServer(MockedHandler{})

func (m MockedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	code := strings.Split(paths[2], "-")[0]
	switch code {
	case "error":
		responseFromFile(w, "mockResponses/empty.json", 500)
	case "notFound":
		responseFromFile(w, "mockResponses/error404.jsom", 404)
	case "empty":
		responseFromFile(w, "mockResponses/emptyArray.json", 200)
	default:
		responseFromFile(w, "mockResponses/success.json", 200)
	}

}

func responseFromFile(w http.ResponseWriter, file string, status int) {
	_, rootPath, _, _ := runtime.Caller(0)
	filePath := fmt.Sprintf("%s/%s", filepath.Dir(rootPath), file)
	body, _ := ioutil.ReadFile(filePath)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

type QuotationMock struct {
	Resp *entities.QuotationData
	Err  error
}

func (q QuotationMock) GetCurrentUSDQuotation(ctx context.Context, code string) (*entities.QuotationData, error) {
	return q.Resp, q.Err
}
