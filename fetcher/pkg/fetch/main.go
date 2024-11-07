package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Rest interface {
	Get(context.Context, string, map[string]interface{}) (map[string]interface{}, error)
	Post(context.Context, string, map[string]interface{}) (map[string]interface{}, error)
}

type Fetch struct {
	httpClient *http.Client
}

func NewFetch() *Fetch {
	return &Fetch{
		httpClient: http.DefaultClient,
	}
}

func (f *Fetch) Get(ctx context.Context, url string, params map[string]interface{}) (map[string]interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (f *Fetch) Post(ctx context.Context, url string, payload map[string]interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
