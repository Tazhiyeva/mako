package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	model "mako/pkg/models"
	"net/http"
	"net/url"
	"time"
)

type DGISReviewFetcher struct {
	client  *http.Client
	baseURL string
}

// конструктор
func NewDGISReviewFetcher(baseURL string) *DGISReviewFetcher {
	return &DGISReviewFetcher{
		client:  &http.Client{Timeout: time.Second * 2},
		baseURL: baseURL,
	}
}

func (fetcher DGISReviewFetcher) FetchReview(params url.Values) (*model.Response, error) {
	fullURL := fetcher.baseURL + "?" + params.Encode()

	fmt.Println(fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := fetcher.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response model.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil

}
