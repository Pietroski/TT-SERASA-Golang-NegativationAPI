package legacy_service

import (
	"fmt"
	"io"
	"net/http"
)

var (
	LegacyClient iLegacyClient = &sLegacyClient{}
)

type iLegacyClient interface {
	NewHTTPClient() *http.Client
	NewHTTPRequest(method string, url string, body io.Reader) *http.Request
	InitCaller(method string, url string, body io.Reader)
	Call() *http.Response
}

type sLegacyClient struct {
	Client  *http.Client
	Request *http.Request
}

func (lc *sLegacyClient) NewHTTPClient() *http.Client {
	httpClient := &http.Client{}

	return httpClient
}

func (lc *sLegacyClient) NewHTTPRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return req
}

func (lc *sLegacyClient) InitCaller(method string, url string, body io.Reader) {
	client := lc.NewHTTPClient()
	lc.Client = client
	request := lc.NewHTTPRequest(method, url, body)
	lc.Request = request
}

func (lc *sLegacyClient) Call() *http.Response {
	resp, err := lc.Client.Do(lc.Request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return resp
}
