package external

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type QuotationMock struct{}

func (m QuotationMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	code := strings.Split(paths[2], "-")[0]
	switch code {
	case "error":
		responseFromFile(w, "mockResponses/empty.json", 500)
	default:
		responseFromFile(w, "mockResponses/success.json", 200)
	}

}

func responseFromFile(w http.ResponseWriter, file string, status int) {
	filePath := fmt.Sprintf("%s/%s", "/home/luizamaciel/estudos/challenge-bravo/external", file)
	body, _ := ioutil.ReadFile(filePath)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
