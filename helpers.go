package plex

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClient struct {
	a *App
}

func (c *httpClient) getSetup(url string) (*httpRequest, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return &httpRequest{r: req, c: c}, nil
}

func (c *httpClient) get(url, token string) ([]byte, error) {
	r, err := c.getSetup(url)
	if err != nil {
		return nil, err
	}
	return r.do(token, true)
}

func (c *httpClient) command(url, token string) error {
	r, err := c.getSetup(url)
	if err != nil {
		return err
	}

	_, err = r.do(token, false)
	if err != nil {
		return err
	}

	return nil
}

func (c *httpClient) getAnon(url, v string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, strings.NewReader(v))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r := &httpRequest{r: req, c: c}
	return r.do("", true)
}

func (c *httpClient) post(url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r := &httpRequest{r: req, c: c}
	return r.do("", true)
}

type httpRequest struct {
	c *httpClient
	r *http.Request
}

func (r *httpRequest) do(token string, decode bool) ([]byte, error) {
	r.r.Header.Set("X-Plex-client-Identifier", r.c.a.o.clientIdentifier)
	r.r.Header.Set("Accept", "application/json")
	if token != "" {
		r.r.Header.Set("X-Plex-Token", token)
	}

	resp, err := r.c.a.o.client.Do(r.r)
	if err != nil {
		return nil, err
	}
	if decode {
		hr := &httpResponse{r: resp}
		return hr.read()
	}
	return nil, nil
}

type httpResponse struct {
	r *http.Response
}

func (r *httpResponse) read() ([]byte, error) {
	body, err := ioutil.ReadAll(r.r.Body)
	if err != nil {
		return nil, err
	}
	err = r.r.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
