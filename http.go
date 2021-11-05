package xfers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Requester is http request interface.
type Requester interface {
	Call(ctx context.Context, method, url, apiKey, secretKey string, header http.Header, request interface{}, response interface{}) (statusCode int, err error)
}

type requester struct {
	client *http.Client
	logger Logger
}

func defaultRequester(client *http.Client, logger Logger) *requester {
	return &requester{
		client: client,
		logger: logger,
	}
}

// Call to prepare request and execute.
func (r *requester) Call(ctx context.Context, method, url, apiKey, secretKey string, header http.Header, request interface{}, response interface{}) (int, error) {
	now := time.Now()

	reqBody, err := json.Marshal(request)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	if header != nil {
		req.Header = header
	}

	req.SetBasicAuth(apiKey, secretKey)
	req.Header.Add("Content-Type", "application/vnd.api+json")
	req.Header.Add("Accept", "application/json")

	r.logger.Debug("%s %s", method, url)
	r.logRequestHeader(req.Header)
	r.logRequestBody(reqBody)
	defer func() { r.logger.Info("%s %s [%s]", method, url, time.Since(now)) }()

	return r.doRequest(req, response)
}

type errReponse struct {
	Errors []struct {
		Code   string `json:"code"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
	} `json:"errors"`
}

func (r *requester) doRequest(req *http.Request, response interface{}) (int, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	r.logResponseBody(resp.StatusCode, respBody)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errResp errReponse
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			r.logger.Error(err.Error())
			return http.StatusInternalServerError, ErrInternal
		}
		if len(errResp.Errors) > 0 {
			return resp.StatusCode, errors.New(errResp.Errors[0].Title)
		}
		return resp.StatusCode, errors.New(http.StatusText(resp.StatusCode))
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	return resp.StatusCode, nil
}

func (r *requester) logRequestHeader(header http.Header) {
	if header == nil || len(header) == 0 {
		return
	}

	for k, h := range header {
		for _, v := range h {
			r.logger.Debug("header: %s: %s", k, v)
		}
	}
}

func (r *requester) logRequestBody(request []byte) {
	if request == nil {
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, request, "", "  "); err != nil {
		r.logger.Error(err.Error())
		r.logger.Debug("request: %s", string(request))
		return
	}

	r.logger.Debug("request: %s", out.String())
}

func (r *requester) logResponseBody(code int, response []byte) {
	if response == nil {
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, response, "", "  "); err != nil {
		r.logger.Error(err.Error())
		r.logger.Debug("response: %d %s", code, string(response))
		return
	}

	r.logger.Debug("response: %d %s", code, out.String())
}
